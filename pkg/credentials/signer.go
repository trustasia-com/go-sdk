// Package credentials provides ...
package credentials

import (
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/trustasia-com/go-van/pkg/server/httpx"
)

// Signature and API related constants.
const (
	signAlgorithmHMAC = "WEKEY-HMAC-SHA256"
	iso8601DateFormat = "20060102T150405Z"
)

// http header
const (
	httpHeaderDate        = "X-WeKey-Date"
	httpHeaderCredential  = "X-WeKey-Credential"
	httpHeaderContentHash = "X-WeKey-Content-Hash"
	httpHeader

	httpHeaderAuthorization = "Authorization"
)

// SignerDefault signatureDefault signer
//   payload: query or body
//   header
func SignerDefault(req *httpx.Request, accessKey, secretKey, location string, payload []byte) {
	// set headers
	now := time.Now().UTC()

	req.SetHeader(httpHeaderCredential, accessKey+"/"+location)
	req.SetHeader(httpHeaderDate, now.Format(iso8601DateFormat))
	hash := sum256(payload)
	req.SetHeader(httpHeaderContentHash, hash)

	h := req.GetHeader()
	signedHeaders := getSignedHeaders(h)
	stringToSign := getStringToSign(h, signedHeaders)
	signature := sumHMAC([]byte(secretKey), []byte(stringToSign))
	auth := []string{
		strings.Join(signedHeaders, ";"),
		signature,
	}
	req.SetHeader(httpHeaderAuthorization, signAlgorithmHMAC+" "+strings.Join(auth, ","))
}

// ValidateDefault validate signature
func ValidateDefault(resp *httpx.Response, secretKey string, payload []byte) error {
	httpResp := resp.HTTP()
	auth := httpResp.Header.Get(httpHeaderAuthorization)
	// check auth header
	if auth == "" {
		return ErrNotFoundAuthorizationHeader
	}
	// check content hash
	hash := httpResp.Header.Get(httpHeaderContentHash)
	if sum256(payload) != hash {
		return ErrNotMatchedPayloadHash
	}
	// check algorithm
	params := strings.Split(auth, " ")
	if len(params) != 2 {
		return ErrInvalidAuthorizationHeader
	}
	algo := params[0]
	if algo != signAlgorithmHMAC {
		return ErrNotMatchedAlgorithmServer
	}
	// check headers
	params = strings.Split(params[1], ",")
	if len(params) != 2 {
		return ErrInvalidAuthorizationHeader
	}
	signedHeaders := strings.Split(params[0], ";")
	stringToSign := getStringToSign(httpResp.Header, signedHeaders)

	gotSig := params[1]
	expectSig := sumHMAC([]byte(secretKey), []byte(stringToSign))
	if expectSig != gotSig {
		return errors.Wrapf(ErrInvalidSignature, "expected %s got %s", expectSig, gotSig)
	}
	return nil
}

func getSignedHeaders(h http.Header) []string {
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
	return hs
}

func getStringToSign(h http.Header, headers []string) string {
	stringToSign := signAlgorithmHMAC
	for _, k := range headers {
		v := h.Get(k)
		stringToSign += "\n" + v
	}
	return stringToSign
}
