// Package credentials provides ...
package credentials

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/trustasia-com/go-van/pkg/server/httpx"
)

// api host
const (
	// test host
	endpointTest = "https://api-test.wekey.com"
	// prod host
	endpointProd = "https://api.wekey.com"
)

// http header
const (
	httpHeaderDate        = "X-WeKey-Date"
	httpHeaderAlgorithm   = "X-WeKey-Algorithm"
	httpHeaderAccessKey   = "X-WeKey-AccessKey"
	httpHeaderContentHash = "X-WeKey-Content-Hash"
	httpHeaderSignature   = "X-WeKey-Signature"

	httpHeaderAuthorization = "Authorization"
)

// SignatureType is type of Authorization requested for a given request.
type SignatureType int

// SignatureType enums
const (
	// hmac-sha256
	SignatureDefault SignatureType = iota
	// and more
)

// Options session options
type Options struct {
	// credential
	AccessKey string
	SecretKey string
	// you can custom fido server if you are privacy
	// eg. https://fido.example.com
	Endpoint   string
	SignerType SignatureType
}

// Session provides a central location to create service clients from
type Session struct {
	Options Options
}

// New session
func New(options Options, isProd bool) (*Session, error) {
	// check options
	if options.AccessKey == "" || options.SecretKey == "" {
		return nil, errors.New("sdk: accessKey or secretKey not specified")
	}
	if options.Endpoint == "" {
		if isProd {
			options.Endpoint = endpointProd
		} else {
			options.Endpoint = endpointTest
		}
	}
	sess := &Session{Options: options}
	return sess, nil
}

// Sign data
func (sess *Session) Sign(data []byte) string {
	return sumHMAC([]byte(sess.Options.SecretKey), data)
}

// SignWithRequest sign & set request header
func (sess *Session) SignWithRequest(req *httpx.Request, payload []byte) {
	var signer Signer
	switch sess.Options.SignerType {

	default:
		signer = SignerDefault
	}
	signer(req, sess.Options.AccessKey, sess.Options.SecretKey, payload)
}

// ValidateSig validate the authorization
func (sess *Session) ValidateSig(resp *httpx.Response) error {
	var validator Validator
	switch sess.Options.SignerType {

	default:
		validator = ValidateDefault
	}
	return validator(resp, sess.Options.SecretKey, resp.Data)
}

// sumHMAC calculate hmac between two input byte array.
func sumHMAC(secret, payload []byte) string {
	h := hmac.New(sha256.New, secret)
	h.Write(payload)
	return hex.EncodeToString(h.Sum(nil))
}

// sum256 calculate sha256 sum for an input byte array.
func sum256(data []byte) string {
	h := sha256.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
