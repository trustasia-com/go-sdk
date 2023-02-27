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
		AccessKey:  "53deb68751cbbbea77d08a2dab740457",
		SecretKey:  "eMJhAQMnsUxtsRMOFwFpA4Ncmyho9mQDaseb5uO2",
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
	e.GET("/create-app", handleCreateOrUpdateApp)
	e.GET("/delete-app", handleDeleteApp)
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

	e.Run(":3002")
}

func handleCreateOrUpdateApp(c *gin.Context) {
	req := app.CreateOrUpdateAppReq{
		Slug:    "test-app",
		Name:    "Test App",
		ExLogin: false,

		RpInfo: app.RpInfo{
			RpID:    "localhost:3002",
			Origins: []string{"https://localhost:3002"},
		},
	}
	err := client.CreateOrUpdateApp(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "")
}

func handleDeleteApp(c *gin.Context) {
	req := app.DeleteAppReq{
		Slug: "test-app",
	}
	err := client.DeleteApp(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "")
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

	typ := app.TypeFidoScan
	if c.Query("type") == "cosign" {
		typ = app.TypeCosignScan
	}
	sdkreq := app.AuthRequestReq{
		Slug: "wekey-dev",
		Type: typ,

		RpUserID:   u.UserID,
		RpUsername: u.Username,
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
}

func handleLoginResult(c *gin.Context) {
	req := app.AuthResultReq{
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
}

func handleUserinfo(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"data": u,
	})
}

func handleRegisterQrCode(c *gin.Context) {
	uid := c.Keys["uid"].(string)

	req := app.RegQRCodeReq{
		Slug:           "wekey-dev",
		CredentialName: "test-cred",

		RpUserID: uid,
		// RpUsername: "",
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
}

func handleRegisterResult(c *gin.Context) {
	req := app.RegResultReq{
		MsgID: c.Query("msg_id"),
	}
	resp, err := client.RegResult(req, func(userID string) error {
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
