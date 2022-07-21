package management

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/trustasia-com/go-sdk/wekey/management/types"
)

// GetIdpWebauthnConfig webauthn
// 获取webauthn身份源（fido）配置
func (c *Client) GetIdpWebauthnConfig() (*types.IdpWebauthnInfoResp, error) {
	path := strings.ReplaceAll(pathGetIdpConfigInfo, ":provider", "webauthn")
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
		Data *types.IdpWebauthnInfoResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// AddIdpWebauthn add webauthn
// 添加webauthn身份源（fido）配置
func (c *Client) AddIdpWebauthn() error {

	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathCreateIdp, ":provider", "webauthn"),
		method: "POST",
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

// DelIdpWebauthn del webauthn
// 删除webauthn身份源（fido）配置
func (c *Client) DelIdpWebauthn() error {
	path := strings.ReplaceAll(pathDelIdp, ":provider", "webauthn")
	path = strings.ReplaceAll(path, ":identifier", "webauthn")
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
