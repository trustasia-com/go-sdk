package types

// BaseResponse base
type BaseResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

// GetPublicKeyResponse public resp
type GetPublicKeyResponse struct {
	BaseResponse
	Data string `json:"data"`
}
