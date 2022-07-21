package management

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/trustasia-com/go-sdk/wekey/management/types"
)

// GetIdpGithubConfig github
// 获取GitHub身份源
func (c *Client) GetIdpGithubConfig() (*types.IdpGithubInfoResp, error) {
	path := strings.ReplaceAll(pathGetIdpConfigInfo, ":provider", "github")
	path = strings.ReplaceAll(path, ":identifier", "github")
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
		Data *types.IdpGithubInfoResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// AddIdpGithub add github
// 添加GitHub身份源
func (c *Client) AddIdpGithub(req types.AddIdpGithubRequest) error {
	if req.Name == "" || req.ClientID == "" || req.ClientSecret == "" {
		return errors.New("invalid params")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathCreateIdp, ":provider", "github"),
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

// UpdateIdpGithub update GitHub conf
// 更新GitHub身份源配置
func (c *Client) UpdateIdpGithub(req types.UpdateIdpGithub) error {
	if req.Name == "" || req.ClientID == "" || req.ClientSecret == "" {
		return errors.New("invalid params")
	}
	data, _ := json.Marshal(req)
	path := strings.ReplaceAll(pathUpdateIdp, ":provider", "github")
	path = strings.ReplaceAll(path, ":identifier", "github")
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

// DelIdpGithub del github
// 删除GitHub身份源
func (c *Client) DelIdpGithub() error {
	path := strings.ReplaceAll(pathDelIdp, ":provider", "github")
	path = strings.ReplaceAll(path, ":identifier", "github")
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

// GetIdpGitlabConfig gitlab
// 获取Gitlab身份源
func (c *Client) GetIdpGitlabConfig() (*types.IdpGitlabInfoResp, error) {
	path := strings.ReplaceAll(pathGetIdpConfigInfo, ":provider", "gitlab")
	path = strings.ReplaceAll(path, ":identifier", "gitlab")
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
		Data *types.IdpGitlabInfoResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// AddIdpGitlab add gitlab
// 添加gitlab身份源
func (c *Client) AddIdpGitlab(req types.AddIdpGitlabRequest) error {
	if req.Name == "" || req.ClientID == "" || req.ClientSecret == "" {
		return errors.New("invalid params")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathCreateIdp, ":provider", "gitlab"),
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

// UpdateIdpGitlab update gitlab conf
// 更新gitlab配置
func (c *Client) UpdateIdpGitlab(req types.UpdateIdpGitlab) error {
	if req.Name == "" || req.ClientID == "" || req.ClientSecret == "" {
		return errors.New("invalid params")
	}
	data, _ := json.Marshal(req)
	path := strings.ReplaceAll(pathUpdateIdp, ":provider", "gitlab")
	path = strings.ReplaceAll(path, ":identifier", "gitlab")
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

// DelIdpGitlab del gitlab
// 删除gitlab
func (c *Client) DelIdpGitlab() error {
	path := strings.ReplaceAll(pathDelIdp, ":provider", "gitlab")
	path = strings.ReplaceAll(path, ":identifier", "gitlab")
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

// GetIdpWeiboConfig weibo
// 获取微博身份源配置
func (c *Client) GetIdpWeiboConfig() (*types.IdpWeiboInfoResp, error) {
	path := strings.ReplaceAll(pathGetIdpConfigInfo, ":provider", "weibo")
	path = strings.ReplaceAll(path, ":identifier", "weibo")
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
		Data *types.IdpWeiboInfoResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// AddIdpWeibo add weibo
// 添加微博身份源
func (c *Client) AddIdpWeibo(req types.AddIdpWeiboRequest) error {
	if req.Name == "" || req.ClientID == "" || req.ClientSecret == "" {
		return errors.New("invalid params")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathCreateIdp, ":provider", "weibo"),
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

// UpdateIdpWeibo update weibo conf
// 更新微博配置
func (c *Client) UpdateIdpWeibo(req types.UpdateIdpWeibo) error {
	if req.Name == "" || req.ClientID == "" || req.ClientSecret == "" {
		return errors.New("invalid params")
	}
	data, _ := json.Marshal(req)
	path := strings.ReplaceAll(pathUpdateIdp, ":provider", "weibo")
	path = strings.ReplaceAll(path, ":identifier", "weibo")
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

// DelIdpWeibo del weibo
// 删除微博身份源
func (c *Client) DelIdpWeibo() error {
	path := strings.ReplaceAll(pathDelIdp, ":provider", "weibo")
	path = strings.ReplaceAll(path, ":identifier", "weibo")
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

// GetIdpGitEEConfig gitee
// 获取gitee身份源配置
func (c *Client) GetIdpGitEEConfig() (*types.IdpGitEEInfoResp, error) {
	path := strings.ReplaceAll(pathGetIdpConfigInfo, ":provider", "gitee")
	path = strings.ReplaceAll(path, ":identifier", "gitee")
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
		Data *types.IdpGitEEInfoResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// AddIdpGitEE add gitee
// 添加gitee身份源
func (c *Client) AddIdpGitEE(req types.AddIdpGitEERequest) error {
	if req.Name == "" || req.ClientID == "" || req.ClientSecret == "" {
		return errors.New("invalid params")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathCreateIdp, ":provider", "gitee"),
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

// UpdateIdpGitEE update gitee conf
// 更新gitee配置
func (c *Client) UpdateIdpGitEE(req types.UpdateIdpGitEE) error {
	if req.Name == "" || req.ClientID == "" || req.ClientSecret == "" {
		return errors.New("invalid params")
	}
	data, _ := json.Marshal(req)
	path := strings.ReplaceAll(pathUpdateIdp, ":provider", "gitee")
	path = strings.ReplaceAll(path, ":identifier", "gitee")
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

// DelIdpGitEE del gitee
// 删除gitee身份源
func (c *Client) DelIdpGitEE() error {
	path := strings.ReplaceAll(pathDelIdp, ":provider", "gitee")
	path = strings.ReplaceAll(path, ":identifier", "gitee")
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

// GetIdpBaiduConfig 百度
// 获取百度身份源配置
func (c *Client) GetIdpBaiduConfig() (*types.IdpBaiduInfoResp, error) {
	path := strings.ReplaceAll(pathGetIdpConfigInfo, ":provider", "baidu")
	path = strings.ReplaceAll(path, ":identifier", "baidu")
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
		Data *types.IdpBaiduInfoResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// AddIdpBaidu add 百度
// 添加百度身份源配置
func (c *Client) AddIdpBaidu(req types.AddIdpBaiduRequest) error {
	if req.Name == "" || req.ClientID == "" || req.ClientSecret == "" {
		return errors.New("invalid params")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathCreateIdp, ":provider", "baidu"),
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

// UpdateIdpBaidu update 百度 conf
// 更细百度身份源
func (c *Client) UpdateIdpBaidu(req types.UpdateIdpBaidu) error {
	if req.Name == "" || req.ClientID == "" || req.ClientSecret == "" {
		return errors.New("invalid params")
	}
	data, _ := json.Marshal(req)
	path := strings.ReplaceAll(pathUpdateIdp, ":provider", "baidu")
	path = strings.ReplaceAll(path, ":identifier", "baidu")
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

// DelIdpBaidu del baidu
// 删除百度身份源
func (c *Client) DelIdpBaidu() error {
	path := strings.ReplaceAll(pathDelIdp, ":provider", "baidu")
	path = strings.ReplaceAll(path, ":identifier", "baidu")
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

// GetIdpQQConfig qq
// 获取qq身份源配置
func (c *Client) GetIdpQQConfig() (*types.IdpQQInfoResp, error) {
	path := strings.ReplaceAll(pathGetIdpConfigInfo, ":provider", "qq")
	path = strings.ReplaceAll(path, ":identifier", "qq")
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
		Data *types.IdpQQInfoResp `json:"data"`
	}
	if err := json.Unmarshal(res.Body, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", resp.Code, resp.Error)
	}
	return resp.Data, nil
}

// AddIdpQQ add qq
// 添加qq身份源
func (c *Client) AddIdpQQ(req types.AddIdpQQRequest) error {
	if req.Name == "" || req.ClientID == "" || req.ClientSecret == "" {
		return errors.New("invalid params")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   strings.ReplaceAll(pathCreateIdp, ":provider", "qq"),
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

// UpdateIdpQQ update qq conf
// 更新qq身份源
func (c *Client) UpdateIdpQQ(req types.UpdateIdpQQ) error {
	if req.Name == "" || req.ClientID == "" || req.ClientSecret == "" {
		return errors.New("invalid params")
	}
	data, _ := json.Marshal(req)
	path := strings.ReplaceAll(pathUpdateIdp, ":provider", "qq")
	path = strings.ReplaceAll(path, ":identifier", "qq")
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

// DelIdpQQ del qq
// 删除qq身份源
func (c *Client) DelIdpQQ() error {
	path := strings.ReplaceAll(pathDelIdp, ":provider", "qq")
	path = strings.ReplaceAll(path, ":identifier", "qq")
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
