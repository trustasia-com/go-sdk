// Package app provides ...
package app

import (
	"github.com/trustasia-com/go-van/pkg/codes/status"
)

// IsNotFoundCredentialErr 是否是无凭证错误
func IsNotFoundCredentialErr(err error) bool {
	code := status.Code(err)
	return code == 11424
}
