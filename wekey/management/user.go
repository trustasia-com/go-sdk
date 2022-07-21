package management

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/trustasia-com/go-sdk/pkg/valid"
	"github.com/trustasia-com/go-sdk/wekey/management/types"
)

const (
	pathQueryUsers                = "/ta-wekey-dash/users"
	pathUserInfo                  = "/ta-wekey-dash/users"
	pathCreateUsers               = "/ta-wekey-dash/users"
	pathDelUsers                  = "/ta-wekey-dash/users"
	pathUpdateUsersBasic          = "/ta-wekey-dash/users/:id/basic"
	pathUpdateUsersPicture        = "/ta-wekey-dash/users/:id/picture"
	pathUpdateUsersPassword       = "/ta-wekey-dash/users/:id/password"
	pathUpdateUsersMetadata       = "/ta-wekey-dash/users/:id/metadata"
	pathBlockUser                 = "/ta-wekey-dash/users/:id/block"
	pathUnBlockUser               = "/ta-wekey-dash/users/:id/unblock"
	pathUserPermissionsInfo       = "/ta-wekey-dash/users/:id/permissions"
	pathUserWithoutPermissionInfo = "/ta-wekey-dash/users/:id/query-permissions"
	pathUserRolesInfo             = "/ta-wekey-dash/users/:id/roles"
	pathUserWithoutRolesInfo      = "/ta-wekey-dash/users/:id/query-roles"
	pathUserApps                  = "/ta-wekey-dash/users/:id/apps"
	pathUserAssignPermissions     = "/ta-wekey-dash/users/:id/permissions"
	pathUserRemovePermissions     = "/ta-wekey-dash/users/:id/permissions/:pid"
	pathUserAssignRoles           = "/ta-wekey-dash/users/:id/roles"
	pathUserRemoveRoles           = "/ta-wekey-dash/users/:id/roles/:rid"
	pathUserRemoveApp             = "/ta-wekey-dash/users/:id/apps/:aid"
)

