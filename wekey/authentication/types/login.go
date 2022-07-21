package types

import "time"

// LoginByPasswordRequest 密码登录请求
type LoginByPasswordRequest struct {
	Target   string `json:"target"`    // 账号，用户名/邮箱/手机号
	Password string `json:"password"`  // 密码
	ClientID string `json:"client_id"` // client id 不用传此参数
}

// SendPasscodeRequest 发送验证码通用接口
type SendPasscodeRequest struct {
	Target string `json:"target" binding:"required"` // 邮箱或者手机号
	Action string `json:"action" binding:"required"` // 操作，目前是login register reset-password
}

// LoginByPasscodeRequest 验证码登录
type LoginByPasscodeRequest struct {
	Target   string `json:"target"`    // 邮箱/手机号
	Passcode string `json:"passcode"`  // 验证码
	ClientID string `json:"client_id"` // client id 不用传此参数
}

// LoginUserInfo 登录响应
type LoginUserInfo struct {
	ID                  string    `json:"id"`                    // 用户id
	Username            string    `json:"username"`              // 用户名
	Email               string    `json:"email"`                 // 邮箱
	EmailVerified       bool      `json:"email_verified"`        // 邮箱验证状态
	PhoneNumber         string    `json:"phone_number"`          // 手机号
	PhoneNumberVerified bool      `json:"phone_number_verified"` // 手机号验证状态
	Status              int       `json:"status"`                // 状态0为未激活，1为已激活
	BlockStatus         int       `json:"block_status"`          // 锁定状态
	Name                string    `json:"name"`                  // name
	GivenName           string    `json:"given_name"`
	FamilyName          string    `json:"family_name"`
	Nickname            string    `json:"nickname"` // 昵称
	Profile             string    `json:"profile"`  // 个人主页
	Picture             string    `json:"picture"`  // 头像
	Website             string    `json:"website"`
	Gender              int       `json:"gender"`
	Birthdate           string    `json:"birthdate"`
	Zoneinfo            string    `json:"zoneinfo"`
	Locale              string    `json:"locale"`
	Address             string    `json:"address"`
	LastLoginIp         string    `json:"last_login_ip"`   // 上次登录ip
	LastLoginTime       time.Time `json:"last_login_time"` // 上次登录时间
	LoginCount          int       `json:"login_count"`     // 登录次数
	UpdatedAt           time.Time `json:"updated_at"`      // 信息更新时间
	CreatedAt           time.Time `json:"created_at"`      // 注册时间
}

// LoginResponse 登录返回
type LoginResponse struct {
	UserInfo LoginUserInfo `json:"user_info"` // 用户信息
	Token    string        `json:"token"`     // token
}
