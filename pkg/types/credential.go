// Package types provides ...
package types

import (
	"github.com/trustasia-com/go-sdk/pkg/base64url"
	"github.com/trustasia-com/go-sdk/pkg/cose"
)

// PublicKeyCredentialEntity describes a user account, or a WebAuthn Relying Party, which a public key
// credential is associated with or scoped to, respectively.
// https://www.w3.org/TR/webauthn/#dictionary-pkcredentialentity
type PublicKeyCredentialEntity struct {
	// A human-palatable name for the entity. Its function depends on what the PublicKeyCredentialEntity
	// represents:
	//
	// When inherited by PublicKeyCredentialRpEntity it is a human-palatable identifier for the Relying
	// Party,
	// intended only for display. For example, "ACME Corporation", "Wonderful Widgets, Inc." or
	// "ОАО Примертех".
	//
	// When inherited by PublicKeyCredentialUserEntity, it is a human-palatable identifier for a user
	// account. It is
	// intended only for display, i.e., aiding the user in determining the difference between user
	// accounts with similar
	// displayNames. For example, "alexm", "alex.p.mueller@example.com" or "+14255551234".
	Name string `json:"name"`
}

// PublicKeyCredentialRpEntity is used to supply additional Relying Party attributes when creating a new
// credential.
// https://www.w3.org/TR/webauthn/#dictionary-rp-credential-params
type PublicKeyCredentialRpEntity struct {
	PublicKeyCredentialEntity

	// A unique identifier for the Relying Party entity, which sets the RP ID
	ID string `json:"id"`

	// Deprecated: new WebAuthn, rp icon
	Icon string `json:"icon,omitempty"`
}

// PublicKeyCredentialUserEntity is used to supply additional user account attributes when creating a
// new credential.
// https://www.w3.org/TR/webauthn/#dictionary-user-credential-params
type PublicKeyCredentialUserEntity struct {
	PublicKeyCredentialEntity

	// A human-palatable name for the user account, intended only for display.
	// For example, "Alex P. Müller" or "田中 倫". The Relying Party SHOULD let
	// the user choose this, and SHOULD NOT restrict the choice more than necessary.
	DisplayName string `json:"displayName"`
	// ID is the user handle of the user account entity. To ensure secure operation,
	// authentication and authorization decisions MUST be made on the basis of this id
	// member, not the displayName nor name members. See Section 6.1 of
	// [RFC8266](https://www.w3.org/TR/webauthn/#biblio-rfc8266).
	ID base64url.Encoding `json:"id"` // max 64 byte

	// Deprecated: new WebAuthn, user icon
	Icon string `json:"icon,omitempty"`
}

// PublicKeyCredentialType This enumeration defines the valid credential types. It is an extension point;
// values can be added to it in the future, as more credential types are defined.
// https://www.w3.org/TR/webauthn/#enumdef-publickeycredentialtype
type PublicKeyCredentialType string

const (
	// CredentialTypePublicKey public key
	CredentialTypePublicKey PublicKeyCredentialType = "public-key"
)

// PublicKeyCredentialParameters used to supply additional parameters when creating a new credential.
// https://www.w3.org/TR/webauthn/#dictionary-credential-params
type PublicKeyCredentialParameters struct {
	// This member specifies the type of credential to be created.
	Type PublicKeyCredentialType `json:"type"`
	// This member specifies the cryptographic signature algorithm with which the newly generated
	// credential will be used, and thus also the type of asymmetric key pair to be generated
	Alg cose.AlgorithmIdentifier `json:"alg"`
}

// PublicKeyCredentialDescriptor This dictionary contains the attributes that are specified by a caller
// when referring to a public key credential as an input parameter to the create() or get() methods.
// It mirrors the fields of the PublicKeyCredential object returned by the latter methods.
// https://www.w3.org/TR/webauthn/#dictionary-credential-descriptor
type PublicKeyCredentialDescriptor struct {
	// the type of the public key credential the caller is referring to.
	Type PublicKeyCredentialType `json:"type"`
	// the credential ID of the public key credential the caller is referring to.
	ID base64url.Encoding `json:"id"`
	// This OPTIONAL member contains a hint as to how the client might communicate with the managing
	// authenticator of the public key credential the caller is referring to
	Transports []AuthenticatorTransport `json:"transports,omitempty"`
}
