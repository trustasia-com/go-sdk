// Package message provides ...
package message

import "encoding/json"

// JSONRawMessage request or response data format
type JSONRawMessage struct {
	Code  int             `json:"code"`
	Data  json.RawMessage `json:"data"`
	Error string          `json:"error"`
}
