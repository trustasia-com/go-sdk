// Package client provides ...
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/trustasia-com/go-sdk/pkg"
	"github.com/trustasia-com/go-sdk/pkg/credentials"
	"github.com/trustasia-com/go-sdk/pkg/message"
)

// HTTPClient http client
type HTTPClient struct {
	session *credentials.Session
	*http.Client
	useragent string
}

// NewHTTPClient new http client
func NewHTTPClient(sess *credentials.Session) *HTTPClient {
	cli := &HTTPClient{
		session:   sess,
		useragent: pkg.BuildUserAgent(),
	}
	cli.Client = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: func(dialer *net.Dialer) func(context.Context, string, string) (net.Conn, error) {
				return dialer.DialContext
			}(&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}),
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
		Timeout: time.Second * 30,
	}
	return cli
}

// Request http request
func (cli *HTTPClient) Request(method, path, scope string, data []byte) (*message.JSONRawMessage, error) {
	var (
		httpReq *http.Request
		err     error
	)
	url := cli.session.Options.Endpoint + path
	if len(data) > 0 {
		httpReq, err = http.NewRequest(method, url, bytes.NewReader(data))
	} else {
		httpReq, err = http.NewRequest(method, url, nil)
	}
	httpReq.Header.Set("User-Agent", cli.useragent)
	if err = cli.session.SignRequest(httpReq, scope); err != nil {
		return nil, err
	}
	httpResp, err := cli.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	data, err = io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	msg := &message.JSONRawMessage{}
	err = json.Unmarshal(data, msg)
	if err != nil {
		return nil, err
	}
	if msg.Code != 0 {
		return nil, fmt.Errorf("(%d)%s", msg.Code, msg.Error)
	}
	return msg, nil
}
