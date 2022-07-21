package management

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/trustasia-com/go-sdk/pkg"
	"github.com/trustasia-com/go-sdk/pkg/client"
)

const (
	accessTokenPath  = "/ta-auth/oidc/token"
	pathGetPublicKey = "/ta-auth/public-key"
)

// Client Management client
type Client struct {
	// your management application client id
	ClientID string
	// your management application client secret
	ClientSecret string
	// your auth endpoint e.g: https://wekey.wekey.com
	Endpoint string

	userAgent string
	// storage to used save your token,you can achieve
	storage Storage
	client  *http.Client
}

// NewClient new
// 初始化一个manage client
func NewClient(clientID string, clientSecret string, endpoint string, storage Storage) (*Client, error) {
	c := &Client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     endpoint,
		userAgent:    pkg.BuildUserAgent(),
		client:       client.NewHTTPClient(),
		storage:      storage,
	}
	accessToken, err := GetAccessToken(c.ClientID, c.ClientSecret, c.Endpoint)
	if err != nil {
		return nil, err
	}
	err = c.storage.SetToken(accessToken)
	if err != nil {
		return nil, err
	}
	pub, err := GetPublicKey(endpoint)
	if err != nil {
		return nil, err
	}
	err = c.storage.SetPublicKey(pub)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// GetAccessToken get access token
// 获取一个access token
func GetAccessToken(clientID, clientSecret, endpoint string) (string, error) {
	data := url.Values{
		"client_id":     {clientID},
		"client_secret": {clientSecret},
		"grant_type":    {"client_credentials"},
		"audience":      {endpoint},
	}
	req, err := http.NewRequest("POST", endpoint+accessTokenPath, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	c := client.NewHTTPClient()
	doResp, err := c.Do(req)
	if err != nil {
		return "", err
	}
	resp, err := ioutil.ReadAll(doResp.Body)
	if err != nil {
		return "", err
	}
	var loginResp struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Expire      int    `json:"expire"`
	}
	if err = json.Unmarshal(resp, &loginResp); err != nil {
		return "", err
	}
	return loginResp.AccessToken, nil
}

// GetPublicKey 获取公钥证书
// 获取当前zone的公钥证书
func GetPublicKey(endpoint string) (*rsa.PublicKey, error) {
	req, err := http.NewRequest("GET", endpoint+pathGetPublicKey, nil)
	if err != nil {
		return nil, err
	}

	c := client.NewHTTPClient()
	doResp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	resp, err := ioutil.ReadAll(doResp.Body)
	if err != nil {
		return nil, err
	}
	var publicKeyResp struct {
		Code  int    `json:"code"`
		Error string `json:"error"`
		Data  string `json:"data"`
	}
	if err = json.Unmarshal(resp, &publicKeyResp); err != nil {
		return nil, err
	}
	b, _ := pem.Decode([]byte(publicKeyResp.Data))
	if b == nil {
		return nil, errors.New("pem decode error")
	}

	certBody, err := x509.ParsePKIXPublicKey(b.Bytes)
	if err != nil {
		return nil, err
	}
	key, is := certBody.(*rsa.PublicKey)
	if !is {
		return nil, errors.New("failed to parse public key")
	}
	return key, nil
}

// 加密密码
func (c *Client) encrypt(msg string) string {
	pub, err := c.storage.GetPublicKey()
	if err != nil {
		return ""
	}
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(msg))
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(ciphertext)
}
