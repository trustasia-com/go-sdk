// Package client provides ...
package client

import (
	"context"
	"net"
	"net/http"
	"time"
)

// NewHTTPClient new http client
func NewHTTPClient() *http.Client {
	return &http.Client{
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
}
