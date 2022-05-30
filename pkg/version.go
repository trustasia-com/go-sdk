// Package pkg provides ...
package pkg

import (
	"fmt"
	"runtime"
)

const (
	// SDKName is the name of this WeKey SDK
	SDKName = "wekey-sdk-go"
	// SDKVersion is the version of this SDK
	SDKVersion = "0.1.1"
)

// BuildUserAgent build the user gaent string
func BuildUserAgent() string {
	return fmt.Sprintf("%s-%s %s", SDKName, SDKVersion, SystemInfo())
}

// SystemInfo returns the system information.
func SystemInfo() string {
	return fmt.Sprintf("(arch=%s; os=%s; go.version=%s)", runtime.GOARCH,
		runtime.GOOS, runtime.Version())
}
