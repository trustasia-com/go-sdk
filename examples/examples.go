// Package examples provides ...
package examples

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Message http响应格式
type Message struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

// RespWithJSON 响应请求
func RespWithJSON(w http.ResponseWriter, httpCode int, data interface{}, err error) {
	msg := Message{}
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
