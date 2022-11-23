// Package finance provides ...
package finance

import (
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

	apiSubscribeCreate = "/ta-finance/subscribe"
)

// Finance instance for RP
type Finance struct {
	userAgent string
	client    *client.HTTPClient
}

// New new Finance instance
func New(sess *credentials.Session) *Finance {
	return &Finance{
		userAgent: pkg.BuildUserAgent(),
		client:    client.NewHTTPClient(sess),
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
	msg, err := f.client.Request(http.MethodGet, apiPaymentsList+"?"+vals.Encode(), scope, nil)
	if err != nil {
		return nil, err
	}
	resp := &PaymentListResp{}
	err = json.Unmarshal(msg.Data, resp)
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
	if req.ProductID == "" {
		return nil, errors.New("Need specify req.ProductID")
	}

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	scope := "finance/"
	msg, err := f.client.Request(http.MethodPost, apiPaymentsCreate, scope, data)
	if err != nil {
		return nil, err
	}
	resp := &PaymentCreateResp{}
	err = json.Unmarshal(msg.Data, resp)
	return resp, err
}

// PaymentRefund refund specify payment
func (f *Finance) PaymentRefund(req PaymentRefundReq) (*PaymentRefundResp, error) {
	// check input
	if req.PaymentID == "" {
		return nil, errors.New("Invalid req.Payment specify")
	}

	path := fmt.Sprintln(apiPaymentsRefund, req.PaymentID)
	scope := "finance/"
	msg, err := f.client.Request(http.MethodPut, path, scope, nil)
	if err != nil {
		return nil, err
	}
	resp := &PaymentRefundResp{}
	err = json.Unmarshal(msg.Data, resp)
	return resp, err
}

// FinanceCallback callback
func (f *Finance) FinanceCallback(r *http.Request) (*CallbackFinanceReq, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	req := &CallbackFinanceReq{}
	err = json.Unmarshal(data, req)
	if err != nil {
		return nil, err
	}
	// validate signature
	vals := url.Values{}
	vals.Set("mch_id", req.MchID)
	vals.Set("do", string(req.Do))
	vals.Set("nonce", req.Nonce)
	vals.Set("content", string(req.Content))
	if req.Sign != f.client.Session.SumHMAC([]byte(vals.Encode())) {
		return nil, errors.New("failed to validate signature")
	}
	return req, nil
}

// SubscribeCreate create subscribe
func (f *Finance) SubscribeCreate(req SubscribeCreateReq) (*SubscribeCreateResp, error) {
	// check input
	if req.UserID == "" {
		return nil, errors.New("Need specify req.UserID")
	}
	if req.Nickname == "" {
		return nil, errors.New("Need specify req.Username")
	}
	if req.ProductID == "" {
		return nil, errors.New("Need specify req.ProductID")
	}

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	scope := "finance/"
	msg, err := f.client.Request(http.MethodPost, apiSubscribeCreate, scope, data)
	if err != nil {
		return nil, err
	}
	resp := &SubscribeCreateResp{}
	err = json.Unmarshal(msg.Data, resp)
	return resp, err
}
