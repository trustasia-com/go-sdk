package management

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/trustasia-com/go-sdk/wekey/management/types"
)

const (
	pathQueryApp          = "/ta-wekey-dash/apps"
	pathAppInfo           = "/ta-wekey-dash/apps"
	pathCreateApp         = "/ta-wekey-dash/apps"
	pathDeleteApp         = "/ta-wekey-dash/apps"
	pathAppIDPs           = "/ta-wekey-dash/apps/:id/idps"
	pathUpdateAppBasic    = "/ta-wekey-dash/apps/:id/basic"
	pathUpdateAppGrant    = "/ta-wekey-dash/apps/:id/grant"
	pathUpdateAppAdvance  = "/ta-wekey-dash/apps/:id/advance"
	pathUpdateAppSecret   = "/ta-wekey-dash/apps/:id/secret"
	pathQueryResource     = "/ta-wekey-dash/resources"
	pathResourceInfo      = "/ta-wekey-dash/resources"
	pathCreateResource    = "/ta-wekey-dash/resources"
	pathDeleteResource    = "/ta-wekey-dash/resources"
	pathUpdateResource    = "/ta-wekey-dash/resources"
	pathQueryPermission   = "/ta-wekey-dash/permissions"
	pathCreatePermission  = "/ta-wekey-dash/permissions"
	pathDeletePermission  = "/ta-wekey-dash/permissions"
	pathResourceApps      = "/ta-wekey-dash/resources/:id/apps"
	pathAppResources      = "/ta-wekey-dash/apps/:id/resources"
	pathOpenAppResources  = "/ta-wekey-dash/apps/:id/resources/:resourceId/open"
	pathCloseAppResources = "/ta-wekey-dash/apps/:id/resources/:resourceId/close"
	pathOpenAppIDP        = "/ta-wekey-dash/apps/:id/conns/:provider/:identifier/open"
	pathCloseAppIDP       = "/ta-wekey-dash/apps/:id/conns/:provider/:identifier/close"
)

