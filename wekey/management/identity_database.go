package management

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/trustasia-com/go-sdk/wekey/management/types"
)

const (
	pathUpdateIdpDatabasePasswordStrength = "/ta-wekey-dash/conns/db/:identifier/password-strength"
	pathUpdateIdpDatabasePasswordRotation = "/ta-wekey-dash/conns/db/:identifier/password-rotation"
)

// GetIdpDatabaseConfig database
// 获取数据库身份源配置
func (c *Client) GetIdpDatabaseConfig(identifier string) (*types.IdpDatabaseInfoResp, error) {
	path := strings.ReplaceAll(pathGetIdpConfigInfo, ":provider", "database")
	path = strings.ReplaceAll(path, ":identifier", identifier)
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
		Data *types.IdpDatabaseInfoResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// AddIdpDatabase add database
// 添加数据库身份源
func (c *Client) AddIdpDatabase(req types.AddIdpDatabaseRequest) error {
	if req.Name == "" || req.Identifier == "" {
		return errors.New("invalid params")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathCreateIdp, ":provider", "database"),
		method: "POST",
		Body:   data,
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return res.Err
	}
	var resp struct {
		types.BaseResponse
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return err
	}
	if resp.Code != 0 {
		return fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return nil
}

// UpdateIdpDatabase update database conf
// 更新数据库身份源
func (c *Client) UpdateIdpDatabase(identifier string, req types.UpdateIdpDatabase) error {
	data, _ := json.Marshal(req)
	path := strings.ReplaceAll(pathUpdateIdp, ":provider", "database")
	path = strings.ReplaceAll(path, ":identifier", identifier)
	httpReq := &ClientRequest{
		Path:   path,
		method: "PUT",
		Body:   data,
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return res.Err
	}
	var resp struct {
		types.BaseResponse
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return err
	}
	if resp.Code != 0 {
		return fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return nil
}

// UpdateIdpDatabasePasswordStrength update database conf
// 更新数据库密码强度
func (c *Client) UpdateIdpDatabasePasswordStrength(identifier string, req types.UpdateIdpDatabasePasswordStrength) error {
	data, _ := json.Marshal(req)
	path := strings.ReplaceAll(pathUpdateIdpDatabasePasswordStrength, ":identifier", identifier)
	httpReq := &ClientRequest{
		Path:   path,
		method: "PUT",
		Body:   data,
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return res.Err
	}
	var resp struct {
		types.BaseResponse
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return err
	}
	if resp.Code != 0 {
		return fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return nil
}

// UpdateIdpDatabasePasswordRotation update database conf
// 更新数据库轮换密钥
func (c *Client) UpdateIdpDatabasePasswordRotation(identifier string, req types.UpdateIdpDatabasePasswordRotation) error {
	data, _ := json.Marshal(req)
	path := strings.ReplaceAll(pathUpdateIdpDatabasePasswordRotation, ":identifier", identifier)
	httpReq := &ClientRequest{
		Path:   path,
		method: "PUT",
		Body:   data,
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return res.Err
	}
	var resp struct {
		types.BaseResponse
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return err
	}
	if resp.Code != 0 {
		return fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return nil
}
