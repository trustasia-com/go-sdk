package types

// IdpConfigsReq  configs req
type IdpConfigsReq struct {
	Type string `form:"type"` //  类型database social enterprise passwordless webauthn
}

// IdpConfigsResponse configs
type IdpConfigsResponse struct {
	ID              string `json:"id"`                // id
	Identifier      string `json:"identifier"`        // identifier
	Provider        string `json:"provider"`          // provider
	Name            string `json:"name"`              // 身份源名字
	Icon            string `json:"icon"`              // icon
	Description     string `json:"description"`       // idp 描述
	Protocol        string `json:"protocol"`          // 协议类型
	EnabledAppCount int    `json:"enabled_app_count"` // 启用数量
}

// IdpAppsReq  get apps
type IdpAppsReq struct {
	Provider   string `uri:"provider"`   // 必填 身份源/github/ldap/saml/...
	Identifier string `uri:"identifier"` // 必填 配置的唯一identifier
}

// IdpAppsResponse idp apps
type IdpAppsResponse struct {
	ID          string `json:"id"`          // 应用id
	Name        string `json:"name"`        // 应用名
	ClientID    string `json:"client_id"`   // Client ID
	Logo        string `json:"logo"`        // 应用logo
	Description string `json:"description"` // 应用描述
	IdpStatus   int    `json:"idp_status"`  // 状态，0代表关，1代表开
}

// IdpConfigAppStateReq  config
type IdpConfigAppStateReq struct {
	Provider   string `uri:"provider"`   // 必传 身份源/github/ldap/saml/...
	Identifier string `uri:"identifier"` // 必传 唯一标识符配置id
	AppID      string `uri:"appId"`      // 必传 应用id
}
