package types

// AddIdpGithubRequest 添加github身份源
type AddIdpGithubRequest struct {
	Name         string   `json:"name"`          // 必填，显示名称，最长64个字符
	ClientID     string   `json:"client_id"`     // 必填，你应用的app id
	ClientSecret string   `json:"client_secret"` // 必填，你应用的app secret
	Scopes       []string `json:"scopes"`        // 选填，scopes
	SyncProfile  bool     `json:"sync_profile"`  // 必填，是否同步
}

// UpdateIdpGithub 编辑Github信息
type UpdateIdpGithub struct {
	Name         string   `json:"name"`          // 必填，显示名称，最长64个字符
	ClientID     string   `json:"client_id"`     // 必填，你应用的app id
	ClientSecret string   `json:"client_secret"` // 必填，你应用的app secret
	Scopes       []string `json:"scopes"`        // 选填，scopes
	SyncProfile  bool     `json:"sync_profile"`  // 必填，是否同步
}

// IdpGithubInfoResp 详情
type IdpGithubInfoResp struct {
	Name       string `json:"name"`       // 显示名称
	Identifier string `json:"identifier"` // 唯一标识符
	Icon       string `json:"icon"`       // icon

	ClientID     string   `json:"client_id"`     // app id
	ClientSecret string   `json:"client_secret"` // app secret
	Scopes       []string `json:"scopes"`        // scopes
	CallbackURL  string   `json:"callback_url"`  // 回调地址
	SyncProfile  bool     `json:"sync_profile"`  // 是否同步
}

// AddIdpGitlabRequest 添加gitlab身份源
type AddIdpGitlabRequest struct {
	Name         string `json:"name"`          // 必填，显示名称，最长64个字符
	ClientID     string `json:"client_id"`     // 必填，你应用的app id
	ClientSecret string `json:"client_secret"` // 必填，你应用的app secret
	SyncProfile  bool   `json:"sync_profile"`  // 必填，是否同步
}

// UpdateIdpGitlab  编辑gitlab
type UpdateIdpGitlab struct {
	Name         string `json:"name"`          // 必填，显示名称，最长64个字符
	ClientID     string `json:"client_id"`     // 必填，你应用的app id
	ClientSecret string `json:"client_secret"` // 必填，你应用的app secret
	SyncProfile  bool   `json:"sync_profile"`  // 必填，是否同步
}

// IdpGitlabInfoResp 详情
type IdpGitlabInfoResp struct {
	Name       string `json:"name"`       // 显示名称
	Identifier string `json:"identifier"` // 唯一标识符
	Icon       string `json:"icon"`       // icon

	ClientID     string `json:"client_id"`     // app id
	ClientSecret string `json:"client_secret"` // app secret
	CallbackURL  string `json:"callback_url"`  // 回调地址
	SyncProfile  bool   `json:"sync_profile"`  // 是否同步
}

// AddIdpWeiboRequest 添加微博身份源
type AddIdpWeiboRequest struct {
	Name         string   `json:"name"`          // 必填，显示名称，最长64个字符
	ClientID     string   `json:"client_id"`     // 必填，你应用的app id
	ClientSecret string   `json:"client_secret"` // 必填，你应用的app secret
	Scopes       []string `json:"scopes"`        // 选填，scopes
	SyncProfile  bool     `json:"sync_profile"`  // 必填，是否同步
}

// UpdateIdpWeibo 编辑Weibo
type UpdateIdpWeibo struct {
	Name         string   `json:"name"`          // 必填，显示名称，最长64个字符
	ClientID     string   `json:"client_id"`     // 必填，你应用的app id
	ClientSecret string   `json:"client_secret"` // 必填，你应用的app secret
	Scopes       []string `json:"scopes"`        // 选填，scopes
	SyncProfile  bool     `json:"sync_profile"`  // 必填，是否同步
}

// IdpWeiboInfoResp 详情
type IdpWeiboInfoResp struct {
	Name       string `json:"name"`       // 显示名称
	Identifier string `json:"identifier"` // 唯一标识符
	Icon       string `json:"icon"`       // icon

	ClientID     string   `json:"client_id"`     // app id
	ClientSecret string   `json:"client_secret"` // app secret
	Scopes       []string `json:"scopes"`        // scopes
	CallbackURL  string   `json:"callback_url"`  // 回调地址
	SyncProfile  bool     `json:"sync_profile"`  // 是否同步
}

