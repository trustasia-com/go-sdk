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
	apiRegQRCode     = "/ta-app/rp/attestation/options"
	apiRegResult     = "/ta-app/rp/attestation/result/%s"
	apiAuthRequest   = "/ta-app/rp/assertion/options"
	apiAuthResult    = "/ta-app/rp/assertion/result/%s"
	apiCredentials   = "/ta-app/rp/credentials?user_id=%s"
	apiCredentialDel = "/ta-app/rp/credentials"
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

// RegQRCode 获取注册扫描二维码
func (a *App) RegQRCode(req RegQRCodeReq) (*RegQRCodeResp, error) {
	if req.UserID == "" {
		return nil, errors.New("Need specify req.UserID")
	}
	if req.Username == "" {
		return nil, errors.New("Need specify req.Username")
	}
	if req.DisplayName == "" {
		return nil, errors.New("Need specify req.DisplayName")
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
	if req.Method != AuthMethodQRCode && req.Method != AuthMethodPush {
		return nil, errors.New("Invalid Auth Method")
	}
	if req.UserID == "" {
		return nil, errors.New("Need specify req.UserID")
	}
	if req.Username == "" {
		return nil, errors.New("Need specify req.Username")
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
		err = callback(resp.UserID)
	}
	return resp, err
}

// UserCredentials 用户凭证列表
func (a *App) UserCredentials(req UserCredentialsReq) (*UserCredentialsResp, error) {
	if req.UserID == "" {
		return nil, errors.New("Need specify req.UserID")
	}

	path := fmt.Sprintf(apiCredentials, req.UserID)
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
	if req.UserID == "" {
		return nil, errors.New("Need specify req.UserID")
	}
	if len(req.CredentialIDs) == 0 {
		return nil, errors.New("Need specify req.CredID")
	}

	data, err := json.Marshal(req)
	scope := "app/"
	msg, err := a.client.Request(http.MethodDelete, apiCredentialDel, scope, data)
	if err != nil {
		return nil, err
	}
	resp := &DeleteCredentialResp{}
	err = json.Unmarshal(msg.Data, resp)
	return resp, err
}
