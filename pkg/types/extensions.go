// Package types provides ...
package types

// AuthenticationExtensionsClientInputs An extension defines one or two request arguments. The client
// extension input, which is a value that can be encoded in JSON, is passed from the WebAuthn Relying Party
// to the client in the get() or create() call, while the CBOR authenticator extension input is passed
// from the client to the authenticator for authenticator extensions during the processing of these calls.
// https://www.w3.org/TR/webauthn/#iface-authentication-extensions-client-inputs
//
// it's map: extension identifier -> client extension input
type AuthenticationExtensionsClientInputs map[string]interface{}