// AddIdpGitEERequest 添加GitEE身份源
type AddIdpGitEERequest struct {
	Name         string   `json:"name"`          // 必填，显示名称，最长64个字符
	ClientID     string   `json:"client_id"`     // 必填，你应用的app id
	ClientSecret string   `json:"client_secret"` // 必填，你应用的app secret
	Scopes       []string `json:"scopes"`        // 选填，scopes
	SyncProfile  bool     `json:"sync_profile"`  // 必填，是否同步
}

// UpdateIdpGitEE 编辑GitEE
type UpdateIdpGitEE struct {
	Name         string   `json:"name"`          // 必填，显示名称，最长64个字符
	ClientID     string   `json:"client_id"`     // 必填，你应用的app id
	ClientSecret string   `json:"client_secret"` // 必填，你应用的app secret
	Scopes       []string `json:"scopes"`        // 选填，scopes
	SyncProfile  bool     `json:"sync_profile"`  // 必填，是否同步
}

// IdpGitEEInfoResp 详情
type IdpGitEEInfoResp struct {
	Name       string `json:"name"`       // 显示名称
	Identifier string `json:"identifier"` // 唯一标识符
	Icon       string `json:"icon"`       // icon

	ClientID     string   `json:"client_id"`     // app id
	ClientSecret string   `json:"client_secret"` // app secret
	CallbackURL  string   `json:"callback_url"`  // 回调地址
	Scopes       []string `json:"scopes"`        // scopes
	SyncProfile  bool     `json:"sync_profile"`  // 是否同步
}

// AddIdpBaiduRequest 添加百度身份源
type AddIdpBaiduRequest struct {
	Name         string `json:"name"`          // 必填，显示名称，最长64个字符
	ClientID     string `json:"client_id"`     // 必填，你应用的app id
	ClientSecret string `json:"client_secret"` // 必填，你应用的app secret
	SyncProfile  bool   `json:"sync_profile"`  // 必填，是否同步
}

// UpdateIdpBaidu 编辑百度身份源
type UpdateIdpBaidu struct {
	Name         string `json:"name"`          // 必填，显示名称，最长64个字符
	ClientID     string `json:"client_id"`     // 必填，你应用的app id
	ClientSecret string `json:"client_secret"` // 必填，你应用的app secret
	SyncProfile  bool   `json:"sync_profile"`  // 必填，是否同步
}

// IdpBaiduInfoResp 详情
type IdpBaiduInfoResp struct {
	Name       string `json:"name"`       // 显示名称
	Identifier string `json:"identifier"` // 唯一标识符
	Icon       string `json:"icon"`       // icon

	ClientID     string `json:"client_id"`     // app id
	ClientSecret string `json:"client_secret"` // app secret
	CallbackURL  string `json:"callback_url"`  // 回调地址
	SyncProfile  bool   `json:"sync_profile"`  // 是否同步
}

// AddIdpQQRequest 添加qq身份源
type AddIdpQQRequest struct {
	Name         string   `json:"name"`          // 必填，显示名称，最长64个字符
	ClientID     string   `json:"client_id"`     // 必填，你应用的app id
	ClientSecret string   `json:"client_secret"` // 必填，你应用的app secret
	Scopes       []string `json:"scopes"`        // 选填，scopes
	SyncProfile  bool     `json:"sync_profile"`  // 必填，是否同步
}

// UpdateIdpQQ 编辑qq身份源
type UpdateIdpQQ struct {
	Name         string   `json:"name"`          // 必填，显示名称，最长64个字符
	ClientID     string   `json:"client_id"`     // 必填，你应用的app id
	ClientSecret string   `json:"client_secret"` // 必填，你应用的app secret
	Scopes       []string `json:"scopes"`        // 选填，scopes
	SyncProfile  bool     `json:"sync_profile"`  // 必填，是否同步
}

// IdpQQInfoResp 详情
type IdpQQInfoResp struct {
	Name       string `json:"name"`       // 显示名称
	Identifier string `json:"identifier"` // 唯一标识符
	Icon       string `json:"icon"`       // icon

	ClientID     string   `json:"client_id"`     // app id
	ClientSecret string   `json:"client_secret"` // app secret
	Scopes       []string `json:"scopes"`        // scopes
	CallbackURL  string   `json:"callback_url"`  // 回调地址
	SyncProfile  bool     `json:"sync_profile"`  // 是否同步
}
