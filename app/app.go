// Package app provides ...
package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/trustasia-com/go-sdk/pkg/client"
	"github.com/trustasia-com/go-sdk/pkg/credentials"
)

// 认证结果状态
const (
	AuthorizeStatusSuccess = "success"
)

// api list
const (
	// app管理
	apiCreateOrUpdateApp = "/ta-app/rp/app"
	apiDeleteApp         = "/ta-app/rp/app"

	apiRegQRCode     = "/ta-app/rp/attestation/options"
	apiRegResult     = "/ta-app/rp/attestation/result/%s"
	apiAuthRequest   = "/ta-app/rp/assertion/options"
	apiAuthResult    = "/ta-app/rp/assertion/result/%s"
	apiCredentials   = "/ta-app/rp/credentials?user_id=%s"
	apiCredentialDel = "/ta-app/rp/credentials"
	apiCosignInfo    = "/ta-app/rp/cosign?slug=%s&rp_user_id=%s"
	apiCosignDelete  = "/ta-app/rp/cosign"
)

// App instance for wekey rp
type App struct {
	client *client.HTTPClient
}

// AuthOKCallback 认证成功回调
type AuthOKCallback func(userID string) error

// New wekey client
func New(sess *credentials.Session) *App {
	return &App{
		client: client.NewHTTPClient(sess),
	}
}

// CreateOrUpdateApp 创建应用
func (a *App) CreateOrUpdateApp(req CreateOrUpdateAppReq) error {
	if req.Slug == "" {
		return errors.New("Need specify req.Slug")
	}
	if req.Name == "" {
		return errors.New("Need specify req.Name")
	}
	if req.RpInfo.RpID == "" {
		return errors.New("Need specify req.RpInfo.RpID")
	}
	if len(req.RpInfo.Origins) == 0 {
		return errors.New("Need specify req.RpInfo.Origins")
	}
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	scope := "app/"
	_, err = a.client.Request(http.MethodPost, apiCreateOrUpdateApp, scope, data)
	return err
}

// DeleteApp 更新应用
func (a *App) DeleteApp(req DeleteAppReq) error {
	if req.Slug == "" {
		return errors.New("Need specify req.Slug")
	}
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	scope := "app/"
	_, err = a.client.Request(http.MethodDelete, apiDeleteApp, scope, data)
	return err
}

// RegQRCode 获取注册扫描二维码
func (a *App) RegQRCode(req RegQRCodeReq) (*RegQRCodeResp, error) {
	if req.Slug == "" {
		return nil, errors.New("Need specify req.Slug")
	}
	if req.CredentialName == "" {
		return nil, errors.New("Need specify req.CredentialName")
	}

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	scope := "app/"
	msg, err := a.client.Request(http.MethodPost, apiRegQRCode, scope, data)
	if err != nil {
		return nil, err
	}
	resp := &RegQRCodeResp{}
	err = json.Unmarshal(msg.Data, resp)
	return resp, err
}

// RegResult 获取扫描认证结果
func (a *App) RegResult(req RegResultReq, callback AuthOKCallback) (*RegResultResp, error) {
	if req.MsgID == "" {
		return nil, errors.New("Need specify req.MsgID")
	}

	path := fmt.Sprintf(apiRegResult, req.MsgID[1:])
	scope := "app/"
	msg, err := a.client.Request(http.MethodGet, path, scope, nil)
	if err != nil {
		return nil, err
	}
	resp := &RegResultResp{}
	err = json.Unmarshal(msg.Data, resp)
	return resp, err
}

// AuthRequest 认证请求
func (a *App) AuthRequest(req AuthRequestReq) (*AuthRequestResp, error) {
	if req.Type != TypeFidoScan && req.Type != TypeCosignScan {
		return nil, errors.New("Invalid Auth Method")
	}
	if req.Slug == "" {
		return nil, errors.New("Need specify req.Slug")
	}

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	scope := "app/"
	msg, err := a.client.Request(http.MethodPost, apiAuthRequest, scope, data)
	if err != nil {
		return nil, err
	}
	resp := &AuthRequestResp{}
	err = json.Unmarshal(msg.Data, resp)
	return resp, err
}

// AuthResult 获取认证结果
func (a *App) AuthResult(req AuthResultReq, callback AuthOKCallback) (*AuthResultResp, error) {
	if req.MsgID == "" {
		return nil, errors.New("Need specify req.MsgID")
	}

	path := fmt.Sprintf(apiAuthResult, req.MsgID[1:])
	scope := "app/"
	msg, err := a.client.Request(http.MethodGet, path, scope, nil)
	if err != nil {
		return nil, err
	}
	resp := &AuthResultResp{}
	err = json.Unmarshal(msg.Data, resp)
	if err != nil {
		return nil, err
	}
	if resp.Status == AuthorizeStatusSuccess {
		err = callback(resp.RpUserID)
	}
	return resp, err
}

// UserCredentials 用户凭证列表
func (a *App) UserCredentials(req UserCredentialsReq) (*UserCredentialsResp, error) {
	if req.RpUserID == "" {
		return nil, errors.New("Need specify req.UserID")
	}

	path := fmt.Sprintf(apiCredentials, req.RpUserID)
	scope := "app/"
	msg, err := a.client.Request(http.MethodGet, path, scope, nil)
	if err != nil {
		return nil, err
	}
	resp := &UserCredentialsResp{}
	err = json.Unmarshal(msg.Data, resp)
	return resp, err
}

// DeleteCredential 删除用户凭证
func (a *App) DeleteCredential(req DeleteCredentialReq) (*DeleteCredentialResp, error) {
	if req.RpUserID == "" {
		return nil, errors.New("Need specify req.UserID")
	}
	if len(req.CredentialIDs) == 0 {
		return nil, errors.New("Need specify req.CredID")
	}

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	scope := "app/"
	msg, err := a.client.Request(http.MethodDelete, apiCredentialDel, scope, data)
	if err != nil {
		return nil, err
	}
	resp := &DeleteCredentialResp{}
	err = json.Unmarshal(msg.Data, resp)
	return resp, err
}

// UserCosignList 用户协同凭证列表
func (a *App) UserCosignList(req CosignListReq) (*CosignListResp, error) {
	if req.RpUserID == "" {
		return nil, errors.New("Need specify req.RpUserID")
	}
	path := fmt.Sprintf(apiCosignInfo, req.Slug, req.RpUserID)
	scope := "app/"
	msg, err := a.client.Request(http.MethodGet, path, scope, nil)
	if err != nil {
		return nil, err
	}
	resp := &CosignListResp{}
	err = json.Unmarshal(msg.Data, resp)
	return resp, err
}

// UserCosignDelete 删除协同凭证
func (a *App) UserCosignDelete(req CosignDeleteReq) (*CosignDeleteResp, error) {
	if req.RpUserID == "" {
		return nil, errors.New("Need specify req.RpUserID")
	}
	if req.UserID == "" {
		return nil, errors.New("Need specify req.UserID")
	}
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	scope := "app/"
	msg, err := a.client.Request(http.MethodDelete, apiCosignDelete, scope, data)
	if err != nil {
		return nil, err
	}
	resp := &CosignDeleteResp{}
	err = json.Unmarshal(msg.Data, resp)
	return resp, err
}
