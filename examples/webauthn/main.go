// Package main provides ...
package main

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/trustasia-com/go-sdk/examples"
	"github.com/trustasia-com/go-sdk/pkg/credentials"
	"github.com/trustasia-com/go-sdk/webauthn"
)

type user struct {
	ID       int
	Username string
	Nickname string
}

func main() {
	// create credential
	opts := credentials.Options{
		AccessKey:  "c676aeb85dc2c553",
		SecretKey:  "a290bc6da6318e5aaf18da9a29bc6ec2",
		Endpoint:   "https://api-dev.wekey.cn",
		SignerType: credentials.SignatureDefault,
	}
	sess, err := credentials.New(opts, true)
	if err != nil {
		panic(err)
	}
	// create client
	client := webauthn.New(sess)

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

		userID := base64.RawURLEncoding.EncodeToString([]byte("test_user_id"))
		req := webauthn.StartSignUpReq{
			Username:    "test_username",
			DisplayName: "test_nickname",
		}
		data, err := client.StartSignUp(req, userID)
		if err != nil {
			examples.RespWithJSON(w, 400, nil, err)
			return
		}
		// 响应给前端
		examples.RespWithJSON(w, 200, data, nil)
	})
	http.HandleFunc("/attestation/result", func(w http.ResponseWriter, r *http.Request) {
		data, err := client.FinishSignUp(r)
		if err != nil {
			examples.RespWithJSON(w, 400, nil, err)
			return
		}
		examples.RespWithJSON(w, 200, data, nil)
	})
	// login
	http.HandleFunc("/assertion/options", func(w http.ResponseWriter, r *http.Request) {
		// 通常通过提交的信息查询用户，比如 username
		//
		// user := ...

		userID := base64.RawURLEncoding.EncodeToString([]byte("test_user_id"))
		req := webauthn.StartSignInReq{
			Username: "test_username",
		}
		data, err := client.StartSignIn(req, userID)
		if err != nil {
			examples.RespWithJSON(w, 400, nil, err)
			return
		}
		examples.RespWithJSON(w, 200, data, nil)
	})
	http.HandleFunc("/assertion/result", func(w http.ResponseWriter, r *http.Request) {
		data, err := client.FinishSignIn(r)
		if err != nil {
			examples.RespWithJSON(w, 400, nil, err)
			return
		}
		// 如果这里err == nil，代表注册成功
		// 设置 cookie
		examples.RespWithJSON(w, 200, data, nil)
	})
	port := "9000"
	fmt.Println("listen and serve: " + port)
	http.ListenAndServe(":"+port, nil)
}
