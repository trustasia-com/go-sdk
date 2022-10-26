// Package main provides ...
package main

import (
	"fmt"

	"github.com/trustasia-com/go-sdk/pkg/credentials"
	"github.com/trustasia-com/go-sdk/wekey"
)

func main() {
	// create credential
	opts := credentials.Options{
		AccessKey:  "cd7q8pc3c37lvrfjsge0",
		SecretKey:  "cd7q8pc3c37lvrfjsgeg",
		Endpoint:   "https://api-dev.wekey.cn",
		SignerType: credentials.SignatureDefault,
	}
	sess, err := credentials.New(opts, false)
	if err != nil {
		panic(err)
	}
	// create client
	client := wekey.New(sess)

	resp, err := client.RegQRCode(wekey.RegQRCodeReq{
		UserID:      "test_user_id",
		Username:    "test_username",
		DisplayName: "test_displayname",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
