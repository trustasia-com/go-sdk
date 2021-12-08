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
