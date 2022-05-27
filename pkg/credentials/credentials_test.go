// Package credentials provides ...
package credentials

import (
	"net/http"
	"testing"

	"github.com/trustasia-com/go-van/pkg/server/httpx"
)

var (
	sess *Session
)

func init() {
	var err error
	sess, err = New(Options{
		AccessKey: "accessKey",
		SecretKey: "secretKey",
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
	query := "page=1&page_size=10"
	req := httpx.NewRequest(http.MethodGet, "/", query, nil)
	sess.SignRequest(req, "id", nil)
	t.Log(req)

	// post
	body := `{"page":1,"page_size":10}`
	req = httpx.NewRequest(http.MethodPost, "/", "", []byte(body))
	sess.SignRequest(req, "id", []byte(body))
	t.Log(req)
}
