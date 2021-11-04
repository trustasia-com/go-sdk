// Package types provides ...
package types

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
	Username               string                               `json:"username"`
	DisplayName            string                               `json:"displayName"`
	Attestation            AttestationConveyancePreference      `json:"attestation"`
	AuthenticatorSelection AuthenticatorSelectionCriteria       `json:"authenticatorSelection"`
	Extensions             AuthenticationExtensionsClientInputs `json:"extensions"`
}

// StartSignUpResp sign up response
type StartSignUpResp struct {
	RP                     PublicKeyCredentialRpEntity          `json:"rp"`
	User                   PublicKeyCredentialUserEntity        `json:"user"`
	Challenge              string                               `json:"challenge"`
	PubKeyCredParams       []PublicKeyCredentialType            `json:"pubKeyCredParams"`
	Timeout                int                                  `json:"timeout"`
	ExcludeCredentials     []PublicKeyCredentialDescriptor      `json:"excludeCredentials"`
	AuthenticatorSelection AuthenticatorSelectionCriteria       `json:"authenticatorSelection"`
	Attestation            AttestationConveyancePreference      `json:"attestation"`
	Extensions             AuthenticationExtensionsClientInputs `json:"extensions"`
}

// FinishSignUpReq sign up request
type FinishSignUpReq struct {
	ID       string                  `json:"id"`
	RawID    string                  `json:"rawId"`
	Type     PublicKeyCredentialType `json:"type"`
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
