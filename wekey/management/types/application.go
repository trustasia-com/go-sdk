package types

// QueryAppRequest   query app
type QueryAppRequest struct {
	Name string `form:"name"` // 名字 最长64个字符
}

// QueryAppResp resp
type QueryAppResp struct {
	ID       string `json:"id"`        // id
	Name     string `json:"name"`      // 名字
	Logo     string `json:"logo"`      // logo
	ClientID string `json:"client_id"` // client id
	LoginURL string `json:"login_url"` // 发起登录uri
}

// AppInfoResp app info resp
type AppInfoResp struct {
	// basic
	ClientID     string `json:"client_id"`     // client id
	ClientSecret string `json:"client_secret"` // client secret
	Name         string `json:"name"`          // app 名字
	Logo         string `json:"logo"`          // app logo
	Description  string `json:"description"`   // app 描述

	// 端点信息
	OauthAuthorize string `json:"oauth_authorize"` // OAuth Authorize
	OauthToken     string `json:"oauth_token"`     // OAuth Token
	OauthUserInfo  string `json:"oauth_user_info"` // OAuth User Info
	OpenIDConfig   string `json:"openid_config"`   // OpenID Config
	JSONWebKey     string `json:"json_web_key"`    // JSON Web Key
	SAMLProtocol   string `json:"saml_protocol"`   // SAML Protocol
	SAMLMetadata   string `json:"smal_metadata"`   // SAML Metadata

	// 授权配置
	TokenEndpointAuthMethod string   `json:"token_endpoint_auth_method"` // Token端点认证
	LoginURL                string   `json:"login_url"`                  // 发起登录uri
	LoginCallbackURLs       []string `json:"login_callback_urls"`        // 登录回调urls
	LogoutCallbackURLs      []string `json:"logout_callback_urls"`       // 登出回调urls
	IDTokenExpire           int      `json:"id_token_expire"`            // ID Token 过期时间
	AccessTokenExpire       int      `json:"access_token_expire"`        // access_token 过期时间
	RefreshTokenExpire      int      `json:"refresh_token_expire"`       // refresh token 过期时间

	// 高级设置
	GrantTypes     []string `json:"grant_types"`     // 授权类型
	SignAlgo       string   `json:"sign_algo"`       // 签名算法
	OIDCConformant bool     `json:"oidc_conformant"` // 是否遵守OIDC
	CertPem        string   `json:"cert_pem"`        // 证书
	CertPemSHA1    string   `json:"cert_pem_sha1"`   // 证书id
}

// CreateAppRequest add app
type CreateAppRequest struct {
	Name string `json:"name"` // app名字，最大64个字符
}

// AppIDPResp idp resp
type AppIDPResp struct {
	Type       string `json:"type"`       // 类型database,social,enterprise,passwordless,webauthn
	Provider   string `json:"provider"`   // 身份源标识
	Identifier string `json:"identifier"` // 唯一标识
	Protocol   string `json:"protocol"`   // 协议类型
	Name       string `json:"name"`       // 身份源名字
	Icon       string `json:"icon"`       // logo
	Status     int    `json:"status"`     // 状态,0关闭，1开启
}

// UpdateAppBasicRequest edit basic
type UpdateAppBasicRequest struct {
	Name        string `json:"name"`        // 必填 app名字，最长64个字符
	Logo        string `json:"logo"`        // 必填，app的logo，一个https图片链接
	Description string `json:"description"` // 选填，app描述，最长140个字符
}

// UpdateAppGrantRequest edit app grant info
type UpdateAppGrantRequest struct {
	TokenEndpointAuthMethod string   `json:"token_endpoint_auth_method"` // 必填，值为Post，Basic，None中的一个
	LoginURL                string   `json:"login_url"`                  // 必填，发起登录uri
	LoginCallbackURLs       []string `json:"login_callback_urls"`        // 选填，登录回调urls
	LogoutCallbackURLs      []string `json:"logout_callback_urls"`       // 选填，登出回调urls
	IDTokenExpire           int      `json:"id_token_expire"`            // 必填，ID Token 过期时间，单位秒
	AccessTokenExpire       int      `json:"access_token_expire"`        // 必填，access_token 过期时间，单位秒
	RefreshTokenExpire      int      `json:"refresh_token_expire"`       // 必填，refresh token 过期时间，单位秒
}

