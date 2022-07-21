package management

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/trustasia-com/go-sdk/wekey/management/types"
)

const (
	pathQueryRoles                 = "/ta-wekey-dash/roles"
	pathRoleInfo                   = "/ta-wekey-dash/roles"
	pathCreateRole                 = "/ta-wekey-dash/roles"
	pathDeleteRole                 = "/ta-wekey-dash/roles"
	pathUpdateRole                 = "/ta-wekey-dash/roles/:id"
	pathGetRolePermissions         = "/ta-wekey-dash/roles/:id/permissions"
	pathGetRoleUsers               = "/ta-wekey-dash/roles/:id/users"
	pathRoleAddPermissions         = "/ta-wekey-dash/roles/:id/permissions"
	pathRoleAddUsers               = "/ta-wekey-dash/roles/:id/users"
	pathRoleRemovePermissions      = "/ta-wekey-dash/roles/:id/permissions/:pid"
	pathQueryPermissionWithoutRole = "/ta-wekey-dash/roles/:id/query-permissions"
	pathQueryUserWithoutRole       = "/ta-wekey-dash/roles/:id/query-users"
)

// QueryRoles query
// 查询角色列表
func (c *Client) QueryRoles(req types.QueryRoleRequest) ([]types.QueryRoleResp, error) {
	query := make(map[string][]string)
	if req.Name != "" {
		query["name"] = []string{req.Name}
	}
	httpReq := &ClientRequest{
		Path:   pathQueryRoles,
		method: "GET",
		Query:  query,
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data []types.QueryRoleResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// RoleInfo info
// 查询角色详情
func (c *Client) RoleInfo(id string) (*types.RoleInfoResp, error) {
	httpReq := &ClientRequest{
		Path:   pathRoleInfo + "/" + id,
		method: "GET",
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data *types.RoleInfoResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// CreateRole create role
// 创建角色
func (c *Client) CreateRole(req types.CreateRoleRequest) (string, error) {

	if req.Name == "" {
		return "", errors.New("invalid app name")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   pathCreateRole,
		method: "POST",
		Body:   data,
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return "", res.Err
	}
	var resp struct {
		types.BaseResponse
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return "", err
	}
	if resp.Code != 0 {
		return "", fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data.ID, nil
}

// DelRole del role
// 删除角色
func (c *Client) DelRole(id string) error {
	if id == "" {
		return errors.New("invalid role id")
	}
	httpReq := &ClientRequest{
		Path:   pathDeleteRole + "/" + id,
		method: "DELETE",
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return res.Err
	}
	var resp struct {
		types.BaseResponse
		Data string `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return err
	}
	if resp.Code != 0 {
		return fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return nil
}

// UpdateRole update role
// 更新角色
func (c *Client) UpdateRole(id string, req types.UpdateRoleRequest) error {
	if id == "" {
		return errors.New("invalid role id")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathUpdateRole, ":id", id),
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

// GetRolePermissions get role permissions
// 获取角色的权限
func (c *Client) GetRolePermissions(id string) ([]types.RolePermissionResp, error) {

	if id == "" {
		return nil, errors.New("invalid params")
	}

	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathGetRolePermissions, ":id", id),
		method: "GET",
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data []types.RolePermissionResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// GetRoleUsers get role users
// 获取角色关联的用户
func (c *Client) GetRoleUsers(id string, req types.BaseQueryRequest) (*types.RoleUsersListResp, error) {

	if id == "" || req.Size < 1 || req.Page < 1 {
		return nil, errors.New("invalid params")
	}
	query := make(map[string][]string)
	query["page"] = []string{strconv.Itoa(req.Page)}
	query["size"] = []string{strconv.Itoa(req.Size)}
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathGetRoleUsers, ":id", id),
		method: "GET",
		Query:  query,
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data *types.RoleUsersListResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// RoleAddPermission role add permissions
// 为角色添加权限
func (c *Client) RoleAddPermission(id string, req types.AddRolePermissionRequest) error {
	if len(req.PermissionIDs) == 0 {
		return errors.New("invalid params")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathRoleAddPermissions, ":id", id),
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

// RoleAddUsers role add users
// 为角色添加用户
func (c *Client) RoleAddUsers(id string, req types.RoleAssignUsersReq) error {
	if len(req.UserIDs) == 0 {
		return errors.New("invalid params")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathRoleAddUsers, ":id", id),
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

// RoleRemovePermission role remove permission
// 移除角色的权限
func (c *Client) RoleRemovePermission(roleID, permissionID string) error {
	if roleID == "" || permissionID == "" {
		return errors.New("invalid params")
	}
	path := strings.ReplaceAll(pathRoleRemovePermissions, ":pid", permissionID)
	path = strings.ReplaceAll(path, ":id", roleID)
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

// GetPermissionsWithoutRole 查询角色可添加的权限
func (c *Client) GetPermissionsWithoutRole(roleID string,
	req types.QueryPermissionsWithoutRoleRequest) ([]types.QueryPermissionsWithoutRoleResp, error) {
	query := make(map[string][]string)

	if roleID == "" || req.ResourceID == "" {
		return nil, errors.New("invalid user id")
	}
	query["resource_id"] = []string{req.ResourceID}
	if req.Scope != "" {
		query["scope"] = []string{req.Scope}
	}

	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathQueryPermissionWithoutRole, ":id", roleID),
		method: "GET",
		Query:  query,
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data []types.QueryPermissionsWithoutRoleResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// GetUsersWithoutRole 查询角色可添加的权限
func (c *Client) GetUsersWithoutRole(roleID string,
	req types.QueryUserWithoutRoleRequest) ([]types.QueryUserWithoutRoleResp, error) {

	if roleID == "" {
		return nil, errors.New("invalid user id")
	}
	query := make(map[string][]string)
	if req.Search != "" {
		query["search"] = []string{req.Search}
	}

	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathQueryUserWithoutRole, ":id", roleID),
		method: "GET",
		Query:  query,
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data []types.QueryUserWithoutRoleResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}
