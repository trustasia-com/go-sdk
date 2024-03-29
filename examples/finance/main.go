// Package main provides ...
package main

import (
	"embed"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/trustasia-com/go-sdk/examples"
	"github.com/trustasia-com/go-sdk/finance"
	"github.com/trustasia-com/go-sdk/pkg/credentials"
)

type order struct {
	ID     string `json:"id"`
	Amount int    `json:"amount"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type user struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
}

var userOrder struct {
	// 登录用户
	User *user `json:"user"`
	// 初始化订单
	Order *order `json:"order"`
}

// 初始化
func init() {
	u := &user{
		ID:       "test_user_id",
		Nickname: "test_nickname",
	}
	o := &order{
		ID:     "test_order_id",
		Amount: 1,
		Title:  "test_order_title",
		Status: "unpaid",
	}
	userOrder.User = u
	userOrder.Order = o
}

//go:embed *.html
var f embed.FS

func main() {
	// create credential
	opts := credentials.Options{
		AccessKey: "test_mch",
		SecretKey: "bff149a0b87f5b0e00d9dd364e9ddaa0",
		Endpoint:  "https://pay-dev.wekey.cn",
		// Endpoint:   "http://localhost:9000",
		SignerType: credentials.SignatureDefault,
	}
	sess, err := credentials.New(opts, true)
	if err != nil {
		panic(err)
	}

	cli := finance.New(sess)

	t, err := template.ParseFS(f, "*.html")
	if err != nil {
		panic(err)
	}
	// index页面
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		if r.FormValue("recreate") == "true" {
			ts := time.Now().Unix()
			userOrder.User.ID = fmt.Sprintf("test_user_id%d", ts)
			userOrder.Order.ID = fmt.Sprintf("test_order_title%d", ts)
		}
		err := t.Execute(w, userOrder)
		if err != nil {
			panic(err)
		}
	})
	// 创建支付订单
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		if userOrder.Order.Status == "complete" {
			examples.RespWithJSON(w, 400, nil, errors.New("订单已完成，无需重复支付"))
			return
		}
		var params struct {
			OrderID string `json:"order_id"`
		}
		data, err := io.ReadAll(r.Body)
		if err != nil {
			examples.RespWithJSON(w, 400, nil, err)
			return
		}
		err = json.Unmarshal(data, &params)
		if err != nil {
			examples.RespWithJSON(w, 400, nil, err)
			return
		}
		if params.OrderID != userOrder.Order.ID {
			examples.RespWithJSON(w, 400, nil, errors.New("订单不存在"))
			return
		}

		req := finance.PaymentCreateReq{
			UserID:   userOrder.User.ID,
			Nickname: userOrder.User.Nickname,

			BusinessCode: "temp",
			OrderID:      userOrder.Order.ID,
			Subject:      userOrder.Order.Title,
			Amount:       userOrder.Order.Amount,
			Note:         "note2",
			Timeout:      3600,
			ReturnURL:    "https://temp.wekey.cn?order_id=" + userOrder.Order.ID,
			ProductID:    "test_product_code",
		}

		resp, err := cli.PaymentCreate(req)
		if err != nil {
			examples.RespWithJSON(w, 400, nil, err)
			return
		}
		examples.RespWithJSON(w, 200, resp.ReturnURL, nil)
	})

	http.HandleFunc("/subscribe", func(w http.ResponseWriter, r *http.Request) {
		req := finance.SubscribeCreateReq{
			UserID:   userOrder.User.ID,
			Nickname: userOrder.User.Nickname,

			ReturnURL: "https://temp.wekey.cn?order_id=" + userOrder.Order.ID,
			ProductID: "test_product_code_subscribe",
		}
		resp, err := cli.SubscribeCreate(req)
		if err != nil {
			examples.RespWithJSON(w, 400, nil, err)
			return
		}
		examples.RespWithJSON(w, 200, resp, nil)
	})
	http.HandleFunc("/plan", func(w http.ResponseWriter, r *http.Request) {
		examples.RespWithJSON(w, 200, userOrder.Order, nil)
	})
	// 回调
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		req, err := cli.FinanceCallback(r)
		if err != nil {
			examples.RespWithJSON(w, 400, nil, err)
			return
		}
		// 区分回调类型
		switch req.Do {
		case finance.PaymentDoDeliver:
			var content struct {
				OrderID string `json:"order_id"`
			}
			data, err := base64.StdEncoding.DecodeString(req.Content)
			if err != nil {
				examples.RespWithJSON(w, 400, nil, errors.New("非base64数据"))
				return
			}
			err = json.Unmarshal(data, &content)
			if err != nil {
				examples.RespWithJSON(w, 400, nil, errors.New("发货内容不对"))
				return
			}
			if content.OrderID != userOrder.Order.ID {
				fmt.Println("orderid not equal: ", req)

				examples.RespWithJSON(w, 400, nil, errors.New("invalid order id"))
				return
			}
			userOrder.Order.Status = "paid"
		case finance.PaymentDoRefund:
			// TODO
		case finance.PaymentDoRenew:
			fmt.Println("subscribe start")
			userOrder.Order.Status = "subscribed"
		default:

		}
	})
	http.ListenAndServe(":12345", nil)
}
