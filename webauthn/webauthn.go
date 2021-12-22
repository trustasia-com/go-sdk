// Package webauthn provides ...
package webauthn

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/trustasia-com/go-sdk/pkg/credentials"
	"github.com/trustasia-com/go-sdk/pkg/types"

	"github.com/trustasia-com/go-van/pkg/server"
	"github.com/trustasia-com/go-van/pkg/server/httpx"
)

// WebAuthn instance for RP
type WebAuthn struct {
	client httpx.Client
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
func (authn *WebAuthn) StartSignUp(req StartSignUpReq, userID string) (*StartSignUpResp, error) {
	// check input
	if req.Username == "" {
		return nil, errors.New("Need specify req.Username")
	}
	if req.DisplayName == "" {
		return nil, errors.New("Need specify req.DisplayName")
	}
	if req.Attestation == "" {
		req.Attestation = types.PreferenceNone
	}
	if !types.IsValidAttestationCP(req.Attestation) {
		return nil, errors.New("Invalid req.Attestation value")
	}

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	loc := "fido/" + userID
	httpxReq := httpx.NewRequest(http.MethodPost, "/ta-fido-server/preregister", bytes.NewReader(data))
	authn.sess.SignRequest(httpxReq, loc, data)
	httpxResp, err := authn.httpRequest(httpxReq)
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
func (authn *WebAuthn) FinishSignUp(req *http.Request) (*FinishSignUpResp, error) {
	if req == nil {
		return nil, errors.New("sdk: http.Request is nil, please specify")
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	httpxReq := httpx.NewRequest(http.MethodPost, "/ta-fido-server/register", bytes.NewReader(data))
	authn.sess.SignRequest(httpxReq, "fido/", data)
	httpxResp, err := authn.httpRequest(httpxReq)
	if err != nil {
		return nil, err
	}
	resp := &FinishSignUpResp{}
	err = httpxResp.Scan(resp)
	return resp, nil
}

// StartSignIn start login
func (authn *WebAuthn) StartSignIn(req StartSignInReq, userID string) (*StartSignInResp, error) {
	if userID != "" {
		if req.Username == "" {
			return nil, errors.New("Need specify req.Username")
		}
		if req.DisplayName == "" {
			req.DisplayName = req.Username
		}
	}
	if req.UserVerification == "" {
		req.UserVerification = types.VerificationPreferred
	}
	if !types.IsValidUserVerificationRequirement(req.UserVerification) {
		return nil, errors.New("Invalid req.UserVerification value")
	}
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	loc := "fido/" + userID
	httpxReq := httpx.NewRequest(http.MethodPost, "/ta-fido-server/preauthenticate", bytes.NewReader(data))
	authn.sess.SignRequest(httpxReq, loc, data)
	httpxResp, err := authn.httpRequest(httpxReq)
	if err != nil {
		return nil, err
	}
	resp := &StartSignInResp{
		AllowCredentials: []types.PublicKeyCredentialDescriptor{},
		Extensions:       types.AuthenticationExtensionsClientInputs{},
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
	authn.sess.SignRequest(httpxReq, "fido/", data)
	httpxResp, err := authn.httpRequest(httpxReq)
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

func (authn *WebAuthn) httpRequest(req *httpx.Request) (httpx.Response, error) {
	httpxResp, err := authn.client.Do(context.Background(), req)
	if err != nil {
		if len(httpxResp.Data) == 0 {
			var badResp struct {
				Code  int    `json:"code"`
				Error string `json:"error"`
			}
			err2 := httpxResp.Scan(&badResp)
			if err2 == nil {
				err = fmt.Errorf("code: %d, err: %s", badResp.Code, badResp.Error)
			}
		}
		return httpxResp, err
	}
	return httpxResp, nil
}
