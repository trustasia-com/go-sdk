// Package credentials provides ...
package credentials

import (
	"net/http"
	"strings"
	"testing"

	"github.com/trustasia-com/go-van/pkg/server/httpx"
)

var (
	sess *Session
)

func init() {
	var err error
	sess, err = New(Options{
		AccessKey: "id",
		SecretKey: "hello",
	}, false)
	if err != nil {
		panic(err)
	}
}

func TestSession_Sign(t *testing.T) {
	sig := sess.Sign([]byte("hello world"))
	t.Log(sig)
}

func TestSession_SignWithRequest(t *testing.T) {
	// get
	payload := "page=1&page_size=10"
	req := httpx.NewRequest(http.MethodGet, "/?"+payload, nil)
	sess.SignWithRequest(req, []byte(payload))
	t.Log(req)

	// post
	payload = `{"page":1,"page_size":10}`
	req = httpx.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
	sess.SignWithRequest(req, []byte(payload))
	t.Log(req)
}
