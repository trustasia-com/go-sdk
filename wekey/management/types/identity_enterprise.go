package types

// AddIdpLDAPRequest 添加LDAP身份源
type AddIdpLDAPRequest struct {
	Name        string `json:"name"`         // 必填，显示名称，最长64个字符
	Identifier  string `json:"identifier"`   // 必填，唯一标识符，最长35个字符
	SyncProfile bool   `json:"sync_profile"` // 必填，是否同步用户信息
}

// UpdateIdpLDAP  编辑LDAP
type UpdateIdpLDAP struct {
	Name        string `json:"name"`         // 必填，显示名称，最长64个字符
	SyncProfile bool   `json:"sync_profile"` // 必填，是否同步用户信息
}

// IdpLDAPInfoResp 详情
type IdpLDAPInfoResp struct {
	Name       string `json:"name"`       // 显示名称
	Identifier string `json:"identifier"` // 唯一标识符
	Icon       string `json:"icon"`       // icon

	SyncProfile     bool   `json:"sync_profile"`     // 是否同步用户信息
	TicketURL       string `json:"ticket_url"`       // ticket url
	ConnectorStatus int    `json:"connector_status"` // LDAP连接器状态，0为未连接，1为已连接
}
