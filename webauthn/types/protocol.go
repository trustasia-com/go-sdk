// Package types provides ...
package types

import (
	"github.com/trustasia-com/go-sdk/internal/base64url"
	"github.com/trustasia-com/go-sdk/internal/cose"
)

// AuthenticatorAttachment This enumeration’s values describe authenticators' attachment modalities.
// Relying Parties use this to express a preferred authenticator attachment modality when calling
// navigator.credentials.create() to create a credential.
// https://www.w3.org/TR/webauthn/#enum-attachment
type AuthenticatorAttachment string

const (
	// Platform - A platform authenticator is attached using a client device-specific transport, called
	// platform attachment, and is usually not removable from the client device. A public key credential
	//  bound to a platform authenticator is called a platform credential.
	Platform AuthenticatorAttachment = "platform"
	// CrossPlatform - A roaming authenticator is attached using cross-platform transports, called
	// cross-platform attachment. Authenticators of this class are removable from, and can "roam"
	// among, client devices. A public key credential bound to a roaming authenticator is called a
	// roaming credential.
	CrossPlatform AuthenticatorAttachment = "cross-platform"
)

// ResidentKeyRequirement this enumeration’s values describe the relying party's requirements for
// client-side discoverable credentials (formerly known as resident credentials or resident keys
// https://www.w3.org/TR/webauthn/#enum-residentKeyRequirement
type ResidentKeyRequirement string

const (
	// ResidentKeyDiscouraged - This value indicates the Relying Party requires a client-side
	// discoverable credential, and is prepared to receive an error if a client-side discoverable
	// credential cannot be created.
	ResidentKeyDiscouraged ResidentKeyRequirement = "discouraged"
	// ResidentKeyPreferred - This value indicates the Relying Party strongly prefers creating a
	// client-side discoverable credential, but will accept a server-side credential.
	// NOTE This is the default
	ResidentKeyPreferred ResidentKeyRequirement = "preferred"
	// ResidentKeyRequired - this value indicates the relying party prefers creating a server-side
	// credential, but will accept a client-side discoverable credential.
	ResidentKeyRequired ResidentKeyRequirement = "required"
)

// UserVerificationRequirement A WebAuthn Relying Party may require user verification for some of its
// operations but not for others, and may use this type to express its needs.
// https://www.w3.org/TR/webauthn/#enum-userVerificationRequirement
type UserVerificationRequirement string

const (
	// VerificationRequired User verification is required to create/release a credential
	VerificationRequired UserVerificationRequirement = "required"
	// VerificationPreferred User verification is preferred to create/release a credential
	// NOTE This is the default
	VerificationPreferred UserVerificationRequirement = "preferred"
	// VerificationDiscouraged The authenticator should not verify the user for the credential
	VerificationDiscouraged UserVerificationRequirement = "discouraged"
)

// AuthenticatorSelectionCriteria WebAuthn Relying Parties may use the AuthenticatorSelectionCriteria
// dictionary to specify their requirements regarding authenticator attributes.
// https://www.w3.org/TR/webauthn/#dictionary-authenticatorSelection
type AuthenticatorSelectionCriteria struct {
	// AuthenticatorAttachment If this member is present, eligible authenticators are filtered to only
	// authenticators attached with the specified AuthenticatorAttachment enum
	AuthenticatorAttachment AuthenticatorAttachment `json:"authenticatorAttachment,omitempty"`
	// Specifies the extent to which the Relying Party desires to create a client-side discoverable
	// credential.
	// For historical reasons the naming retains the deprecated “resident” terminology.
	ResidentKey ResidentKeyRequirement `json:"residentKey,omitempty"`
	// RequireResidentKey this member describes the Relying Party's requirements regarding resident
	// credentials. If the parameter is set to true, the authenticator MUST create a client-side-resident
	// public key credential source when creating a public key credential.
	RequireResidentKey *bool `json:"requireResidentKey,omitempty"`
	// UserVerification This member describes the Relying Party's requirements regarding user verification
	// for the create() operation. Eligible authenticators are filtered to only those capable of
	// satisfying this requirement.
	UserVerification UserVerificationRequirement `json:"userVerification,omitempty"`
}

// AttestationConveyancePreference WebAuthn Relying Parties may use AttestationConveyancePreference to
// specify their preference regarding attestation conveyance during credential generation.
// https://www.w3.org/TR/webauthn/#enum-attestation-convey
type AttestationConveyancePreference string

const (
	// PreferenceNone - This value indicates that the Relying Party is not interested in authenticator
	// attestation.
	PreferenceNone AttestationConveyancePreference = "none" // This is the default
	// PreferenceIndirect - This value indicates that the Relying Party prefers an attestation conveyance
	// yielding verifiable attestation statements, but allows the client to decide how to obtain such
	// attestation statements.
	PreferenceIndirect AttestationConveyancePreference = "indirect"
	// PreferenceDirect - This value indicates that the Relying Party wants to receive the attestation
	// statement as generated by the authenticator.
	PreferenceDirect AttestationConveyancePreference = "direct"
	// PreferenceEnterprise - This value indicates that the Relying Party wants to receive an attestation
	// statement that may include uniquely identifying information.
	PreferenceEnterprise AttestationConveyancePreference = "enterprise"
)

// AuthenticationExtensionsClientInputs An extension defines one or two request arguments. The client
// extension input, which is a value that can be encoded in JSON, is passed from the WebAuthn Relying Party
// to the client in the get() or create() call, while the CBOR authenticator extension input is passed
// from the client to the authenticator for authenticator extensions during the processing of these calls.
// https://www.w3.org/TR/webauthn/#iface-authentication-extensions-client-inputs
//
// it's map: extension identifier -> client extension input
type AuthenticationExtensionsClientInputs map[string]interface{}

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
	ID []byte `json:"id"` // max 64 byte

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

// AuthenticatorTransport Authenticators may implement various transports for communicating with clients.
// This enumeration defines hints as to how clients might communicate with a particular authenticator in
// order to obtain an assertion for a specific credential. Note that these hints represent the WebAuthn
// Relying Party's best belief as to how an authenticator may be reached. A Relying Party will typically
// learn of the supported transports for a public key credential via getTransports().
// https://www.w3.org/TR/webauthn/#enumdef-authenticatortransport
type AuthenticatorTransport string

const (
	// TransportUSB The authenticator should transport information over USB
	TransportUSB AuthenticatorTransport = "usb"
	// TransportNFC The authenticator should transport information over Near Field Communication Protocol
	TransportNFC AuthenticatorTransport = "nfc"
	// TransportBLE The authenticator should transport information over Bluetooth
	TransportBLE AuthenticatorTransport = "ble"
	// TransportInternal the client should use an internal source like a TPM or SE
	TransportInternal AuthenticatorTransport = "internal"
	// TransportMQ transport with mq
	TransportMQ AuthenticatorTransport = "mq"
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
