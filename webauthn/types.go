// Package webauthn provides ...
package webauthn

import "github.com/trustasia-com/go-sdk/pkg/types"

// StartSignUpReq sign up request
type StartSignUpReq struct {
	Username               string                                     `json:"username"`
	DisplayName            string                                     `json:"displayName"`
	Attestation            types.AttestationConveyancePreference      `json:"attestation"`
	AuthenticatorSelection types.AuthenticatorSelectionCriteria       `json:"authenticatorSelection"`
	Extensions             types.AuthenticationExtensionsClientInputs `json:"extensions"`
}

// StartSignUpResp sign up response
// https://www.w3.org/TR/webauthn/#dictionary-makecredentialoptions
type StartSignUpResp struct {
	RP               types.PublicKeyCredentialRpEntity     `json:"rp"`
	User             types.PublicKeyCredentialUserEntity   `json:"user"`
	Challenge        string                                `json:"challenge"`
	PubKeyCredParams []types.PublicKeyCredentialParameters `json:"pubKeyCredParams"`
	Timeout          int                                   `json:"timeout"`

	ExcludeCredentials     []types.PublicKeyCredentialDescriptor      `json:"excludeCredentials"`
	AuthenticatorSelection types.AuthenticatorSelectionCriteria       `json:"authenticatorSelection"`
	Attestation            types.AttestationConveyancePreference      `json:"attestation"`
	Extensions             types.AuthenticationExtensionsClientInputs `json:"extensions"`
}

// FinishSignUpReq sign up request
type FinishSignUpReq struct {
	Name string `json:"name"` // 凭证/设备名称

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
type StartSignInReq struct {
	Username         string                                     `json:"username"`
	DisplayName      string                                     `json:"displayName"`
	UserVerification types.UserVerificationRequirement          `json:"userVerification"`
	Extensions       types.AuthenticationExtensionsClientInputs `json:"extensions"`
}

// StartSignInResp sign in response
// https://www.w3.org/TR/webauthn/#dictionary-assertion-options
type StartSignInResp struct {
	Challenge string `json:"challenge"`
	Timeout   int    `json:"timeout"`
	RpID      string `json:"rpId"`

	AllowCredentials []types.PublicKeyCredentialDescriptor      `json:"allowCredentials"`
	UserVerification types.UserVerificationRequirement          `json:"userVerification,omitempty"`
	Extensions       types.AuthenticationExtensionsClientInputs `json:"extensions"`
}

// FinishSignInReq sign in request
type FinishSignInReq struct {
	ID       string                        `json:"id"`
	RawID    string                        `json:"rawId"`
	Type     types.PublicKeyCredentialType `json:"type"`
	Response struct {
		ClientDataJSON    string `json:"clientDataJSON"`
		AuthenticatorData string `json:"authenticatorData"`
		Signature         string `json:"signature"`
		UserHandle        string `json:"userHandle"`
	} `json:"response"`
}

// FinishSignInResp sign in response
type FinishSignInResp struct {
	UserID []byte `json:"userId"`
}
