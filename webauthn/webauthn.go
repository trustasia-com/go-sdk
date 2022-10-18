// Package webauthn provides ...
package webauthn

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/trustasia-com/go-sdk/pkg"
	"github.com/trustasia-com/go-sdk/pkg/client"
	"github.com/trustasia-com/go-sdk/pkg/credentials"
	"github.com/trustasia-com/go-sdk/pkg/message"
	"github.com/trustasia-com/go-sdk/pkg/types"
)

// api list
const (
	apiPreregister     = "/ta-fido-server/preregister"
	apiRegister        = "/ta-fido-server/register"
	apiPreauthenticate = "/ta-fido-server/preauthenticate"
	apiAuthenticate    = "/ta-fido-server/authenticate"
)

// WebAuthn instance for RP
type WebAuthn struct {
	userAgent string
	sess      *credentials.Session
	client    *http.Client
}

// New new WebAuthn instance
func New(sess *credentials.Session) *WebAuthn {
	return &WebAuthn{
		userAgent: pkg.BuildUserAgent(),
		sess:      sess,
		client:    client.NewHTTPClient(),
	}
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
	scope := "fido-server/" + userID
	msg, err := authn.httpRequest(http.MethodPost, apiPreregister, scope, data)
	if err != nil {
		return nil, err
	}
	resp := &StartSignUpResp{}
	err = json.Unmarshal(msg.Data, resp)
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
	scope := "fido-server/"
	msg, err := authn.httpRequest(http.MethodPost, apiRegister, scope, data)
	if err != nil {
		return nil, err
	}
	resp := &FinishSignUpResp{}
	err = json.Unmarshal(msg.Data, resp)
	return resp, err
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
	scope := "fido-server/" + userID
	msg, err := authn.httpRequest(http.MethodPost, apiPreauthenticate, scope, data)
	if err != nil {
		return nil, err
	}
	resp := &StartSignInResp{}
	err = json.Unmarshal(msg.Data, resp)
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
	scope := "fido-server/"
	msg, err := authn.httpRequest(http.MethodPost, apiAuthenticate, scope, data)
	if err != nil {
		return nil, err
	}
	resp := &FinishSignInResp{}
	err = json.Unmarshal(msg.Data, resp)
	return resp, err
}

// DeleteCredential delete exists credential
func (authn *WebAuthn) DeleteCredential(userID string) {

}

// SelectCredentials query credential list
func (authn *WebAuthn) SelectCredentials() {

}

// DestroyUser delete user's all credentials
func (authn *WebAuthn) DestroyUser() {

}

func (authn *WebAuthn) httpRequest(method, path, scope string, data []byte) (*message.JSONRawMessage, error) {
	var (
		httpReq *http.Request
		err     error
	)
	url := authn.sess.Options.Endpoint + path
	if len(data) > 0 {
		httpReq, err = http.NewRequest(method, url, bytes.NewReader(data))
	} else {
		httpReq, err = http.NewRequest(method, url, nil)
	}
	httpReq.Header.Set("User-Agent", authn.userAgent)

	if err = authn.sess.SignRequest(httpReq, scope); err != nil {
		return nil, err
	}
	httpResp, err := authn.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	data, err = io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	resp := &message.JSONRawMessage{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, fmt.Errorf("code: %d, err: %s", resp.Code, resp.Error)
	}
	if resp.Code != 0 {
		return nil, errors.New(resp.Error)
	}
	return resp, err
}
