// Package wekey provides ...
package wekey

// RegQRCodeReq 注册请求
type RegQRCodeReq struct {
	UserID      string `json:"user_id"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
}

// RegQRCodeResp 响应请求
type RegQRCodeResp struct {
	QRCode  string `json:"qr_code"`
	Timeout int    `json:"timeout"`
}

// RegResultReq 注册结果获取
type RegResultReq struct {
	MsgID string `json:"msg_id"`
}

// RegResultResp 注册结果响应
type RegResultResp struct{}

// AuthMethod method
type AuthMethod string

// authmethod list
const (
	AuthMethodQRCode AuthMethod = "qrcode"
	AuthMethodPush   AuthMethod = "push"
)

// AuthRequestReq 认证请求
type AuthRequestReq struct {
	Method   AuthMethod `json:"method"`
	UserID   string     `json:"user_id"`
	Username string     `json:"username"`
}

// AuthRequestResp 认证请求响应
type AuthRequestResp struct {
}

// AuthResultReq 认证结果
type AuthResultReq struct {
	MsgID string `json:"msg_id"`
}

// AuthResultResp 认证结果响应
type AuthResultResp struct{}
