// Package webauthn provides ...
package webauthn

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/trustasia-com/go-sdk/webauthn/types"
	"github.com/trustasia-com/go-sdk/wekey/session"

	"github.com/trustasia-com/go-van/pkg/server"
	"github.com/trustasia-com/go-van/pkg/server/httpx"
)

// WebAuthn instance for RP
type WebAuthn struct {
	client httpx.HTTPClient
	sess   *session.Session
}

// NewWebAuthn new WebAuthn instance
func NewWebAuthn(sess *session.Session) *WebAuthn {
	cli := httpx.NewClient(
		server.WithEndpoint(sess.Options.Host),
	)
	return &WebAuthn{sess: sess, client: cli}
}

// StartSignUp start registration process
func (authn *WebAuthn) StartSignUp(user types.User, req *http.Request) (*types.StartSignUpResp, error) {
	if req == nil {
		return nil, errors.New("sdk: http.Request is nil, please specify")
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	// json unmarshal
	input := types.StartSignUpReq{}
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
	sig := authn.sess.Sign(data)
	httpxReq := httpx.NewRequest(http.MethodPost, "/ta-fido-server/preregister", bytes.NewReader(data))
	httpxReq.AddHeader("Authenrization", sig)
	httpxResp, err := authn.client.Do(httpxReq)
	if err != nil {
		return nil, err
	}
	resp := &types.StartSignUpResp{}
	err = httpxResp.Scan(resp)
	return resp, err
}

// FinishSignUp registration process
func (authn *WebAuthn) FinishSignUp(user types.User, req *http.Request) (*types.FinishSignInResp, error) {
	if req == nil {
		return nil, errors.New("sdk: http.Request is nil, please specify")
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	sig := authn.sess.Sign(data)
	httpxReq := httpx.NewRequest(http.MethodPost, "/ta-fido-server/preregister", bytes.NewReader(data))
	httpxReq.AddHeader("Authenrization", sig)
	httpxResp, err := authn.client.Do(httpxReq)
	if err != nil {
		return nil, err
	}
	resp := &types.FinishSignInResp{}
	err = httpxResp.Scan(resp)
	return resp, nil
}

// StartSignIn start login
func (authn *WebAuthn) StartSignIn(user types.User, req *http.Request) (*types.StartSignInResp, error) {
	if req == nil {
		return nil, errors.New("sdk: http.Request is nil, please specify")
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	// json unmarshal
	input := types.StartSignInReq{}
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
	sig := authn.sess.Sign(data)
	httpxReq := httpx.NewRequest(http.MethodPost, "/ta-fido-server/preauthenticate", bytes.NewReader(data))
	httpxReq.AddHeader("Authenrization", sig)
	httpxResp, err := authn.client.Do(httpxReq)
	if err != nil {
		return nil, err
	}
	resp := &types.StartSignInResp{}
	err = httpxResp.Scan(resp)
	return resp, err
}

// FinishSignIn finish login
func (authn *WebAuthn) FinishSignIn(user types.User, req *http.Request) (*types.FinishSignInResp, error) {
	if req == nil {
		return nil, errors.New("sdk: http.Request is nil, please specify")
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	sig := authn.sess.Sign(data)
	httpxReq := httpx.NewRequest(http.MethodPost, "/ta-fido-server/preauthenticate", bytes.NewReader(data))
	httpxReq.AddHeader("Authenrization", sig)
	httpxResp, err := authn.client.Do(httpxReq)
	if err != nil {
		return nil, err
	}
	resp := &types.FinishSignInResp{}
	err = httpxResp.Scan(resp)
	return resp, nil
}

// DeleteCredential delete exists credential
func (authn *WebAuthn) DeleteCredential() {

}

// SelectCredentials query credential list
func (authn *WebAuthn) SelectCredentials() {

}
