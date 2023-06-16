// Package app provides ...
package app

import "time"

// AuthMethod 认证类型
type AuthMethod string

// 类型
const (
	TypeFidoScan   AuthMethod = "fido-scan"
	TypeCosignScan AuthMethod = "cosign-scan"
)

// ClientType 客户端类型
type ClientType string

func (ct ClientType) isValid() bool {
	return ct == ClientTypeWeb || ct == ClientTypeMobile
}

// 客户端类型
const (
	ClientTypeWeb    ClientType = "web"
	ClientTypeMobile ClientType = "mobile"
)

// BindQRCodeReq 注册请求
type BindQRCodeReq struct {
	RpUserDisplayName string     `json:"rp_user_display_name"`
	RpUserID          string     `json:"rp_user_id"`
	ClientType        ClientType `json:"client_type"`
}

// BindQRCodeResp 响应请求
type BindQRCodeResp struct {
	ExpiredAt int64  `json:"expired_at"` // 消息过期时间
	URL       string `json:"url"`        // 响应消息地址
	MsgID     string `json:"msg_id"`
}

// BindResultReq 注册结果获取
type BindResultReq struct {
	MsgID string `json:"msg_id"`
}

// BindResultResp 注册结果响应
type BindResultResp struct {
	Status string `json:"status"`

	ErrorMsg      string `json:"error_msg"`
	WekeyNickname string `json:"wekey_nickname"`
}

// LoginQRCodeReq 认证请求
type LoginQRCodeReq struct {
	ClientType ClientType `json:"client_type"`
}

// LoginQRCodeResp 认证请求响应
type LoginQRCodeResp struct {
	ExpiredAt int64  `json:"expired_at"` // 消息过期时间
	URL       string `json:"url"`        // 响应消息地址
	MsgID     string `json:"msg_id"`
}

// LoginResultReq 认证结果
type LoginResultReq struct {
	MsgID string `json:"msg_id"`
}

// LoginResultResp 认证结果响应
type LoginResultResp struct {
	RpUserDisplayName string `json:"rp_user_display_name"`
	RpUserID          string `json:"rp_user_id"`
	Status            string `json:"status"`
	ErrorMsg          string `json:"error_msg"`
}

// SignRequestReq 签名请求
type SignRequestReq struct {
	ClientType  ClientType `json:"client_type"`
	NeedPKCS7   bool       `json:"need_pkcs7"`
	Data        string     `json:"data"`
	CallbackURL string     `json:"callback_url"`
}

// SignRequestResp 请求响应
type SignRequestResp struct {
	ExpiredAt int64  `json:"expired_at"`
	URL       string `json:"url"`
	MsgID     string `json:"msg_id"`
}

// SignResultReq 签名结果
type SignResultReq struct {
	MsgID string `json:"msg_id"`
}

// SignResultResp 签名结果
type SignResultResp struct {
	CertPEM       string `json:"cert_pem"` // 证书pem内容
	Data          string `json:"data"`     // 原始签名数据, 16进制
	ErrorMsg      string `json:"error_msg"`
	PKCS7         string `json:"pkcs7"`     // PKCS7签名
	Signature     string `json:"signature"` // 签名数据
	Status        string `json:"status"`
	WekeyNickname string `json:"wekey_nickname"`
}

/////////////////////////////////////////////

// CredentialsReq 协同列表
type CredentialsReq struct {
	Page     int    `json:"page"`
	Size     int    `json:"size"`
	RpUserID string `json:"rp_user_id"`
}

// CredentialInfo 信息
type CredentialInfo struct {
	ID            string    `json:"id"`
	WeKeyNickname string    `json:"we_key_nickname"`
	CertPEM       string    `json:"cert_pem"`
	CreatedAt     time.Time `json:"created_at"`
}

// CredentialsResp 响应
type CredentialsResp struct {
	List  []CredentialInfo `json:"list"`
	Total int              `json:"total"`
}

// CredentialDeleteReq 删除
type CredentialDeleteReq struct {
	RpUserID string `json:"rp_user_id"`
	CertID   string `json:"cert_id"`
}

// CredentialDeleteResp 响应
type CredentialDeleteResp struct{}
