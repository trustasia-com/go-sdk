// Package types provides ...
package types

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
