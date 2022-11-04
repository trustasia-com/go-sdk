// Package message provides ...
package message

import (
	"encoding/json"

	"github.com/trustasia-com/go-van/pkg/codes"
)

// JSONRawMessage request or response data format
type JSONRawMessage struct {
	Code  codes.Code      `json:"code"`
	Data  json.RawMessage `json:"data"`
	Error string          `json:"error"`
}
