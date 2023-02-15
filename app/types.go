// Package app provides ...
package app

import "time"

// RpInfo rp info
type RpInfo struct {
	RpID        string   `json:"rp_id"`         // 租户唯一, sso.example.com
	Origins     []string `json:"origins"`       // eg. https://sso.example.com
	Icon        string   `json:"icon"`          // eg. url, 过时
	RpPolicy    int      `json:"rp_pilicy"`     // rp严格校验模式
	RpAllowList []string `json:"rp_allow_list"` // 允许列表
}

// CreateOrUpdateAppReq 创建请求
type CreateOrUpdateAppReq struct {
	Slug    string `json:"slug"`
	Name    string `json:"name"`
	ExLogin bool   `json:"ex_login"`

	RpInfo RpInfo `json:"rp_info"`
}

// DeleteAppReq 删除请求
type DeleteAppReq struct {
	Slug string `json:"slug"`
}

/////////////////////////////////////

// RegQRCodeReq 注册请求
type RegQRCodeReq struct {
	Slug           string `json:"slug"`
	CredentialName string `json:"credential_name"`

	RpUserID   string `json:"rp_user_id"`
	RpUsername string `json:"rp_username"`
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
	Username string `json:"username"`
	Status   string `json:"status"`
	Error    string `json:"error"`
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
	Slug   string     `json:"slug"`
	Method AuthMethod `json:"method"`

	RpUserID   string `json:"rp_user_id"`
	RpUsername string `json:"rp_username"`
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
	Username string `json:"username"`
	Status   string `json:"status"`
	Error    string `json:"error"`
	UserID   string `json:"user_id"`
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
