// Package main provides ...
package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/trustasia-com/go-sdk/pkg/credentials"
	"github.com/trustasia-com/go-sdk/webauthn"
)

type user struct {
	ID       int
	Username string
	Nickname string
}

type message struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

func main() {
	// create credential
	opts := credentials.Options{
		AccessKey:  "c676aeb85dc2c553",
		SecretKey:  "a290bc6da6318e5aaf18da9a29bc6ec2",
		Endpoint:   "https://api-dev.wekey.com",
		SignerType: credentials.SignatureDefault,
	}
	sess, err := credentials.New(opts, true)
	if err != nil {
		panic(err)
	}
	// create client
	client := webauthn.New(sess)

	// 这里以 gin 为例子
	// index.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// 我们建议所有的接口都用post
	// register
	http.HandleFunc("/attestation/options", func(w http.ResponseWriter, r *http.Request) {
		// 构建注册请求，这里有可能用户在登录状态，你可以通过session/cookie
		// 方式查询到用户
		// user := ...

		userID := hex.EncodeToString([]byte("test_user_id"))
		req := webauthn.StartSignUpReq{
			Username:    "test_username",
			DisplayName: "test_nickname",
		}
		data, err := client.StartSignUp(req, userID)
		if err != nil {
			respWithJSON(w, 400, nil, err)
			return
		}
		// 响应给前端
		respWithJSON(w, 200, data, nil)
	})
	http.HandleFunc("/attestation/result", func(w http.ResponseWriter, r *http.Request) {
		data, err := client.FinishSignUp(r)
		if err != nil {
			respWithJSON(w, 400, nil, err)
			return
		}
		respWithJSON(w, 200, data, nil)
	})
	// login
	http.HandleFunc("/assertion/options", func(w http.ResponseWriter, r *http.Request) {
		// 通常通过提交的信息查询用户，比如 username
		//
		// user := ...

		userID := hex.EncodeToString([]byte("test_user_id"))
		req := webauthn.StartSignInReq{
			Username: "test_username",
		}
		data, err := client.StartSignIn(req, userID)
		if err != nil {
			respWithJSON(w, 400, nil, err)
			return
		}
		respWithJSON(w, 200, data, nil)
	})
	http.HandleFunc("/assertion/result", func(w http.ResponseWriter, r *http.Request) {
		data, err := client.FinishSignIn(r)
		if err != nil {
			respWithJSON(w, 400, nil, err)
			return
		}
		respWithJSON(w, 200, data, nil)
	})
	port := "9000"
	fmt.Println("listen and serve: " + port)
	http.ListenAndServe(":"+port, nil)
}

func respWithJSON(w http.ResponseWriter, httpCode int, data interface{}, err error) {
	msg := message{}
	if httpCode != 200 {
		msg.Code = 1
	}
	msg.Data = data
	if err != nil {
		fmt.Println(err)
		msg.Error = err.Error()
	}
	raw, _ := json.Marshal(msg)
	w.Header().Set("Content-Type", "application/json")
	w.Write(raw)
}
