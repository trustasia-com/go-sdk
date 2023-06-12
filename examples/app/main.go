// Package main provides ...
package main

import (
	"net/http"

	"github.com/trustasia-com/go-sdk/app"
	"github.com/trustasia-com/go-sdk/pkg/credentials"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type user struct {
	UserID      string `json:"user_id"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
}

var (
	u      user
	client *app.App
)

func init() {
	u = user{
		UserID:      "uid",
		Username:    "username",
		DisplayName: "displayName",
	}
	// create credential
	opts := credentials.Options{
		AccessKey:  "ce6pdnjtlntq3g6d22v0",
		SecretKey:  "3bb5c5ab4e4801ca0cf7c77f5a996020958a71c8",
		Endpoint:   "http://localhost:9000",
		SignerType: credentials.SignatureDefault,
	}
	sess, err := credentials.New(opts, false)
	if err != nil {
		panic(err)
	}
	// create client
	client = app.New(sess)
}

func main() {

	e := gin.Default()
	e.LoadHTMLGlob("./*.html")
	store := cookie.NewStore([]byte("secret"))
	e.Use(sessions.Sessions("mysession", store))
	// static page
	e.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	e.GET("/dashboard", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.html", nil)
	})
	e.GET("/login", handleLogin)
	// auth
	e.POST("/login/qrcode", handleLoginQrCode)
	e.GET("/login/result", handleLoginResult)
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
	e.GET("/logout", handleLogout)
	e.GET("/userinfo", handleUserinfo)
	e.POST("/register/qrcode", handleRegisterQrCode)
	e.GET("/register/result", handleRegisterResult)
	e.GET("/cosign", handleCosignList)

	e.Run(":3002")
}

// 模拟直接登录
func handleLogin(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("user_id", u.UserID)
	session.Save()

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
	})
}

// 登出
func handleLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user_id")
	session.Save()

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
	})
}

func handleLoginQrCode(c *gin.Context) {
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

	sdkreq := app.LoginQRCodeReq{
		ClientType: app.ClientTypeWeb,
	}
	resp, err := client.LoginQRCode(sdkreq)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"data": resp,
	})
}

func handleLoginResult(c *gin.Context) {
	req := app.LoginResultReq{
		MsgID: c.Query("msg_id"),
	}
	resp, err := client.LoginResult(req, func(userID string) error {
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
}

func handleUserinfo(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"data": u,
	})
}

func handleRegisterQrCode(c *gin.Context) {
	uid := c.Keys["uid"].(string)

	req := app.BindQRCodeReq{
		ClientType:        app.ClientTypeWeb,
		RpUserID:          uid,
		RpUserDisplayName: "hello",
	}
	resp, err := client.BindQRCode(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"data": resp,
	})
}

func handleRegisterResult(c *gin.Context) {
	req := app.BindResultReq{
		MsgID: c.Query("msg_id"),
	}
	resp, err := client.BindResult(req, func(userID string) error {
		// TODO something
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
}

func handleCosignList(c *gin.Context) {
	req := app.CredentialsReq{
		RpUserID: "akfjasl",
	}
	resp, err := client.CosignCredentials(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"data": resp,
	})
}
