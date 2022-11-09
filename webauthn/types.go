// Package webauthn provides ...
package webauthn

import (
	"github.com/trustasia-com/go-sdk/pkg/base64url"
	"github.com/trustasia-com/go-sdk/pkg/fido"
)

// StartSignUpReq sign up request
type StartSignUpReq struct {
	Username               string                                    `json:"username"`
	DisplayName            string                                    `json:"displayName"`
	Attestation            fido.AttestationConveyancePreference      `json:"attestation"`
	AuthenticatorSelection fido.AuthenticatorSelectionCriteria       `json:"authenticatorSelection"`
	Extensions             fido.AuthenticationExtensionsClientInputs `json:"extensions"`
	CredentialName         string                                    `json:"credentialName"`
}

// StartSignUpResp sign up response
// https://www.w3.org/TR/webauthn/#dictionary-makecredentialoptions
type StartSignUpResp struct {
	RP               fido.PublicKeyCredentialRpEntity     `json:"rp"`
	User             fido.PublicKeyCredentialUserEntity   `json:"user"`
	Challenge        string                               `json:"challenge"`
	PubKeyCredParams []fido.PublicKeyCredentialParameters `json:"pubKeyCredParams"`
	Timeout          int                                  `json:"timeout"`

	ExcludeCredentials     []fido.PublicKeyCredentialDescriptor      `json:"excludeCredentials,omitempty"`
	AuthenticatorSelection fido.AuthenticatorSelectionCriteria       `json:"authenticatorSelection"`
	Attestation            fido.AttestationConveyancePreference      `json:"attestation"`
	Extensions             fido.AuthenticationExtensionsClientInputs `json:"extensions"`
}

// FinishSignUpReq sign up request
type FinishSignUpReq struct {
	ID       string                       `json:"id"`
	RawID    string                       `json:"rawId"`
	Type     fido.PublicKeyCredentialType `json:"type"`
	Response struct {
		ClientDataJSON    string `json:"clientDataJSON"`
		AttestationObject string `json:"attestationObject"`
	} `json:"response"`
}

// FinishSignUpResp sign up response
type FinishSignUpResp struct{}

// StartSignInReq sign in request
type StartSignInReq struct {
	Username         string                                    `json:"username"`
	DisplayName      string                                    `json:"displayName"`
	UserVerification fido.UserVerificationRequirement          `json:"userVerification"`
	Extensions       fido.AuthenticationExtensionsClientInputs `json:"extensions"`
}

// StartSignInResp sign in response
// https://www.w3.org/TR/webauthn/#dictionary-assertion-options
type StartSignInResp struct {
	Challenge string `json:"challenge"`
	Timeout   int    `json:"timeout"`
	RpID      string `json:"rpId"`

	AllowCredentials []fido.PublicKeyCredentialDescriptor      `json:"allowCredentials,omitempty"`
	UserVerification fido.UserVerificationRequirement          `json:"userVerification,omitempty"`
	Extensions       fido.AuthenticationExtensionsClientInputs `json:"extensions"`
}

// FinishSignInReq sign in request
type FinishSignInReq struct {
	ID       string                       `json:"id"`
	RawID    string                       `json:"rawId"`
	Type     fido.PublicKeyCredentialType `json:"type"`
	Response struct {
		ClientDataJSON    string `json:"clientDataJSON"`
		AuthenticatorData string `json:"authenticatorData"`
		Signature         string `json:"signature"`
		UserHandle        string `json:"userHandle"`
	} `json:"response"`
}

// FinishSignInResp sign in response
type FinishSignInResp struct {
	UserID base64url.Encoding `json:"userId"`
}
