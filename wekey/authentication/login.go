package authentication

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/trustasia-com/go-sdk/pkg/valid"
	"github.com/trustasia-com/go-sdk/wekey/authentication/types"
)

const (
	pathLoginByPassword = "/ta-auth/login/password"
	pathSendPasscode    = "/ta-auth/passcode"
	pathLoginByPasscode = "/ta-auth/login/passcode"
)

// LoginByPassword 密码登录
func (c *Client) LoginByPassword(req types.LoginByPasswordRequest) (*types.LoginResponse, error) {
	req.Password = c.encrypt(req.Password)
	req.ClientID = c.ClientID
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   pathLoginByPassword,
		method: "POST",
		Body:   data,
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var loginResp struct {
		types.BaseResponse
		Data *types.LoginResponse `json:"data"`
	}

	if err := json.Unmarshal(res.Body, &loginResp); err != nil {
		return nil, err
	}
	if loginResp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", loginResp.Code, loginResp.Error)
	}
	return loginResp.Data, nil
}

// SendPasscode 发送验证码
// 目前支持的Action是登录(login)，注册(register)，忘记密码(reset-password)
func (c *Client) SendPasscode(req types.SendPasscodeRequest) error {
	// 前置校验一下是否合法
	if !valid.IsASCIIEmail(req.Target) && !valid.IsChinaPhoneNo(req.Target) {
		return errors.New("invalid phone number or email address")
	}
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   pathSendPasscode,
		method: "POST",
		Body:   data,
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return res.Err
	}
	var passcodeResp struct {
		types.BaseResponse
	}

	if err := json.Unmarshal(res.Body, &passcodeResp); err != nil {
		return err
	}
	if passcodeResp.Code != 0 {
		return fmt.Errorf("code:[%d],error:%s", passcodeResp.Code, passcodeResp.Error)
	}
	return nil
}

// LoginByPasscode 验证码登录
func (c *Client) LoginByPasscode(req types.LoginByPasscodeRequest) (*types.LoginResponse, error) {

	// 前置校验一下是否合法
	if !valid.IsASCIIEmail(req.Target) && !valid.IsChinaPhoneNo(req.Target) {
		return nil, errors.New("invalid phone number or email address")
	}
	req.ClientID = c.ClientID
	data, _ := json.Marshal(req)
	httpReq := &ClientRequest{
		Path:   pathLoginByPasscode,
		method: "POST",
		Body:   data,
	}
	res := c.httpRequest(httpReq)
	if res.Err != nil {
		return nil, res.Err
	}
	var loginResp struct {
		types.BaseResponse
		Data *types.LoginResponse `json:"data"`
	}

	if err := json.Unmarshal(res.Body, &loginResp); err != nil {
		return nil, err
	}
	if loginResp.Code != 0 {
		return nil, fmt.Errorf("code:[%d],error:%s", loginResp.Code, loginResp.Error)
	}
	return loginResp.Data, nil
}
