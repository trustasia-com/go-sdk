// Package finance provides ...
package finance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"github.com/trustasia-com/go-sdk/pkg"
	"github.com/trustasia-com/go-sdk/pkg/client"
	"github.com/trustasia-com/go-sdk/pkg/credentials"
)

// api list
const (
	apiPaymentsList   = "/ta-finance/payments"
	apiPaymentsCreate = "/ta-finance/payments"
	apiPaymentsRefund = "/ta-finance/payments/%s/refund"
)

// Finance instance for RP
type Finance struct {
	userAgent string
	sess      *credentials.Session
	client    *http.Client
}

// New new Finance instance
func New(sess *credentials.Session) *Finance {
	return &Finance{
		userAgent: pkg.BuildUserAgent(),
		sess:      sess,
		client:    client.NewHTTPClient(),
	}
}

// PaymentList payment list of user
func (f *Finance) PaymentList(req PaymentListReq) (*PaymentListResp, error) {
	// check input
	if req.Limit < 1 {
		return nil, errors.New("Need specify req.Limit and min 1")
	}
	if req.Offset < 0 {
		return nil, errors.New("Need specify req.Offset and min 0")
	}
	if len(req.Search) != len(req.Target) {
		return nil, errors.New("Search & Target length should be equal")
	}

	vals := url.Values{}
	vals.Set("limit", fmt.Sprint(req.Limit))
	vals.Set("offset", fmt.Sprint(req.Offset))
	for _, v := range req.Search {
		vals.Add("search", v)
	}
	for _, v := range req.Target {
		vals.Add("search", v)
	}
	scope := "finance/"
	data, err := f.httpRequest(http.MethodGet, apiPaymentsList+"?"+vals.Encode(), scope, nil)
	if err != nil {
		return nil, err
	}
	resp := &PaymentListResp{}
	err = json.Unmarshal(data, resp)
	return resp, err
}

// PaymentCreate create payment for user
func (f *Finance) PaymentCreate(req PaymentCreateReq) (*PaymentCreateResp, error) {
	// check input
	if req.UserID == "" {
		return nil, errors.New("Need specify req.UserID")
	}
	if req.Nickname == "" {
		return nil, errors.New("Need specify req.Username")
	}
	if req.OrderID == "" {
		return nil, errors.New("Need specify req.OrderID")
	}
	if req.Subject == "" {
		return nil, errors.New("Need specify req.Subject")
	}
	if req.Amount < 0 {
		return nil, errors.New("Invalid req.Amount specify")
	}
	if req.Timeout < 60 {
		return nil, errors.New("Invalid req.Timeout specify, should more than 60")
	}

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	scope := "finance/"
	data, err = f.httpRequest(http.MethodPost, apiPaymentsCreate, scope, data)
	if err != nil {
		return nil, err
	}
	resp := &PaymentCreateResp{}
	err = json.Unmarshal(data, resp)
	return resp, err
}

// PaymentRefund refund specify payment
func (f *Finance) PaymentRefund(req PaymentRefundReq) (*PaymentRefundResp, error) {
	// check input

	path := fmt.Sprintln(apiPaymentsRefund, req.PaymentID)
	scope := "finance/"
	data, err := f.httpRequest(http.MethodPut, path, scope, nil)
	if err != nil {
		return nil, err
	}
	resp := &PaymentRefundResp{}
	err = json.Unmarshal(data, resp)
	return resp, err
}

func (f *Finance) httpRequest(method, path, scope string, data []byte) ([]byte, error) {
	var (
		httpReq *http.Request
		err     error
	)
	url := f.sess.Options.Endpoint + path
	if len(data) > 0 {
		httpReq, err = http.NewRequest(method, url, bytes.NewReader(data))
	} else {
		httpReq, err = http.NewRequest(method, url, nil)
	}
	httpReq.Header.Set("User-Agent", f.userAgent)

	if err = f.sess.SignRequest(httpReq, scope); err != nil {
		return nil, err
	}
	httpResp, err := f.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	return io.ReadAll(httpResp.Body)
}
