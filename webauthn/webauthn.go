// Package webauthn provides ...
package webauthn

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/trustasia-com/go-sdk/pkg/client"
	"github.com/trustasia-com/go-sdk/pkg/credentials"
	"github.com/trustasia-com/go-sdk/pkg/fido"
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
	client *client.HTTPClient
}

// New new WebAuthn instance
func New(sess *credentials.Session) *WebAuthn {
	return &WebAuthn{
		client: client.NewHTTPClient(sess),
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
		req.Attestation = fido.PreferenceNone
	}
	if !fido.IsValidAttestationCP(req.Attestation) {
		return nil, errors.New("Invalid req.Attestation value")
	}

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	scope := "fido-server/" + userID
	msg, err := authn.client.Request(http.MethodPost, apiPreregister, scope, data)
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
	msg, err := authn.client.Request(http.MethodPost, apiRegister, scope, data)
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
		req.UserVerification = fido.VerificationPreferred
	}
	if !fido.IsValidUserVerificationRequirement(req.UserVerification) {
		return nil, errors.New("Invalid req.UserVerification value")
	}
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	scope := "fido-server/" + userID
	msg, err := authn.client.Request(http.MethodPost, apiPreauthenticate, scope, data)
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
	msg, err := authn.client.Request(http.MethodPost, apiAuthenticate, scope, data)
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
