// Package finance provides ...
package finance

import (
	"time"
)

// PaymentDetail 付款详细
type PaymentDetail struct {
	PaymentID string `json:"payment_id"`
	MchName   string `json:"mch_name"`
	Nickname  string `json:"nickname"`

	OrderID   string    `json:"order_id"`
	Subject   string    `json:"subject"`
	Amount    int       `json:"amount"`
	Note      string    `json:"note"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`

	Class   string    `json:"class,omitempty"`
	Status  string    `json:"status,omitempty"`
	PayChan string    `json:"pay_chan,omitempty"`
	PaidAt  time.Time `json:"paid_at,omitempty"`
}

// PaymentListReq req
type PaymentListReq struct {
	Search []string
	Target []string
	Limit  int
	Offset int
}

// PaymentListResp resp
type PaymentListResp struct {
	Total int64           `json:"total"`
	List  []PaymentDetail `json:"list"`
}

// PaymentCreateReq req
type PaymentCreateReq struct {
	UserID   string `json:"user_id"`
	Nickname string `json:"nickname"`

	BusinessCode string `json:"business_code"`
	OrderID      string `json:"order_id"`
	Subject      string `json:"subject"`
	Amount       int    `json:"amount"`
	Note         string `json:"note"`
	Timeout      int    `json:"timeout"`
	ReturnURL    string `json:"return_url"`
	ProductID    string `json:"product_id"`
}

// PaymentCreateResp resp
type PaymentCreateResp struct {
	ReturnURL string `json:"return_url"`
}

// PaymentRefundReq req
type PaymentRefundReq struct {
	PaymentID string
}

// PaymentRefundResp resp
type PaymentRefundResp struct {
	PaymentID string `json:"payment_id"`
}

// PaymentDo 操作类型
type PaymentDo string

// payment action
var (
	PaymentDoDeliver PaymentDo = "deliver" // 发货
	PaymentDoRefund  PaymentDo = "refund"  // 退款
	PaymentDoRenew   PaymentDo = "renew"   // 订阅更新
)

// CallbackFinanceReq 回调请求
type CallbackFinanceReq struct {
	MchID   string    `json:"mch_id"`
	Do      PaymentDo `json:"do"` // 发货 deliver， 退货 refund
	Nonce   string    `json:"nonce"`
	Content []byte    `json:"content"`

	Sign string `json:"sign"`
}

// SubscribeCreateReq 创建订阅单
type SubscribeCreateReq struct {
	UserID   string `json:"user_id"`
	Nickname string `json:"nickname"`

	ReturnURL string `json:"return_url"`
	ProductID string `json:"product_id"`
}

// SubscribeCreateResp 响应
type SubscribeCreateResp struct {
	Status string `json:"status"`
	Token  string `json:"token"`

	Redirect string `json:"redirect,omitempty"`
}
