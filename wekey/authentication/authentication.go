package authentication

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"net/http"

	"github.com/trustasia-com/go-sdk/pkg"
	"github.com/trustasia-com/go-sdk/pkg/client"
	"github.com/trustasia-com/go-sdk/wekey/authentication/types"
)

const (
	pathGetPublicKey = "/ta-auth/public-key"
)

// Client authentication client
type Client struct {
	// your application client secret
	ClientID string
	// your auth endpoint e.g: https://wekey.wekey.com
	Endpoint string

	userAgent string
	publicKey *rsa.PublicKey
	client    *http.Client
}

// NewClient new
func NewClient(clientID string, endpoint string) (*Client, error) {
	c := &Client{
		ClientID:  clientID,
		Endpoint:  endpoint,
		userAgent: pkg.BuildUserAgent(),
		client:    client.NewHTTPClient(),
	}
	public, err := c.GetPublicKey()
	if err != nil {
		return nil, err
	}
	c.publicKey = public
	return c, nil
}

// GetPublicKey 获取公钥证书
func (c *Client) GetPublicKey() (*rsa.PublicKey, error) {

	req := &ClientRequest{
		Path:   pathGetPublicKey,
		method: "GET",
	}
	res := c.httpRequest(req)
	if res.Err != nil {
		return nil, res.Err
	}
	var publicResponse types.GetPublicKeyResponse
	if err := json.Unmarshal(res.Body, &publicResponse); err != nil {
		return nil, err
	}

	b, _ := pem.Decode([]byte(publicResponse.Data))
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
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, c.publicKey, []byte(msg))
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(ciphertext)
}
