// Package finance provides ...
package finance

import (
	"time"
)

// PaymentDetail 付款详细
type PaymentDetail struct {
	PaymentID string `json:"payment_id"` // 支付ID
	MchName   string `json:"mch_name"`   // 商户名称
	Nickname  string `json:"nickname"`   // 用户昵称

	OrderID   string    `json:"order_id"`   // 商户订单ID
	Subject   string    `json:"subject"`    // 商户订单标题
	Amount    int       `json:"amount"`     // 金额，分
	Note      string    `json:"note"`       // 订单备注
	UpdatedAt time.Time `json:"updated_at"` // 付款更新
	CreatedAt time.Time `json:"created_at"` // 付款创建
	ExpiredAt time.Time `json:"expired_at"` // 付款失效

	Class   string    `json:"class,omitempty"`    // 付款分类，payment/subscribe/refunds
	Status  string    `json:"status,omitempty"`   // 付款状态
	PayChan string    `json:"pay_chan,omitempty"` // 付款渠道
	PaidAt  time.Time `json:"paid_at,omitempty"`  // 付款时间
}

// PaymentListReq req
type PaymentListReq struct {
	Search []string // 搜索项
	Target []string // 搜索内容
	Limit  int      // 条数限制
	Offset int      // 偏移数
}

// PaymentListResp resp
type PaymentListResp struct {
	Total int64           `json:"total"` // 总共条数
	List  []PaymentDetail `json:"list"`  // 列表
}

// PaymentCreateReq req
type PaymentCreateReq struct {
	UserID   string `json:"user_id"`  // 商户用户唯一ID
	Nickname string `json:"nickname"` // 商户用户下单时昵称

	// 业务ID，如果商户使用同一收银台收款，但有不同业务，可以用来区分下级业务
	BusinessCode string `json:"business_code"`
	OrderID      string `json:"order_id"`   // 商户订单ID，回调回原样传回
	Subject      string `json:"subject"`    // 商户订单标题
	Amount       int    `json:"amount"`     // 订单实际金额，分
	Note         string `json:"note"`       // 订单备注，回调原样传回
	Timeout      int    `json:"timeout"`    // 支付到期时间，最少60s
	ReturnURL    string `json:"return_url"` // 支付完成，收银台跳转地址
	ProductID    string `json:"product_id"` // 订单产品ID
}

// PaymentCreateResp resp
type PaymentCreateResp struct {
	ReturnURL string `json:"return_url"` // 商户创建支付后引导用户跳转支付地址
}

// PaymentRefundReq req
type PaymentRefundReq struct {
	PaymentID string // 支付ID
}

// PaymentRefundResp resp
type PaymentRefundResp struct {
	PaymentID string `json:"payment_id"` // 支付ID
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
	MchID   string    `json:"mch_id"`  // 商户ID
	Do      PaymentDo `json:"do"`      // 发货 deliver， 退货 refund
	Nonce   string    `json:"nonce"`   // 随机数
	Content string    `json:"content"` // stdencoding base64

	Sign string `json:"sign"` // 签名内容
}

// CallbackContent 回调内容
type CallbackContent struct {
	BusinessCode string    `json:"business_code"` // 业务ID
	OrderID      string    `json:"order_id"`      // 订单ID
	UserID       string    `json:"user_id"`       // 用户ID
	Amount       int       `json:"amount"`        // 支付金额
	Class        string    `json:"class"`         // 商品分类
	Note         string    `json:"note"`          // 备注
	ProductID    string    `json:"product_id"`    // 产品ID
	EffctedAt    time.Time `json:"effcted_at"`    // 产生时间：付款时间内/退款时间
	PayChan      string    `json:"pay_chan"`      // 支付渠道
}

// SubscribeCreateReq 创建订阅单
type SubscribeCreateReq struct {
	UserID   string `json:"user_id"`  // 订阅用户ID
	Nickname string `json:"nickname"` // 订阅用户昵称

	ReturnURL string `json:"return_url"` // 订阅完成返回地址
	ProductID string `json:"product_id"` // 订阅产品ID
}

// SubscribeCreateResp 响应
type SubscribeCreateResp struct {
	Status string `json:"status"` // 订阅状态
	Token  string `json:"token"`  // 鉴权token

	Redirect string `json:"redirect,omitempty"` // 重定向地址
}
