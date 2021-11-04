// Package cose provides ...
package cose

import (
	"crypto"
	"crypto/x509"

	"github.com/pkg/errors"
)

//
// https://www.iana.org/assignments/cose/cose.xhtm
//

// some error
var (
	ErrUnsupportedKey          = errors.New("Unsupported Public Key Type")
	ErrUnsupportedAlgorithm    = errors.New("Unsupported public key algorithm")
	ErrSigNotProvidedOrInvalid = errors.New("Signature invalid or not provided")
)

// AlgorithmIdentifier From §5.10.5. A number identifying a cryptographic algorithm. The algorithm
// identifiers SHOULD be values registered in the IANA COSE Algorithms registry
// [https://www.w3.org/TR/webauthn/#biblio-iana-cose-algs-reg], for instance, -7 for "ES256"
//  and -257 for "RS256".
type AlgorithmIdentifier int

const (
	// AlgUnknown unknown algorithm
	AlgUnknown AlgorithmIdentifier = 0

	// AlgES256 ECDSA with SHA-256
	AlgES256 AlgorithmIdentifier = -7
	// AlgES384 ECDSA with SHA-384
	AlgES384 AlgorithmIdentifier = -35
	// AlgES512 ECDSA with SHA-512
	AlgES512 AlgorithmIdentifier = -36
	// AlgRS1 RSASSA-PKCS1-v1_5 with SHA-1
	AlgRS1 AlgorithmIdentifier = -65535
	// AlgRS256 RSASSA-PKCS1-v1_5 with SHA-256
	AlgRS256 AlgorithmIdentifier = -257
	// AlgRS384 RSASSA-PKCS1-v1_5 with SHA-384
	AlgRS384 AlgorithmIdentifier = -258
	// AlgRS512 RSASSA-PKCS1-v1_5 with SHA-512
	AlgRS512 AlgorithmIdentifier = -259
	// AlgPS256 RSASSA-PSS with SHA-256
	AlgPS256 AlgorithmIdentifier = -37
	// AlgPS384 RSASSA-PSS with SHA-384
	AlgPS384 AlgorithmIdentifier = -38
	// AlgPS512 RSASSA-PSS with SHA-512
	AlgPS512 AlgorithmIdentifier = -39
	// AlgEdDSA EdDSA
	AlgEdDSA AlgorithmIdentifier = -8
)

// EllipticCurve curve value
type EllipticCurve int64

const (
	// CurveP256 NIST P-256 also known as secp256r1
	CurveP256 EllipticCurve = 1
	// CurveP384 NIST P-384 also known as secp384r1
	CurveP384 EllipticCurve = 2
	// CurveP521 NIST P-521 also known as secp521r1
	CurveP521 EllipticCurve = 3
	// CurveX25519 X25519 for use w/ ECDH only
	CurveX25519 EllipticCurve = 4
	// CurveX448 X448 for use w/ ECDH only
	CurveX448 EllipticCurve = 5
	// CurveEd25519 Ed25519 for use w/ EdDSA only
	CurveEd25519 EllipticCurve = 6
	// CurveEd448 Ed448 for use w/ EdDSA only
	CurveEd448 EllipticCurve = 7
	// CurveSecp256k1 SECG secp256k1 curve
	CurveSecp256k1 EllipticCurve = 8
)

// PublicKeyData The public key portion of a Relying Party-specific credential key pair, generated
// by an authenticator and returned to a Relying Party at registration time. We unpack this object
// using fxamacker's cbor library ("github.com/fxamacker/cbor/v2") which is why there are cbor tags
// included. The tag field values correspond to the IANA COSE keys that give their respective
// values.
// See §6.4.1.1 https://www.w3.org/TR/webauthn/#sctn-encoded-credPubKey-examples for examples of this
// COSE data.
type PublicKeyData struct {
	// The type of key created. Should be OKP, EC2, or RSA.
	KeyType KeyType `cbor:"1,keyasint,omitempty" json:"kty"`
	// A COSEAlgorithmIdentifier for the algorithm used to derive the key signature.
	Algorithm AlgorithmIdentifier `cbor:"3,keyasint,omitempty" json:"alg"`
}

// KeyType The Key Type derived from the IANA COSE AuthData
type KeyType int64

const (
	// OctetKey is an Octet Key
	OctetKey KeyType = 1
	// EllipticKey is an Elliptic Curve Public Key
	EllipticKey KeyType = 2
	// RSAKey is an RSA Public Key
	RSAKey KeyType = 3
)

// SignatureAlgorithmDetails signature with algo, coseAlg, name, hasher
var SignatureAlgorithmDetails = []struct {
	algo    x509.SignatureAlgorithm
	coseAlg AlgorithmIdentifier
	name    string
	hasher  crypto.Hash
}{
	{x509.SHA1WithRSA, AlgRS1, "SHA1-RSA", crypto.SHA1},
	{x509.SHA256WithRSA, AlgRS256, "SHA256-RSA", crypto.SHA256},
	{x509.SHA384WithRSA, AlgRS384, "SHA384-RSA", crypto.SHA384},
	{x509.SHA512WithRSA, AlgRS512, "SHA512-RSA", crypto.SHA512},
	{x509.SHA256WithRSAPSS, AlgPS256, "SHA256-RSAPSS", crypto.SHA256},
	{x509.SHA384WithRSAPSS, AlgPS384, "SHA384-RSAPSS", crypto.SHA384},
	{x509.SHA512WithRSAPSS, AlgPS512, "SHA512-RSAPSS", crypto.SHA512},
	{x509.ECDSAWithSHA256, AlgES256, "ECDSA-SHA256", crypto.SHA256},
	{x509.ECDSAWithSHA384, AlgES384, "ECDSA-SHA384", crypto.SHA384},
	{x509.ECDSAWithSHA512, AlgES512, "ECDSA-SHA512", crypto.SHA512},
	{x509.PureEd25519, AlgEdDSA, "EdDSA", crypto.SHA512},
}
