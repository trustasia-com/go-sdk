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
	AuthorizeStatusSuccess = "success" // 成功
	AuthorizeStatusInit    = "init"    // 初始化
	AuthorizeStatusBind    = "bind"    // 绑定
	AuthorizeStatusRefuse  = "refuse"  // 拒绝
	AuthorizeStatusFail    = "fail"    // 失败
)

// api list
const (
	apiCosignPreBind          = "/ta-app/sdk/cosign/pre-bind"
	apiCosignBind             = "/ta-app/sdk/cosign/bind?msg_id=%s"
	apiCosignPreLogin         = "/ta-app/sdk/cosign/pre-login"
	apiCosignLogin            = "/ta-app/sdk/cosign/login?msg_id=%s"
	apiCosignPreSign          = "/ta-app/sdk/cosign/pre-sign"
	apiCosignSign             = "/ta-app/sdk/cosign/sign?msg_id=%s"
	apiCosignCredentials      = "/ta-app/sdk/cosign/credentials?rp_user_id=%s&page=%d&size=%d"
	apiCosignCredentialDelete = "/ta-app/sdk/cosign/credentials"
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

// BindQRCode 获取注册扫描二维码
func (a *App) BindQRCode(req BindQRCodeReq) (*BindQRCodeResp, error) {
	if req.RpUserDisplayName == "" {
		return nil, errors.New("No DisplayName found")
	}
	if req.RpUserID == "" {
		return nil, errors.New("No UserID found")
	}
	if !req.ClientType.isValid() {
		return nil, errors.New("ClientType is invalid")
	}

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	scope := "app/"
	msg, err := a.client.Request(http.MethodPost, apiCosignPreBind, scope, data)
	if err != nil {
		return nil, err
	}
	resp := &BindQRCodeResp{}
	err = json.Unmarshal(msg.Data, resp)
	return resp, err
}

// BindResult 获取扫描认证结果
func (a *App) BindResult(req BindResultReq, callback AuthOKCallback) (*BindResultResp, error) {
	if req.MsgID == "" {
		return nil, errors.New("No MsgID found")
	}

	path := fmt.Sprintf(apiCosignBind, req.MsgID)
	scope := "app/"
	msg, err := a.client.Request(http.MethodGet, path, scope, nil)
	if err != nil {
		return nil, err
	}
	resp := &BindResultResp{}
	err = json.Unmarshal(msg.Data, resp)
	return resp, err
}

// LoginQRCode 认证请求
func (a *App) LoginQRCode(req LoginQRCodeReq) (*LoginQRCodeResp, error) {
	if !req.ClientType.isValid() {
		return nil, errors.New("ClientType is invalid")
	}

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	scope := "app/"
	msg, err := a.client.Request(http.MethodPost, apiCosignPreLogin, scope, data)
	if err != nil {
		return nil, err
	}
	resp := &LoginQRCodeResp{}
	err = json.Unmarshal(msg.Data, resp)
	return resp, err
}

// LoginResult 获取认证结果
func (a *App) LoginResult(req LoginResultReq, callback AuthOKCallback) (*LoginResultResp, error) {
	if req.MsgID == "" {
		return nil, errors.New("No MsgID found")
	}

	path := fmt.Sprintf(apiCosignLogin, req.MsgID)
	scope := "app/"
	msg, err := a.client.Request(http.MethodGet, path, scope, nil)
	if err != nil {
		return nil, err
	}
	resp := &LoginResultResp{}
	err = json.Unmarshal(msg.Data, resp)
	if err != nil {
		return nil, err
	}
	if resp.Status == AuthorizeStatusSuccess {
		err = callback(resp.RpUserID)
	}
	return resp, err
}

// SignRequest 签名请求
func (a *App) SignRequest(req SignRequestReq) (*SignRequestResp, error) {
	if !req.ClientType.isValid() {
		return nil, errors.New("ClientType is invalid")
	}
	if len(req.Data) == 0 {
		return nil, errors.New("No Data found")
	}

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	scope := "app/"
	msg, err := a.client.Request(http.MethodPost, apiCosignPreSign, scope, data)
	if err != nil {
		return nil, err
	}
	resp := &SignRequestResp{}
	err = json.Unmarshal(msg.Data, resp)
	return resp, err
}

// SignResult 签名结果
func (a *App) SignResult(req SignResultReq) (*SignResultResp, error) {
	if req.MsgID == "" {
		return nil, errors.New("No MsgID found")
	}

	path := fmt.Sprintf(apiCosignSign, req.MsgID)
	scope := "app/"
	msg, err := a.client.Request(http.MethodGet, path, scope, nil)
	if err != nil {
		return nil, err
	}
	resp := &SignResultResp{}
	err = json.Unmarshal(msg.Data, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// CosignCredentials 用户协同凭证列表
func (a *App) CosignCredentials(req CredentialsReq) (*CredentialsResp, error) {
	if req.RpUserID == "" {
		return nil, errors.New("Need specify req.RpUserID")
	}
	if req.Page < 1 {
		req.Page = 1
	}
	if req.Size < 1 {
		req.Size = 10
	}
	path := fmt.Sprintf(apiCosignCredentials, req.RpUserID, req.Page, req.Size)
	scope := "app/"
	msg, err := a.client.Request(http.MethodGet, path, scope, nil)
	if err != nil {
		return nil, err
	}
	resp := &CredentialsResp{}
	err = json.Unmarshal(msg.Data, resp)
	return resp, err
}

// CredentialDelete 删除协同凭证
func (a *App) CredentialDelete(req CredentialDeleteReq) (*CredentialDeleteResp, error) {
	if req.CertID == "" {
		return nil, errors.New("No CertID found")
	}
	if req.RpUserID == "" {
		return nil, errors.New("No RpUserID found")
	}
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	scope := "app/"
	msg, err := a.client.Request(http.MethodDelete, apiCosignCredentialDelete, scope, data)
	if err != nil {
		return nil, err
	}
	resp := &CredentialDeleteResp{}
	err = json.Unmarshal(msg.Data, resp)
	return resp, err
}
