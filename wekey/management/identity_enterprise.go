package management

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/trustasia-com/go-sdk/wekey/management/types"
)

// GetIdpLDAPConfig ldap
// 获取ldap的配置信息
func (c *Client) GetIdpLDAPConfig(identifier string) (*types.IdpLDAPInfoResp, error) {
	path := strings.ReplaceAll(pathGetIdpConfigInfo, ":provider", "ldap")
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
		Data *types.IdpLDAPInfoResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// AddIdpLDAP add ldap
// 添加ldap身份源
func (c *Client) AddIdpLDAP(req types.AddIdpLDAPRequest) error {
	if req.Name == "" || req.Identifier == "" {
		return errors.New("invalid params")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathCreateIdp, ":provider", "ldap"),
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

// UpdateIdpLDAP update ldap conf
// 更新ldap配置
func (c *Client) UpdateIdpLDAP(identifier string, req types.UpdateIdpLDAP) error {
	if req.Name == "" {
		return errors.New("invalid params")
	}
	data, _ := json.Marshal(req)
	path := strings.ReplaceAll(pathUpdateIdp, ":provider", "ldap")
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

// DelIdpLDAP del ldap
// 删除ldap身份源
func (c *Client) DelIdpLDAP(identifier string) error {
	path := strings.ReplaceAll(pathDelIdp, ":provider", "ldap")
	path = strings.ReplaceAll(path, ":identifier", identifier)
	httpReq := &ClientRequest{
		Path:   path,
		method: "DELETE",
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
