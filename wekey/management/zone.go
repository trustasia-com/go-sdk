package management

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/trustasia-com/go-sdk/pkg/valid"
	"github.com/trustasia-com/go-sdk/wekey/management/types"
)

const (
	pathGetZoneBasicInfo     = "/ta-wekey-dash/zones/basic"
	pathGetZoneMemberInfo    = "/ta-wekey-dash/zones/members"
	pathGetZoneSecretInfo    = "/ta-wekey-dash/zones/secrets"
	pathGetZoneAdvanceInfo   = "/ta-wekey-dash/zones/advance"
	pathUpdateZoneBasicInfo  = "/ta-wekey-dash/zones/basic"
	pathUpdateZoneMemberInfo = "/ta-wekey-dash/zones/members/:id"
	pathDeleteZoneMemberInfo = "/ta-wekey-dash/zones/members/:id"
	pathDeleteZoneSecrets    = "/ta-wekey-dash/zones/secrets"
	pathDeleteZoneAdvance    = "/ta-wekey-dash/zones/advance"
)

// GetZoneBasicInfo get zone basic info
// 获取空间基础信息
func (c *Client) GetZoneBasicInfo() (*types.ZoneInfoBasicResp, error) {
	httpReq := &ClientRequest{
		Path:   pathGetZoneBasicInfo,
		method: "GET",
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data *types.ZoneInfoBasicResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// GetZoneMembersInfo get zone members info
// 获取空间成员信息
func (c *Client) GetZoneMembersInfo() ([]types.ZoneInfoMemberResp, error) {
	httpReq := &ClientRequest{
		Path:   pathGetZoneMemberInfo,
		method: "GET",
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data []types.ZoneInfoMemberResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// GetZoneSecretsInfo get zone secrets info
// 获取空间密钥信息
func (c *Client) GetZoneSecretsInfo() (*types.ZoneInfoSecrets, error) {
	httpReq := &ClientRequest{
		Path:   pathGetZoneSecretInfo,
		method: "GET",
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data *types.ZoneInfoSecrets `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// GetZoneAdvanceInfo get zone advance info 高级设置信息
// 获取空间高级信息
func (c *Client) GetZoneAdvanceInfo() (*types.ZoneInfoAdvanceResp, error) {
	httpReq := &ClientRequest{
		Path:   pathGetZoneAdvanceInfo,
		method: "GET",
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data *types.ZoneInfoAdvanceResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// UpdateZoneBasicInfo update zone basic info 编辑空间基础信息
func (c *Client) UpdateZoneBasicInfo(req types.UpdateZoneBasicInfoReq) error {

	if !valid.IsURL(req.Logo) || req.Name == "" {
		return errors.New("invalid params")
	}

	if req.Logo != "" && !valid.IsURL(req.Logo) {
		return errors.New("invalid params")
	}

	if req.EmailSupport != "" && !valid.IsASCIIEmail(req.EmailSupport) {
		return errors.New("invalid params")
	}

	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   pathUpdateZoneBasicInfo,
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

// UpdateZoneMembersInfo update zone members info 编辑空间成员信息
func (c *Client) UpdateZoneMembersInfo(id string, req types.UpdateZoneMembersReq) error {

	if len(req.Roles) == 0 || id == "" {
		return errors.New("invalid params")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathUpdateZoneMemberInfo, ":id", id),
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

// DeleteZoneMembers del zone members 删除空间成员
func (c *Client) DeleteZoneMembers(id string) error {
	if id == "" {
		return errors.New("invalid id")
	}

	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathDeleteZoneMemberInfo, ":id", id),
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

// DeleteZoneSecrets 轮换空间密钥
func (c *Client) DeleteZoneSecrets() error {
	httpReq := &ClientRequest{
		Path:   pathDeleteZoneSecrets,
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

// UpdateZoneAdvanceInfo 更新空间高级信息
func (c *Client) UpdateZoneAdvanceInfo(req types.UpdateZoneAdvanceReq) error {

	if req.CookieMode != "persistent" && req.CookieMode != "non-persistent" {
		return errors.New("invalid params")
	}

	if req.FinalTimeout < 60 || req.FinalTimeout > 2592000 {
		return errors.New("invalid params")
	}

	if req.InactivityTimeout < 60 || req.InactivityTimeout > 2592000 {
		return errors.New("invalid params")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   pathDeleteZoneAdvance,
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