// UpdateAppAdvanceRequest edit app advance info
type UpdateAppAdvanceRequest struct {
	GrantTypes     []string `json:"grant_types"`     // 必填，授权类型
	SignAlgo       string   `json:"sign_algo"`       // 必填，签名算法HS256或RS256
	OIDCConformant bool     `json:"oidc_conformant"` // 选填，是否遵守OIDC
}

// QueryResourceRequest   query resource
type QueryResourceRequest struct {
	Name string `form:"name"` // 必填，名字，最长64字符
}

// QueryResourceResp resp
type QueryResourceResp struct {
	ID       string `json:"id"`       // id
	Type     int    `json:"type"`     // 类型0用户资源组，1为系统资源组，系统资源组不可做任何写操作
	Name     string `json:"name"`     // 名字
	Audience string `json:"audience"` // API Audience
	Editable bool   `json:"editable"` // 是否可编辑
}

// ResourceInfoResp  resp
type ResourceInfoResp struct {
	ID                  string `json:"id"`                    // id
	Type                int    `json:"type"`                  // 类型0用户资源组，1为系统资源组，系统资源组不可做任何写操作
	Name                string `json:"name"`                  // 名字
	Editable            bool   `json:"editable"`              // 是否可编辑
	Audience            string `json:"audience"`              // 唯一标识
	RBACEnable          bool   `json:"rabc_enable"`           // RBAC开关
	TokenWithPermission bool   `json:"token_with_permission"` // 是否添加权限到access token
	AllowSkipConsent    bool   `json:"allow_skip_consent"`    // 是否跳过用户consent
	AllowOfflineAccess  bool   `json:"allow_offline_access"`  // 是否允许离线访问
}

// CreateResourceRequest add resource
type CreateResourceRequest struct {
	Name     string `json:"name"`     // 必填，名字，最长64个字符
	Audience string `json:"audience"` // 必填，资源组地址,当前空间唯一，一个url地址
}

// UpdateResourceRequest  edit
type UpdateResourceRequest struct {
	Name                string `json:"name"`                  // 必填，名字，最长64个字符
	RBACEnable          bool   `json:"rabc_enable"`           // 必填，RBAC开关
	TokenWithPermission bool   `json:"token_with_permission"` // 必填，是否添加权限到access token
	AllowSkipConsent    bool   `json:"allow_skip_consent"`    // 必填，是否跳过用户consent
	AllowOfflineAccess  bool   `json:"allow_offline_access"`  // 必填，是否允许离线访问
}

// QueryPermissionRequest query
type QueryPermissionRequest struct {
	ResourceID string `form:"resource_id"` // 必填，资源组id
	Scope      string `form:"scope"`       // 选填，权限，最长64个字符
}

// QueryPermissionResp query permission
type QueryPermissionResp struct {
	ID          string `json:"id"`          // 权限id
	Scope       string `json:"scope"`       // 权限
	Description string `json:"description"` // 描述
}

// CreatePermissionRequest add
type CreatePermissionRequest struct {
	ResourceID  string `json:"resource_id"` // 必填，资源组id
	Scope       string `json:"scope"`       // 必填，权限scope，最长64个字符
	Description string `json:"description"` // 选填，描述，最长512字符
}

// ResourceAppsResp resource apps
type ResourceAppsResp struct {
	ID       string `json:"id"`        // app id
	ClientID string `json:"client_id"` // app client id
	Name     string `json:"name"`      // app名字
	Status   int    `json:"status"`    // 与资源组关系，0为未关联，1为已关联
}

// AppResourcesResp app resources
type AppResourcesResp struct {
	ID       string `json:"id"`       // 资源组id
	Name     string `json:"name"`     // resource name
	Audience string `json:"audience"` // audience
	Status   int    `json:"status"`   // 与应用关系，0为未关联，1为已关联
}
