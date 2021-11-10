// Package webauthn provides ...
package webauthn

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/trustasia-com/go-sdk/pkg/credentials"
	"github.com/trustasia-com/go-sdk/pkg/types"

	"github.com/trustasia-com/go-van/pkg/server"
	"github.com/trustasia-com/go-van/pkg/server/httpx"
)

// WebAuthn instance for RP
type WebAuthn struct {
	client httpx.HTTPClient
	sess   *credentials.Session
}

// New new WebAuthn instance
func New(sess *credentials.Session) *WebAuthn {
	cli := httpx.NewClient(
		server.WithEndpoint(sess.Options.Endpoint),
	)
	return &WebAuthn{sess: sess, client: cli}
}

// StartSignUp start registration process
func (authn *WebAuthn) StartSignUp(req *http.Request, user User) (*StartSignUpResp, error) {
	if req == nil {
		return nil, errors.New("sdk: http.Request is nil, please specify")
	}
	if user == nil {
		return nil, errors.New("sdk: user is nil, please specify")
	}
	loc := "fido/" + hex.EncodeToString(user.ID())
	data, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	// json unmarshal
	input := StartSignUpReq{}
	err = json.Unmarshal(data, &input)
	if err != nil {
		return nil, err
	}
	input.Username = user.Name()
	input.DisplayName = user.DisplayName()

	data, err = json.Marshal(input)
	if err != nil {
		return nil, err
	}
	httpxReq := httpx.NewRequest(http.MethodPost, "/ta-fido-server/preregister", bytes.NewReader(data))
	authn.sess.SignWithRequest(httpxReq, loc, data)
	httpxResp, err := authn.client.Do(httpxReq)
	if err != nil {
		return nil, err
	}
	resp := &StartSignUpResp{
		ExcludeCredentials: []types.PublicKeyCredentialDescriptor{},
		Extensions:         types.AuthenticationExtensionsClientInputs{},
	}
	err = httpxResp.Scan(resp)
	return resp, err
}

// FinishSignUp registration process
func (authn *WebAuthn) FinishSignUp(req *http.Request, user User) (*FinishSignUpResp, error) {
	if req == nil {
		return nil, errors.New("sdk: http.Request is nil, please specify")
	}
	if user == nil {
		return nil, errors.New("sdk: user is nil, please specify")
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	httpxReq := httpx.NewRequest(http.MethodPost, "/ta-fido-server/register", bytes.NewReader(data))
	authn.sess.SignWithRequest(httpxReq, "fido/"+hex.EncodeToString(user.ID()), data)
	httpxResp, err := authn.client.Do(httpxReq)
	if err != nil {
		return nil, err
	}
	resp := &FinishSignUpResp{}
	err = httpxResp.Scan(resp)
	return resp, nil
}

// StartSignIn start login
func (authn *WebAuthn) StartSignIn(req *http.Request, user User) (*StartSignInResp, error) {
	if req == nil {
		return nil, errors.New("sdk: http.Request is nil, please specify")
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	// json unmarshal
	input := StartSignInReq{}
	err = json.Unmarshal(data, &input)
	if err != nil {
		return nil, err
	}
	data, err = json.Marshal(input)
	if err != nil {
		return nil, err
	}
	loc := "fido/"
	if user != nil {
		input.Username = user.Name()
		input.DisplayName = user.DisplayName()

		loc += hex.EncodeToString(user.ID())
	}
	httpxReq := httpx.NewRequest(http.MethodPost, "/ta-fido-server/preauthenticate", bytes.NewReader(data))
	authn.sess.SignWithRequest(httpxReq, loc, data)
	httpxResp, err := authn.client.Do(httpxReq)
	if err != nil {
		return nil, err
	}
	resp := &StartSignInResp{
		ExcludeCredentials: []types.PublicKeyCredentialDescriptor{},
		Extensions:         types.AuthenticationExtensionsClientInputs{},
	}
	err = httpxResp.Scan(resp)
	return resp, err
}

// FinishSignIn finish login
func (authn *WebAuthn) FinishSignIn(req *http.Request) (*FinishSignInResp, error) {
	if req == nil {
		return nil, errors.New("sdk: http.Request is nil, please specify")
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	httpxReq := httpx.NewRequest(http.MethodPost, "/ta-fido-server/authenticate", bytes.NewReader(data))
	authn.sess.SignWithRequest(httpxReq, "fido/", data)
	httpxResp, err := authn.client.Do(httpxReq)
	if err != nil {
		return nil, err
	}
	resp := &FinishSignInResp{}
	err = httpxResp.Scan(resp)
	return resp, nil
}

// DeleteCredential delete exists credential
func (authn *WebAuthn) DeleteCredential() {

}

// SelectCredentials query credential list
func (authn *WebAuthn) SelectCredentials() {

}
