// Package session provides ...
package session

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
)

// Session provides a central location to create service clients from
type Session struct {
	Options Options
}

// New session
func New(opts ...Option) (*Session, error) {
	options := Options{Host: testHost}
	for _, o := range opts {
		o(&options)
	}
	// check options
	if options.AccessKey == "" || options.SecretKey == "" {
		return nil, errors.New("sdk: accessKey or secretKey not specified")
	}
	sess := &Session{Options: options}
	return sess, nil
}

// Must is a helper function to ensure the session is valid and there was no
// error when calling a New function.
func Must(sess *Session, err error) *Session {
	if err != nil {
		panic(err)
	}
	return sess
}

// Sign data
// <accessKey>#<data>
func (sess *Session) Sign(data []byte) string {
	h := hmac.New(sha1.New, []byte(sess.Options.SecretKey))
	h.Write([]byte(sess.Options.AccessKey))
	h.Write([]byte("#"))
	h.Write(data)
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}
