package types

import "time"

// ZoneInfoBasicResp req
type ZoneInfoBasicResp struct {
	Name         string `json:"name"`          // 别名
	Slug         string `json:"slug"`          // slug
	Logo         string `json:"logo"`          // logo
	EmailSupport string `json:"email_support"` // 邮件支持
	SupportUrl   string `json:"support_url"`   // 支持URL
	Description  string `json:"description"`
}

// ZoneInfoMemberResp 空间成员
type ZoneInfoMemberResp struct {
	UserID         string   `json:"user_id"`         // 用户id
	IsCreator      bool     `json:"is_creator"`      // 是否是创建者，如果是创建者，不允许对该成员做删除操作
	Email          string   `json:"email"`           // 邮箱
	Username       string   `json:"username"`        // 用户名
	Nickname       string   `json:"nickname"`        // 用户昵称
	Picture        string   `json:"picture"`         // 用户头像
	Roles          []string `json:"roles"`           // 用户角色 admin editor:apps editor:identity editor:users viewer:users viewer:config
	WebauthnEnable bool     `json:"webauthn_enable"` // mfa 这个只要和otp_enable有一个为true就表示开启了
	OtpEnable      bool     `json:"otp_enable"`      // mfa 这个只要和webauthn_enable有一个为true就表示开启了
}

// ZoneRevokeSecretItem 吊销的secret
type ZoneRevokeSecretItem struct {
	SecretID  string    `json:"secret_id"`  // 密钥id
	CreatedAt time.Time `json:"created_at"` // 吊销时间
}

// ZoneInfoSecrets 空间密钥响应
type ZoneInfoSecrets struct {
	CertID          string                 `json:"cert_id"`            // 证书sha1,也就是证书id
	CertPem         string                 `json:"cert_pem"`           // 证书
	CertPemSHA1     string                 `json:"cert_pem_sha1"`      // 证书sha1
	NextCertID      string                 `json:"next_cert_id"`       // 下一个证书sha1，也就是证书id
	NextCertPem     string                 `json:"next_cert_pem"`      // 下一个证书
	NextCertPemSHA1 string                 `json:"next_cert_pem_sha1"` // 下一个证书sha1
	RevokeSecrets   []ZoneRevokeSecretItem `json:"revoke_secrets"`     // 吊销的证书
}

// ZoneInfoAdvanceResp 空间高级设置
type ZoneInfoAdvanceResp struct {
	CookieMode        string `json:"cookie_mode"`        // cookie模式 persistent持续 non-persistent非持久
	InactivityTimeout int    `json:"inactivity_timeout"` // 不活动超时,单位秒
	FinalTimeout      int    `json:"final_timeout"`      // 最终超时,单位秒
}

// UpdateZoneBasicInfoReq req
type UpdateZoneBasicInfoReq struct {
	Name         string `json:"name"`          // 必填 别名
	Description  string `json:"description"`   // 选填 最长144个字符
	Logo         string `json:"logo"`          // 必填 logo图
	EmailSupport string `json:"email_support"` // 选填 邮件支持地址
	SupportUrl   string `json:"support_url"`   // 选填 支持URL
}

// UpdateZoneMembersReq  编辑空间成员
type UpdateZoneMembersReq struct {
	Roles []string `json:"roles"` // 必传 用户角色 admin editor:apps editor:identity editor:users viewer:users viewer:config
}

// UpdateZoneAdvanceReq  update zone advance
type UpdateZoneAdvanceReq struct {
	CookieMode        string `json:"cookie_mode"`        // 必填 cookie模式 persistent持续 non-persistent非持久
	InactivityTimeout int    `json:"inactivity_timeout"` // 必填 不活动超时,单位秒，最小60 ，最大259200
	FinalTimeout      int    `json:"final_timeout" `     // 必填 最终超时,单位秒 最小60，最大2592000
}
