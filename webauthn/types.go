// Package webauthn provides ...
package webauthn

import "github.com/trustasia-com/go-sdk/pkg/types"

// User interface
type User interface {
	// User ID according to the Relying Party
	ID() []byte
	// Display Name of the user
	DisplayName() string
	// User Name according to the Relying Party
	Name() string
	// Deprecated: User's Icon url
	Icon() string
}

// StartSignUpReq sign up request
type StartSignUpReq struct {
	Username               string                                     `json:"username"`
	DisplayName            string                                     `json:"displayName"`
	Attestation            types.AttestationConveyancePreference      `json:"attestation"`
	AuthenticatorSelection types.AuthenticatorSelectionCriteria       `json:"authenticatorSelection"`
	Extensions             types.AuthenticationExtensionsClientInputs `json:"extensions"`
}

// StartSignUpResp sign up response
type StartSignUpResp struct {
	RP                     types.PublicKeyCredentialRpEntity          `json:"rp"`
	User                   types.PublicKeyCredentialUserEntity        `json:"user"`
	Challenge              string                                     `json:"challenge"`
	PubKeyCredParams       []types.PublicKeyCredentialType            `json:"pubKeyCredParams"`
	Timeout                int                                        `json:"timeout"`
	ExcludeCredentials     []types.PublicKeyCredentialDescriptor      `json:"excludeCredentials"`
	AuthenticatorSelection types.AuthenticatorSelectionCriteria       `json:"authenticatorSelection"`
	Attestation            types.AttestationConveyancePreference      `json:"attestation"`
	Extensions             types.AuthenticationExtensionsClientInputs `json:"extensions"`
}

// FinishSignUpReq sign up request
type FinishSignUpReq struct {
	ID       string                        `json:"id"`
	RawID    string                        `json:"rawId"`
	Type     types.PublicKeyCredentialType `json:"type"`
	Response struct {
		ClientDataJSON    string `json:"clientDataJSON"`
		AttestationObject string `json:"attestationObject"`
	} `json:"response"`
}

// FinishSignUpResp sign up response
type FinishSignUpResp struct{}

// StartSignInReq sign in request
type StartSignInReq = StartSignUpReq

// StartSignInResp sign in response
type StartSignInResp = StartSignUpResp

// FinishSignInReq sign in request
type FinishSignInReq = FinishSignUpReq

// FinishSignInResp sign in response
type FinishSignInResp struct{}