// QueryUsers query user
// 查询用户列表
func (c *Client) QueryUsers(req types.QueryUserReq) (*types.UserListResp, error) {
	query := make(map[string][]string)
	if req.Search != "" {
		query["search"] = []string{req.Search}
	}
	query["page"] = []string{strconv.Itoa(req.Page)}
	query["size"] = []string{strconv.Itoa(req.Size)}
	httpReq := &ClientRequest{
		Path:   pathQueryUsers,
		method: "GET",
		Query:  query,
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data *types.UserListResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// UserInfo user info
// 查询用户详情
func (c *Client) UserInfo(id string) (*types.UserDetailResp, error) {
	if id == "" {
		return nil, errors.New("invalid user id")
	}

	httpReq := &ClientRequest{
		Path:   pathUserInfo + "/" + id,
		method: "GET",
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data *types.UserDetailResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// CreateUser create user
// 创建用户
func (c *Client) CreateUser(req types.CreateUserReq) (string, error) {

	if req.Username == "" && req.PhoneNumber == "" && req.Email == "" {
		return "", errors.New("invalid params")
	}

	if req.Password == "" {
		return "", errors.New("invalid password")
	}

	req.Password = c.encrypt(req.Password)

	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   pathCreateUsers,
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

// DelUser del user
// 删除用户
func (c *Client) DelUser(id string) error {
	if id == "" {
		return errors.New("invalid user id")
	}
	httpReq := &ClientRequest{
		Path:   pathDelUsers + "/" + id,
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

// UpdateUserBasic update user basic info
// 更新用户基本信息
func (c *Client) UpdateUserBasic(id string, req types.UpdateUserBasicRequest) error {
	if req.Nickname == "" && req.PhoneNumber == "" && req.Email == "" {
		return errors.New("invalid params")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathUpdateUsersBasic, ":id", id),
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

// UpdateUserPicture update user picture info
// 更新用户头像
func (c *Client) UpdateUserPicture(id string, req types.UpdateUserPictureRequest) error {
	if !valid.IsURL(req.Picture) {
		return errors.New("invalid params")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathUpdateUsersPicture, ":id", id),
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

// UpdateUserPassword update user password info
// 修改用户密码
func (c *Client) UpdateUserPassword(id string, req types.UpdateUserPwdRequest) error {
	if req.Password == "" {
		return errors.New("invalid params")
	}

	req.Password = c.encrypt(req.Password)
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathUpdateUsersPassword, ":id", id),
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

// UpdateUserMetadata update user metadata
// 更新用户metadata信息
func (c *Client) UpdateUserMetadata(id string, req types.UpdateUserMetadataRequest) error {

	data, err := json.Marshal(req)
	if err != nil {
		return errors.New("invalid params")
	}
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathUpdateUsersMetadata, ":id", id),
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
	if err = json.Unmarshal(res.Body, &resp); err != nil {
		return err
	}
	if resp.Code != 0 {
		return fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return nil
}

// BlockUser block user 锁定用户，用户将无法再登录
// 锁定用户
func (c *Client) BlockUser(id string) error {

	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathBlockUser, ":id", id),
		method: "PUT",
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

// UnBlockUser unblock user
// 解锁用户
func (c *Client) UnBlockUser(id string) error {

	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathUnBlockUser, ":id", id),
		method: "PUT",
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

// GetUserPermissions get user's permission
// 获取用户权限，注意不包含用户角色下的权限
func (c *Client) GetUserPermissions(id string) ([]types.UserPermissionInfo, error) {
	if id == "" {
		return nil, errors.New("invalid user id")
	}

	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathUserPermissionsInfo, ":id", id),
		method: "GET",
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data []types.UserPermissionInfo `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// GetPermissionsWithoutUser 获取用户可添加的权限列表
func (c *Client) GetPermissionsWithoutUser(id string,
	req types.UserWithoutPermissionRequest) ([]types.UserWithoutPermissionResp, error) {

	if req.ResourceID == "" || id == "" {
		return nil, errors.New("invalid user id or resource id")
	}
	query := make(map[string][]string)
	query["resource_id"] = []string{req.ResourceID}
	if req.Scope != "" {
		query["scope"] = []string{req.Scope}
	}

	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathUserWithoutPermissionInfo, ":id", id),
		method: "GET",
		Query:  query,
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data []types.UserWithoutPermissionResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// GetUserRoles get user's roles
// 获取用户的角色
func (c *Client) GetUserRoles(id string) ([]types.UserRoleInfoResp, error) {
	if id == "" {
		return nil, errors.New("invalid user id")
	}

	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathUserRolesInfo, ":id", id),
		method: "GET",
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data []types.UserRoleInfoResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// GetRolesWithoutUser  查询用户可以添加的角色列表
func (c *Client) GetRolesWithoutUser(id string,
	req types.UserWithoutRoleRequest) ([]types.UserQueryRoleResp, error) {

	if id == "" {
		return nil, errors.New("invalid user id")
	}
	query := make(map[string][]string)
	if req.Name != "" {
		query["name"] = []string{req.Name}
	}

	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathUserWithoutRolesInfo, ":id", id),
		method: "GET",
		Query:  query,
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data []types.UserQueryRoleResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// GetUserGrantApps get user grant user
// 获取用户已授权的应用
func (c *Client) GetUserGrantApps(id string) ([]types.UserAppInfoResp, error) {

	if id == "" {
		return nil, errors.New("invalid user id")
	}
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathUserApps, ":id", id),
		method: "GET",
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data []types.UserAppInfoResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// UserAssignPermissions user assign permissions
// 为用户授予权限
func (c *Client) UserAssignPermissions(id string, req types.UserAssignPermissionReq) error {
	if len(req.PermissionIDs) == 0 {
		return errors.New("invalid params")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathUserAssignPermissions, ":id", id),
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

// UserRemovePermission user remove permission
// 移除用户权限
func (c *Client) UserRemovePermission(userID, permissionID string) error {
	if userID == "" || permissionID == "" {
		return errors.New("invalid params")
	}
	path := strings.ReplaceAll(pathUserRemovePermissions, ":pid", permissionID)
	path = strings.ReplaceAll(path, ":id", userID)
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

// UserAssignRoles user assign roles
// 为用户授予角色
func (c *Client) UserAssignRoles(id string, req types.UserAssignRolesReq) error {
	if len(req.RoleIDs) == 0 {
		return errors.New("invalid params")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathUserAssignRoles, ":id", id),
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

// UserRemoveRole user remove role
// 移除用户角色
func (c *Client) UserRemoveRole(userID, roleID string) error {
	if userID == "" || roleID == "" {
		return errors.New("invalid params")
	}
	path := strings.ReplaceAll(pathUserRemoveRoles, ":rid", roleID)
	path = strings.ReplaceAll(path, ":id", userID)
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

// UserRemoveApp user remove app
// 移除用户授权应用
func (c *Client) UserRemoveApp(userID, appID string) error {
	if userID == "" || appID == "" {
		return errors.New("invalid params")
	}
	path := strings.ReplaceAll(pathUserRemoveApp, ":aid", appID)
	path = strings.ReplaceAll(path, ":id", userID)
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
