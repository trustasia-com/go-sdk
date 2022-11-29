// Package main provides ...
package main

import (
	"errors"
	"net/http"

	"github.com/trustasia-com/go-sdk/pkg/credentials"
	"github.com/trustasia-com/go-sdk/wekey"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type user struct {
	UserID      string `json:"user_id"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
}

var u user

func init() {
	u = user{
		UserID:      "uid",
		Username:    "username",
		DisplayName: "displayName",
	}
}

func main() {
	// create credential
	opts := credentials.Options{
		AccessKey:  "cd7q8pc3c37lvrfjsge0",
		SecretKey:  "cd7q8pc3c37lvrfjsgeg",
		Endpoint:   "http://localhost:9000",
		SignerType: credentials.SignatureDefault,
	}
	sess, err := credentials.New(opts, false)
	if err != nil {
		panic(err)
	}
	// create client
	client := wekey.New(sess)

	e := gin.Default()
	e.LoadHTMLGlob("./*.html")
	store := cookie.NewStore([]byte("secret"))
	e.Use(sessions.Sessions("mysession", store))

	e.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	e.GET("/dashboard", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.html", nil)
	})
	e.GET("/login", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("user_id", u.UserID)
		session.Save()

		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 0,
		})
	})
	e.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Delete("user_id")
		session.Save()

		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 0,
		})
	})
	// auth
	e.POST("/login/qrcode", func(c *gin.Context) {
		var req struct {
			Username string `json:"username" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(http.StatusBadRequest, "invalid username")
			return
		}
		if req.Username != u.Username {
			c.String(http.StatusBadRequest, "user not exists")
			return
		}

		method := wekey.AuthMethodQRCode
		if c.Query("method") == "push" {
			method = wekey.AuthMethodPush
		}
		sdkreq := wekey.AuthRequestReq{
			Method:   method,
			UserID:   u.UserID,
			Username: u.Username,
		}
		resp, err := client.AuthRequest(sdkreq)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 0,
			"data": resp,
		})
	})
	e.GET("/login/result", func(c *gin.Context) {
		req := wekey.AuthResultReq{
			MsgID: c.Query("msg_id"),
		}
		resp, err := client.AuthResult(req, func(userID string) error {
			session := sessions.Default(c)
			session.Set("user_id", userID)
			session.Save()
			return nil
		})
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 0,
			"data": resp,
		})
	})
	// 登录鉴权
	e.Use(func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("user_id")
		if uid == nil {
			c.Abort()
			c.String(http.StatusUnauthorized, "unauthorized")
			return
		}
		c.Keys["uid"] = uid
		c.Next()
	})
	e.GET("/userinfo", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 0,
			"data": u,
		})
	})
	e.POST("/register/qrcode", func(c *gin.Context) {
		uid := c.Keys["uid"].(string)

		req := wekey.RegQRCodeReq{
			DisplayName: u.DisplayName,
			UserID:      uid,
			Username:    u.Username,
		}
		resp, err := client.RegQRCode(req)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 0,
			"data": resp,
		})
	})
	e.GET("/register/result", func(c *gin.Context) {
		uid := c.Keys["uid"].(string)

		req := wekey.RegResultReq{
			MsgID: c.Query("msg_id"),
		}
		resp, err := client.RegResult(req, func(userID string) error {
			if userID != uid {
				return errors.New("user id not equal")
			}
			session := sessions.Default(c)
			session.Set("user_id", uid)
			session.Save()
			return nil
		})
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 0,
			"data": resp,
		})
	})

	e.Run(":3002")
}
