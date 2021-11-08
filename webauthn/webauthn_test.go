// Package webauthn provides ...
package webauthn

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/trustasia-com/go-sdk/pkg/credentials"
)

var (
	webauthn *WebAuthn
	u        *user
)

func init() {
	sess, err := credentials.New(credentials.Options{
		AccessKey: "test",
		SecretKey: "secret",
		Endpoint:  "http://localhost:9000",
	}, false)
	if err != nil {
		panic(err)
	}
	webauthn = New(sess)
	u = &user{}
}

type user struct{}

// User ID according to the Relying Party
func (u *user) ID() []byte {
	return []byte("userID")
}

// Display Name of the user
func (u *user) DisplayName() string {
	return "displayName"
}

// User Name according to the Relying Party
func (u *user) Name() string {
	return "name"
}

// Deprecated: User's Icon url
func (u *user) Icon() string {
	return ""
}

func TestStartSignUp(t *testing.T) {
	data := `{"authenticatorSelection":{"userVerification":"preferred","requireResidentKey":true},"attestation":"direct","extensions":{}}`
	req := httptest.NewRequest(http.MethodPost, "/ta-fido-server/preregister", strings.NewReader(data))
	req.Header.Set("Content-Type", "application/json")

	resp, err := webauthn.StartSignUp(u, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestFinishSignUp(t *testing.T) {

}

func TestStartSignIn(t *testing.T) {

}

func TestFinishSignIn(t *testing.T) {

}
