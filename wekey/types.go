// Package wekey provides ...
package wekey

import "time"

// RegQRCodeReq 注册请求
type RegQRCodeReq struct {
	UserID      string `json:"user_id"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
}

// RegQRCodeResp 响应请求
type RegQRCodeResp struct {
	URL       string `json:"url"`
	ExpiresAt int64  `json:"expires_at"`
}

// RegResultReq 注册结果获取
type RegResultReq struct {
	MsgID string `json:"msg_id"`
}

// RegResultResp 注册结果响应
type RegResultResp struct {
	WekeyUser string `json:"wekey_user"`
	Status    string `json:"status"`
	Error     string `json:"error"`
}

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
	URL       string `json:"url"`
	ExpiresAt int64  `json:"expires_at"`
}

// AuthResultReq 认证结果
type AuthResultReq struct {
	MsgID string `json:"msg_id"`
}

// AuthResultResp 认证结果响应
type AuthResultResp struct {
	WekeyUser string `json:"wekey_user"`
	Status    string `json:"status"`
	Error     string `json:"error"`
	UserID    string `json:"user_id"`
}

// UserCredentialsReq 获取凭证列表请求
type UserCredentialsReq struct {
	UserID string `json:"user_id"`
}

// Credential 凭证
type Credential struct {
	CredentialID   string `json:"credential_id"`   // 凭证ID
	CredentialName string `json:"credential_name"` // 凭证名称

	UserID    string    `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"` // 最近使用
	CreatedAt time.Time `json:"created_at"` // 添加时间
}

// UserCredentialsResp 凭证列表响应
type UserCredentialsResp struct {
	List  []Credential `json:"list"`
	Total int64        `json:"total"`
}

// DeleteCredentialReq 删除用户凭证请求
type DeleteCredentialReq struct {
	UserID        string   `json:"user_id"`        // 用户ID
	CredentialIDs []string `json:"credential_ids"` // 凭证ID
}

// DeleteCredentialResp 删除响应
type DeleteCredentialResp struct{}
