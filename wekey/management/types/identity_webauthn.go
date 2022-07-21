package types

// IdpWebauthnInfoResp 详情
type IdpWebauthnInfoResp struct {
	Name       string `json:"name"`       // 显示名称
	Identifier string `json:"identifier"` // 唯一标识符
	Icon       string `json:"icon"`       // icon
}
