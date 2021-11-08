// Package credentials provides ...
package credentials

import (
	"net/url"
	"time"

	"github.com/pkg/errors"
	"github.com/trustasia-com/go-van/pkg/server/httpx"
)

// Signature and API related constants.
const (
	signAlgorithmHMAC = "WEKEY-HMAC-SHA256"
	iso8601DateFormat = "20060102T150405Z"
)

// Signer sign the request before Do()
type Signer func(req *httpx.Request, accessKey, secretKey string, payload []byte)

// Validator validate the request data
type Validator func(resp *httpx.Response, secretKey string, payload []byte) error

// SignerDefault signatureDefault signer
//   query
//   payload: query or body
//   header
func SignerDefault(req *httpx.Request, accessKey, secretKey string, payload []byte) {
	// init time
	now := time.Now().UTC()

	hash := sum256(payload)

	vals := url.Values{}
	vals.Set(httpHeaderDate, now.Format(iso8601DateFormat))
	vals.Set(httpHeaderAlgorithm, signAlgorithmHMAC)
	vals.Set(httpHeaderAccessKey, accessKey)
	vals.Set(httpHeaderContentHash, hash)
	stringToSign := vals.Encode()

	sig := sumHMAC([]byte(secretKey), []byte(stringToSign))
	vals.Set(httpHeaderSignature, sig)
	req.AddHeader(httpHeaderAuthorization, vals.Encode())
}

// ValidateDefault validate signature
func ValidateDefault(resp *httpx.Response, secretKey string, payload []byte) error {
	httpResp := resp.ToHTTP()
	auth := httpResp.Header.Get("Authorization")
	if auth == "" {
		return errors.New("sdk: not found http header: Authorization")
	}
	vals, err := url.ParseQuery(auth)
	if err != nil {
		return err
	}
	gotSig := vals.Get(httpHeaderSignature)
	vals.Del(httpHeaderSignature)

	// check content hash
	hash := sum256(payload)
	if val := vals.Get(httpHeaderContentHash); val != hash {
		return errors.Errorf("sdk: check content hash failed: expected %s got %s",
			hash, val)
	}
	stringToSign := vals.Encode()
	expectSig := sumHMAC([]byte(secretKey), []byte(stringToSign))
	if expectSig != gotSig {
		return errors.Errorf("sdk: invalid signature: expected %s got %s", expectSig, gotSig)
	}
	return nil
}
