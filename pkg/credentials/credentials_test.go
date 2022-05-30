// Package credentials provides ...
package credentials

import (
	"net/http"
	"strings"
	"testing"
)

var (
	sess *Session
)

func init() {
	var err error
	sess, err = New(Options{AccessKey: "accessKey", SecretKey: "secretKey"}, false)
	if err != nil {
		panic(err)
	}
}

func TestSession_Sign(t *testing.T) {
	sig := sess.SumHMAC([]byte("hello world"))
	t.Log(sig)
}

func TestSession_SignWithRequest(t *testing.T) {
	// get
	req, err := http.NewRequest(http.MethodGet, "/?page=1&page_size=10", nil)
	if err != nil {
		panic(err)
	}
	sess.SignRequest(req, "id")
	t.Log(req)

	// post
	body := `{"page":1,"page_size":10}`
	req, err = http.NewRequest(http.MethodPost, "httP://127.0.0.1", strings.NewReader(body))
	if err != nil {
		panic(err)
	}
	sess.SignRequest(req, "id/12")
	t.Log(req)

	cred, err := ValidateDefault(req, sess.Options.SecretKey)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cred)
}
