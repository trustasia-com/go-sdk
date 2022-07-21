package types

// AddIdpDatabaseRequest 添加数据库身份源
type AddIdpDatabaseRequest struct {
	Name       string `json:"name"`        // 必填，显示名称，最长64个字符
	Identifier string `json:"identifier"`  // 必填， 唯一标识符，最长35个字符
	ImportUser bool   `json:"import_user"` // 必填，是否导入用户
}

// UpdateIdpDatabase 编辑信息
type UpdateIdpDatabase struct {
	ImportUser      bool     `json:"import_user"`      // 必填,是否导入用户
	LoginMethods    []string `json:"login_methods" `   // 选填，登录方式
	RegisterMethods []string `json:"register_methods"` // 选填，注册方式
}

// UpdateIdpDatabasePasswordStrength 编辑密码强度
type UpdateIdpDatabasePasswordStrength struct {
	PasswordPolicy string `json:"password_policy"` // 必填，密码策略，从高到低为：excellent good fair low none
	PasswordLength int    `json:"password_length"` // 必填，密码长度，最长64个字符
}

// UpdateIdpDatabasePasswordRotation  编辑密码轮换
type UpdateIdpDatabasePasswordRotation struct {
	PasswordDicEnable bool     `json:"password_dic_enable"` // 必填 密码字典开关
	PasswordDic       []string `json:"password_dic"`        // 必填 密码字典
}

// IdpDatabaseInfoResp 详情
type IdpDatabaseInfoResp struct {
	Name       string `json:"name"`       // 显示名称
	Identifier string `json:"identifier"` // 唯一标识符
	Icon       string `json:"icon"`       // icon

	PasswordPolicy    string   `json:"password_policy"`     // 密码策略
	PasswordLength    int      `json:"password_length"`     // 密码长度
	PasswordDicEnable bool     `json:"password_dic_enable"` // 密码字典开关
	PasswordDic       []string `json:"password_dic"`        // 密码字典
	CustomDb          bool     `json:"custom_db"`           // 自定义数据库
	CustomScripts     string   `json:"custom_scripts"`      // 自定义脚本
	ImportUser        bool     `json:"import_user"`         // 是否导入用户
	AllowSignup       bool     `json:"allow_signup"`        // 是否允许注册
	LoginMethods      []string `json:"login_methods" `      // 登录方式
	RegisterMethods   []string `json:"register_methods"`    // 注册方式
}
