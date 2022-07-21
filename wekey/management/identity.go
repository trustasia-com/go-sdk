package management

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/trustasia-com/go-sdk/wekey/management/types"
)

const (
	pathGetIdpConfigs    = "/ta-wekey-dash/conns"
	pathGetIdpConfigInfo = "/ta-wekey-dash/conns/:provider/:identifier/info"
	pathGetIdpApps       = "/ta-wekey-dash/conns/:provider/:identifier/apps"
	pathCreateIdp        = "/ta-wekey-dash/conns/:provider"
	pathUpdateIdp        = "/ta-wekey-dash/conns/:provider/:identifier"
	pathDelIdp           = "/ta-wekey-dash/conns/:provider/:identifier"
)

// GetIdpConfigs 获取已配置的连接源
func (c *Client) GetIdpConfigs(req types.IdpConfigsReq) ([]types.IdpConfigsResponse, error) {
	query := make(map[string][]string)
	if req.Type == "" {
		return nil, errors.New("invalid type")
	}
	query["type"] = []string{req.Type}
	httpReq := &ClientRequest{
		Path:   pathGetIdpConfigs,
		method: "GET",
		Query:  query,
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data []types.IdpConfigsResponse `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// GetIdpApps get idp apps
// 获取身份源关联应用
func (c *Client) GetIdpApps(req types.IdpAppsReq) ([]types.IdpAppsResponse, error) {
	if req.Provider == "" || req.Identifier == "" {
		return nil, errors.New("invalid provider or identifier")
	}
	path := strings.ReplaceAll(pathGetIdpApps, ":provider", req.Provider)
	path = strings.ReplaceAll(path, ":identifier", req.Identifier)
	httpReq := &ClientRequest{
		Path:   path,
		method: "GET",
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data []types.IdpAppsResponse `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}
