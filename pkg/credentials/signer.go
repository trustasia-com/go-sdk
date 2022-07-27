// Package credentials provides ...
package credentials

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// Signer sign the request before Do()
type Signer func(req *http.Request, accessKey, secretKey, scope string) error

// Validator validate the request data
type Validator func(req *http.Request, secretKey string) ([]string, error)

// Signature and API related constants.
const (
	signAlgorithmHMAC = "WEKEY-HMAC-SHA256"
	iso8601DateFormat = "20060102T150405Z"
)

// http header
const (
	httpHeaderDate = "X-WeKey-Date"
	httpHeaderHost = "Host"

	httpHeaderAuthorization = "Authorization"
)

// SignerDefault signatureDefault signer
//   payload: query or body
//   header
// https://docs.aws.amazon.com/general/latest/gr/sigv4-create-canonical-request.html
// eg.
//   GET
//   /
//   Action=ListUsers&Version=2010-05-08
//   content-type:application/x-www-form-urlencoded; charset=utf-8
//   host:iam.amazonaws.com
//   x-amz-date:20150830T123600Z
//
//   content-type;host;x-amz-date
//   e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
func SignerDefault(req *http.Request, accessKey, secretKey, scope string) error {
	// must Host header
	req.Header.Set(httpHeaderHost, req.Host)

	canonicalReq := bytes.Buffer{}
	canonicalReq.WriteString(req.Method + "\n") // method
	path := req.URL.Path
	if path == "" {
		path = "/"
	}
	canonicalReq.WriteString(path + "\n") // path
	vals, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		return err
	}
	canonicalReq.WriteString(vals.Encode() + "\n") // query
	// other headers
	signedHeaders, canonicalHeaders := getCanonicalHeaders(req.Header)
	canonicalReq.WriteString(canonicalHeaders + "\n") // headers
	canonicalReq.WriteString(signedHeaders + "\n")    // headers
	var data []byte
	if req.GetBody != nil {
		body, err := req.GetBody()
		if err != nil {
			return err
		}
		data, err = io.ReadAll(body)
		if err != nil {
			return err
		}
	}
	hash := sum256(data)
	canonicalReq.WriteString(hash) // body
	hash = sum256(canonicalReq.Bytes())

	// set headers
	date := time.Now().UTC().Format(iso8601DateFormat)
	req.Header.Set(httpHeaderDate, date)
	credential := accessKey + "/" + scope

	stringToSign := fmt.Sprintf("%s\n%s\n%s\n%s",
		signAlgorithmHMAC,
		date,
		scope,
		hash,
	)
	signature := sumHMAC([]byte(secretKey), []byte(stringToSign))
	req.Header.Set(httpHeaderAuthorization, signAlgorithmHMAC+" "+credential+","+signedHeaders+","+signature)
	return nil
}

// ValidateDefault validate signature
func ValidateDefault(req *http.Request, secretKey string) ([]string, error) {
	// Authorization: algorithm <access key ID>/<credential scope>,SignedHeaders,signature
	auth := req.Header.Get(httpHeaderAuthorization)
	// check auth header
	if auth == "" {
		return nil, ErrNotFoundAuthorizationHeader
	}
	params := strings.Fields(auth)
	if len(params) != 2 {
		return nil, ErrInvalidAuthorizationHeader
	}
	algo := params[0]
	if algo != signAlgorithmHMAC {
		return nil, ErrNotMatchedAlgorithmServer
	}
	params = strings.Split(params[1], ",")
	if len(params) != 3 {
		return nil, ErrInvalidAuthorizationHeader
	}
	// credential
	cred := strings.SplitN(params[0], "/", 2)
	if len(cred) != 2 {
		return nil, ErrInvalidCredentialHeader
	}
	// validate signature
	canonicalReq := bytes.Buffer{}
	canonicalReq.WriteString(req.Method + "\n") // method
	path := req.URL.RawPath
	if path == "" {
		path = "/"
	}
	canonicalReq.WriteString(path + "\n") // path
	vals, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		return nil, err
	}
	canonicalReq.WriteString(vals.Encode() + "\n") // query
	signedHeaders := params[1]
	hostHeader, canonicalHeaders := getCanonicalSignedHeaders(req.Header, signedHeaders)
	if !hostHeader {
		return nil, ErrInvalidHostHeader
	}
	canonicalReq.WriteString(canonicalHeaders + "\n") // headers
	canonicalReq.WriteString(signedHeaders + "\n")    // header
	var data []byte
	if req.Body != http.NoBody {
		var err error
		data, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		req.Body = io.NopCloser(bytes.NewReader(data))
	}
	hash := sum256(data)
	canonicalReq.WriteString(hash) // body
	hash = sum256(canonicalReq.Bytes())

	date := req.Header.Get(httpHeaderDate)
	stringToSign := fmt.Sprintf("%s\n%s\n%s\n%s",
		signAlgorithmHMAC,
		date,
		cred[1], // scope
		hash,
	)
	expectSig := sumHMAC([]byte(secretKey), []byte(stringToSign))
	if expectSig != params[2] {
		return nil, errors.Wrapf(ErrInvalidSignature, "expected %s got %s", expectSig, params[2])
	}
	return cred, nil
}

func getCanonicalHeaders(h http.Header) (string, string) {
	var hs []string

	for k := range h {
		if k == "User-Agent" {
			continue
		}
		if k == "Authorization" {
			continue
		}
		hs = append(hs, strings.ToLower(k))
	}
	sort.Strings(hs)

	buf := bytes.Buffer{}
	for _, v := range hs {
		buf.WriteString(strings.ToLower(v) + ":" + h.Get(v) + "\n")
	}
	return strings.Join(hs, ";"), buf.String()
}

func getCanonicalSignedHeaders(h http.Header, signedHeaders string) (bool, string) {
	headers := strings.Split(signedHeaders, ";")

	buf := bytes.Buffer{}
	hostHeader := false
	for _, v := range headers {
		lv := strings.ToLower(v)
		buf.WriteString(v + ":" + h.Get(v) + "\n")

		if lv == "host" {
			hostHeader = true
		}
	}
	return hostHeader, buf.String()
}
