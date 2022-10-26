// Package wekey provides ...
package wekey

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/trustasia-com/go-sdk/pkg/client"
	"github.com/trustasia-com/go-sdk/pkg/credentials"
)

// api list
const (
	apiRegQRCode   = "/ta-app/attestation/options"
	apiRegResult   = "/ta-app/attestation/result/%s"
	apiAuthRequest = "/ta-app/assertion/options"
	apiAuthResult  = "/ta-app/assertion/result/%s"
)

// WeKey instance for wekey rp
type WeKey struct {
	client *client.HTTPClient
}

// New wekey client
func New(sess *credentials.Session) *WeKey {
	return &WeKey{
		client: client.NewHTTPClient(sess),
	}
}

// RegQRCode 获取注册扫描二维码
func (we *WeKey) RegQRCode(req RegQRCodeReq) (*RegQRCodeResp, error) {
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
	scope := "wekey/"
	msg, err := we.client.Request(http.MethodPost, apiRegQRCode, scope, data)
	if err != nil {
		return nil, err
	}
	return &RegQRCodeResp{QRCode: string(msg.Data)}, nil
}

// RegResult 获取扫描认证结果
func (we *WeKey) RegResult(req RegResultReq) (*RegResultResp, error) {
	if req.MsgID == "" {
		return nil, errors.New("Need specify req.MsgID")
	}

	path := fmt.Sprintf(apiRegResult, req.MsgID)
	scope := "wekey/"
	msg, err := we.client.Request(http.MethodPost, path, scope, nil)
	if err != nil {
		return nil, err
	}
	resp := &RegResultResp{}
	err = json.Unmarshal(msg.Data, resp)
	return resp, err
}

// AuthRequest 认证请求
func (we *WeKey) AuthRequest(req AuthRequestReq) (*AuthRequestResp, error) {
	if req.Method != AuthMethodQRCode && req.Method == AuthMethodPush {
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
	scope := "wekey/"
	msg, err := we.client.Request(http.MethodPost, apiAuthResult, scope, data)
	if err != nil {
		return nil, err
	}
	resp := &AuthRequestResp{}
	err = json.Unmarshal(msg.Data, resp)
	return resp, err
}

// AuthResult 获取认证结果
func (we *WeKey) AuthResult(req AuthResultReq) (*AuthResultResp, error) {
	if req.MsgID == "" {
		return nil, errors.New("Need specify req.MsgID")
	}

	path := fmt.Sprintf(apiAuthResult, req.MsgID)
	scope := "wekey/"
	msg, err := we.client.Request(http.MethodGet, path, scope, nil)
	if err != nil {
		return nil, err
	}
	resp := &AuthResultResp{}
	err = json.Unmarshal(msg.Data, resp)
	return resp, err
}
