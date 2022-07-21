package types

import "time"

// QueryUserReq 获取用户请求
type QueryUserReq struct {
	BaseQueryRequest

	Search string `form:"search"` // 查询条件，支持username/phone/email的模糊查询
}

// UserListRespItem 用户列表item
type UserListRespItem struct {
	ID            string    `json:"id"`
	Username      string    `json:"username"`
	PhoneNumber   string    `json:"phone_number"`
	Email         string    `json:"email"`
	Picture       string    `json:"picture"`         // 头像
	Status        int       `json:"status"`          // 激活状态，0未激活，1已激活
	BlockStatus   int       `json:"block_status"`    // 锁定状态，0正常未锁定，1已锁定
	LastLoginTime time.Time `json:"last_login_time"` // 上次登录时间
	LoginCount    int       `json:"login_count"`     // 登录次数
	IdpID         int       `json:"idp_id"`          //
	IdpProvider   string    `json:"idp_provider"`
	IdpIdentifier string    `json:"idp_identifier"`
	IdpName       string    `json:"idp_name"`
}

// UserListResp 用户列表响应
type UserListResp struct {
	Total int                `json:"total"` // 总数据
	List  []UserListRespItem `json:"list"`  // 数据列表
}

// UserInfoIdpItems 身份源item
type UserInfoIdpItems struct {
	Name         string `json:"name"`           // 身份源名称
	Provider     string `json:"provider"`       // 身份源
	Identifier   string `json:"identifier"`     // 唯一标识
	IsPrimaryIdp bool   `json:"is_primary_idp"` // 是否是主要身份源
}

// UserDetailResp 用户详细
type UserDetailResp struct {
	// 基础信息
	ID                  string    `json:"id"`                    // 用户id
	Username            string    `json:"username"`              // 用户名
	Picture             string    `json:"picture"`               // 头像
	Nickname            string    `json:"nickname"`              // 昵称
	Email               string    `json:"email"`                 // 邮箱
	EmailVerified       bool      `json:"email_verified"`        // 邮箱是否验证
	PhoneNumber         string    `json:"phone_number"`          // 手机号
	PhoneNumberVerified bool      `json:"phone_number_verified"` // 手机号是否验证
	LastLoginIp         string    `json:"last_login_ip"`         // 上次登录ip
	LastLoginTime       time.Time `json:"last_login_time"`       // 上次登录时间
	LoginCount          int       `json:"login_count"`           // 登录总次数
	CreatedAt           time.Time `json:"created_at"`            // 注册时间
	UserAgent           string    `json:"user_agent"`            // 浏览器
	// MFA
	WebauthnEnable bool `json:"webauthn_enable"` // 多因子，webauthn开关
	OtpEnable      bool `json:"otp_enable"`      // 多因子，otp开关

	// metadata
	UserMetadata string `json:"user_metadata"` // 用户metadata
	AppMetadata  string `json:"app_metadata"`  // 应用metadata
	// idp
	Identities []UserInfoIdpItems `json:"identities"` // 身份源信息

	// other
	Status      int       `json:"status"`       // 账号状态，0代表未初始化，1已初始化
	BlockStatus int       `json:"block_status"` // 锁定状态，0正常，1已锁定
	Name        string    `json:"name"`         // 名字
	GivenName   string    `json:"given_name"`   // 名
	FamilyName  string    `json:"family_name"`  // 姓
	Profile     string    `json:"profile"`      // 个人信息主页
	Website     string    `json:"website"`      // 个人blog，官网等
	Gender      int       `json:"gender"`       // 性别,0未设置，1男，2女
	Birthdate   time.Time `json:"birthdate"`    // 生日
	TimeZone    string    `json:"time_zone"`    // 时区
	Locale      string    `json:"locale"`       // 区域
	Address     string    `json:"address"`      // 地址
}

// CreateUserReq create user
type CreateUserReq struct {
	Username    string `json:"username" binding:"max=64"`   // 选填 用户名 最长64个字符
	PhoneNumber string `json:"phone_number"`                // 选填 手机号
	Email       string `json:"email"`                       // 选填 邮箱
	Password    string `json:"password" binding:"required"` // 必填 密码
	// connections
	ConnectionID string `json:"connection_id"` // 选填 身份源id
}

// UpdateUserBasicRequest 更新用户基础信息
type UpdateUserBasicRequest struct {
	Email       string `json:"email"`        // 选填 邮箱地址 最长64个字符
	PhoneNumber string `json:"phone_number"` // 选填 电话号码 11位
	Nickname    string `json:"nickname"`     // 选填 昵称 最长128个字符
}

// UpdateUserPictureRequest  更新用户头像信息
type UpdateUserPictureRequest struct {
	Picture string `json:"picture"` // 必填 头像 url
}

// UpdateUserPwdRequest req
type UpdateUserPwdRequest struct {
	Password string `json:"password" binding:"required"` // 密码,最少6位，后续会修改位加密传输
}

// UpdateUserMetadataRequest req
type UpdateUserMetadataRequest struct {
	UserMetadata interface{} `json:"user_metadata" binding:"required"` // 用户metadata
	AppMetadata  interface{} `json:"app_metadata" binding:"required"`  // 应用metadata
}

// UserPermissionInfo 权限信息
type UserPermissionInfo struct {
	PermissionID string    `json:"permission_id"` // 权限id
	Scope        string    `json:"scope"`         // scope
	Description  string    `json:"description"`   // 描述
	ResourceID   string    `json:"resource_id"`   // 资源组id
	ResourceName string    `json:"resource_name"` // 资源组名字
	CreatedAt    time.Time `json:"created_at"`
}

// UserWithoutPermissionRequest query
type UserWithoutPermissionRequest struct {
	ResourceID string `form:"resource_id"` // 必填 资源组id
	Scope      string `form:"scope"`       // 选填 权限
}

// UserWithoutPermissionResp item
type UserWithoutPermissionResp struct {
	ID          string `json:"id"`          // 权限id
	Scope       string `json:"scope"`       // 权限
	Description string `json:"description"` // 描述
}

// UserRoleInfoResp  用户角色信息
type UserRoleInfoResp struct {
	RoleID      string    `json:"role_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// UserWithoutRoleRequest QueryRoleRequest query
type UserWithoutRoleRequest struct {
	Name string `form:"name"` // 选填 角色名
}

// UserQueryRoleResp  resp
type UserQueryRoleResp struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`        // 名字
	Description string    `json:"description"` // 描述
	CreatedAt   time.Time `json:"created_at"`  // 创建时间
}

// UserAppInfoResp  用户应用信息
type UserAppInfoResp struct {
	AppID    string `json:"app_id"`    // app id
	ClientID string `json:"client_id"` // client id
	Name     string `json:"name"`      // 应用名字
	Logo     string `json:"logo"`      // logo
}

// UserAssignPermissionReq assign permission
type UserAssignPermissionReq struct {
	PermissionIDs []string `json:"permission_ids"` // 必填 权限ids
}

// UserAssignRolesReq 分配角色
type UserAssignRolesReq struct {
	RoleIDs []string `json:"role_ids"` // 必填 角色roles
}