// QueryApp query
// 查询应用列表
func (c *Client) QueryApp(req types.QueryAppRequest) ([]types.QueryAppResp, error) {
	query := make(map[string][]string)
	if req.Name != "" {
		query["name"] = []string{req.Name}
	}
	httpReq := &ClientRequest{
		Path:   pathQueryApp,
		method: "GET",
		Query:  query,
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data []types.QueryAppResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// AppInfo info
// 查询应用详情
func (c *Client) AppInfo(id string) (*types.AppInfoResp, error) {
	httpReq := &ClientRequest{
		Path:   pathAppInfo + "/" + id,
		method: "GET",
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data *types.AppInfoResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// CreateApp create app
// 创建一个应用
func (c *Client) CreateApp(req types.CreateAppRequest) (string, error) {

	if req.Name == "" {
		return "", errors.New("invalid app name")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   pathCreateApp,
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

// DelApp del app
// 删除一个应用
func (c *Client) DelApp(id string) error {
	if id == "" {
		return errors.New("invalid app id")
	}
	httpReq := &ClientRequest{
		Path:   pathDeleteApp + "/" + id,
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

// GetAppIDPs get app idps
// 获取应用的身份源信息
func (c *Client) GetAppIDPs(id string) ([]types.AppIDPResp, error) {

	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathAppIDPs, ":id", id),
		method: "GET",
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data []types.AppIDPResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// UpdateAppBasic update app basic info
// 更新应用基础信息
func (c *Client) UpdateAppBasic(id string, req types.UpdateAppBasicRequest) error {
	if id == "" {
		return errors.New("invalid app id")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathUpdateAppBasic, ":id", id),
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

// UpdateAppGrant update app grant
// 更新应用授权信息
func (c *Client) UpdateAppGrant(id string, req types.UpdateAppGrantRequest) error {
	if id == "" {
		return errors.New("invalid app id")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathUpdateAppGrant, ":id", id),
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

// UpdateAppAdvance update app advance
// 更新应用高级信息
func (c *Client) UpdateAppAdvance(id string, req types.UpdateAppAdvanceRequest) error {
	if id == "" {
		return errors.New("invalid app id")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathUpdateAppAdvance, ":id", id),
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

// UpdateAppSecret update app secret
// 更新应用密钥
func (c *Client) UpdateAppSecret(id string) error {
	if id == "" {
		return errors.New("invalid app id")
	}
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathUpdateAppSecret, ":id", id),
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

// QueryResource query
// 查询资源组列表
func (c *Client) QueryResource(req types.QueryResourceRequest) ([]types.QueryResourceResp, error) {
	query := make(map[string][]string)
	if req.Name != "" {
		query["name"] = []string{req.Name}
	}
	httpReq := &ClientRequest{
		Path:   pathQueryResource,
		method: "GET",
		Query:  query,
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data []types.QueryResourceResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// ResourceInfo info
// 查询资源组详情
func (c *Client) ResourceInfo(id string) (*types.ResourceInfoResp, error) {
	httpReq := &ClientRequest{
		Path:   pathResourceInfo + "/" + id,
		method: "GET",
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data *types.ResourceInfoResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// CreateResource create resource
// 创建一个资源组
func (c *Client) CreateResource(req types.CreateResourceRequest) (string, error) {

	if req.Name == "" {
		return "", errors.New("invalid resource name")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   pathCreateResource,
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

// DelResource del resource
// 删除一个资源组
func (c *Client) DelResource(id string) error {
	if id == "" {
		return errors.New("invalid resource id")
	}
	httpReq := &ClientRequest{
		Path:   pathDeleteResource + "/" + id,
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

// UpdateResource update resource info
// 更新资源组信息
func (c *Client) UpdateResource(id string, req types.UpdateResourceRequest) error {
	if id == "" {
		return errors.New("invalid resource id")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   pathUpdateResource + "/" + id,
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

// QueryPermission query permission
// 查询权限列表
func (c *Client) QueryPermission(req types.QueryPermissionRequest) ([]types.QueryPermissionResp, error) {
	if req.ResourceID == "" {
		return nil, errors.New("invalid resource id")
	}
	query := make(map[string][]string)
	query["resource_id"] = []string{req.ResourceID}
	if req.Scope != "" {
		query["scope"] = []string{req.Scope}
	}
	httpReq := &ClientRequest{
		Path:   pathQueryPermission,
		method: "GET",
		Query:  query,
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data []types.QueryPermissionResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// CreatePermission create permission
// 添加权限
func (c *Client) CreatePermission(req types.CreatePermissionRequest) (string, error) {

	if req.ResourceID == "" || req.Scope == "" {
		return "", errors.New("invalid resource id or scope")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   pathCreatePermission,
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

// DelPermission del permission
// 删除权限
func (c *Client) DelPermission(id string) error {
	if id == "" {
		return errors.New("invalid permission id")
	}
	httpReq := &ClientRequest{
		Path:   pathDeletePermission + "/" + id,
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

// QueryResourceApps query resource apps
// 查询资源组关联的应用
func (c *Client) QueryResourceApps(resourceID string) ([]types.ResourceAppsResp, error) {

	if resourceID == "" {
		return nil, errors.New("invalid resource id")
	}
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathResourceApps, ":id", resourceID),
		method: "GET",
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data []types.ResourceAppsResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// QueryAppResources query app resources
// 查询应用关联的资源组
func (c *Client) QueryAppResources(appID string) ([]types.AppResourcesResp, error) {

	if appID == "" {
		return nil, errors.New("invalid app id")
	}
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathAppResources, ":id", appID),
		method: "GET",
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var resp struct {
		types.BaseResponse
		Data []types.AppResourcesResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// OpenAppResource open app with resource
// 开启应用资源组的关联
func (c *Client) OpenAppResource(appID, resourceID string) error {
	if appID == "" || resourceID == "" {
		return errors.New("invalid app id or resource id")
	}
	path := strings.ReplaceAll(pathOpenAppResources, ":id", appID)
	path = strings.ReplaceAll(path, ":resourceId", resourceID)

	httpReq := &ClientRequest{
		Path:   path,
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

// CloseAppResource open app with resource
// 关闭应用资源组关联
func (c *Client) CloseAppResource(appID, resourceID string) error {
	if appID == "" || resourceID == "" {
		return errors.New("invalid app id or resource id")
	}
	path := strings.ReplaceAll(pathCloseAppResources, ":id", appID)
	path = strings.ReplaceAll(path, ":resourceId", resourceID)

	httpReq := &ClientRequest{
		Path:   path,
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

// OpenAppIDP open app identity
// 为应用开启身份源
// provider指身份源类型，如: ldap,qq等
// identifier通常指您设置的唯一标识，社会化身份源的provider和identifier一直，如：qq，github
func (c *Client) OpenAppIDP(appID, provider, identifier string) error {
	if appID == "" || provider == "" || identifier == "" {
		return errors.New("invalid app id or provider or identifier")
	}
	path := strings.ReplaceAll(pathOpenAppIDP, ":id", appID)
	path = strings.ReplaceAll(path, ":provider", provider)
	path = strings.ReplaceAll(path, ":identifier", identifier)

	httpReq := &ClientRequest{
		Path:   path,
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

// CloseAppIDP close app identity
// provider指身份源类型，如: ldap,qq等
// identifier通常指您设置的唯一标识，社会化身份源的provider和identifier一直，如：qq，github
func (c *Client) CloseAppIDP(appID, provider, identifier string) error {
	if appID == "" || provider == "" || identifier == "" {
		return errors.New("invalid app id or provider or identifier")
	}
	path := strings.ReplaceAll(pathCloseAppIDP, ":id", appID)
	path = strings.ReplaceAll(path, ":provider", provider)
	path = strings.ReplaceAll(path, ":identifier", identifier)

	httpReq := &ClientRequest{
		Path:   path,
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
