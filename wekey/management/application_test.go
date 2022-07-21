package management

import (
	"fmt"
	"log"
	"testing"

	"github.com/trustasia-com/go-sdk/wekey/management/types"
)

func TestClient_QueryApp(t *testing.T) {
	initTestClient()
	data, err := testClient.QueryApp(types.QueryAppRequest{Name: ""})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

func TestClient_AppInfo(t *testing.T) {
	initTestClient()
	data, err := testClient.AppInfo("VrobkZpx1L3Mj6aW87gK")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)

}

func TestClient_CreateApp(t *testing.T) {
	initTestClient()
	data, err := testClient.CreateApp(types.CreateAppRequest{Name: "test123"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

func TestClient_DelApp(t *testing.T) {
	initTestClient()
	err := testClient.DelApp("736p4N0A15GdnoQj5xmB")
	if err != nil {
		log.Fatal(err)
	}
}

func TestClient_GetAppIDPs(t *testing.T) {
	initTestClient()
	data, err := testClient.GetAppIDPs("VrobkZpx1L3Mj6aW87gK")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

func TestClient_UpdateAppBasic(t *testing.T) {
	initTestClient()
	data, err := testClient.CreateApp(types.CreateAppRequest{Name: "test123"})
	if err != nil {
		log.Fatal(err)
	}
	err = testClient.UpdateAppBasic(data, types.UpdateAppBasicRequest{
		Name:        "test3121",
		Logo:        "https://cdn.wekey.com/website/images/download-sec2-icon2.png",
		Description: "update test",
	})
	if err != nil {
		log.Fatal(err)
	}
	app, err := testClient.AppInfo(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(app)
}

func TestClient_UpdateAppGrant(t *testing.T) {
	initTestClient()
	data, err := testClient.CreateApp(types.CreateAppRequest{Name: "test123"})
	if err != nil {
		log.Fatal(err)
	}
	err = testClient.UpdateAppGrant(data, types.UpdateAppGrantRequest{
		TokenEndpointAuthMethod: "Post",
		LoginURL:                fmt.Sprintf("https://me-dev.wekey.cn/ta-wekey-dash/demo/auth-url/%s", data),
		LoginCallbackURLs:       []string{},
		LogoutCallbackURLs:      []string{},
		IDTokenExpire:           1209600,
		AccessTokenExpire:       1209600,
		RefreshTokenExpire:      1209600,
	})
	if err != nil {
		log.Fatal(err)
	}
	app, err := testClient.AppInfo(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(app)
}

func TestClient_UpdateAppAdvance(t *testing.T) {
	initTestClient()
	data, err := testClient.CreateApp(types.CreateAppRequest{Name: "test123"})
	if err != nil {
		log.Fatal(err)
	}
	err = testClient.UpdateAppAdvance(data, types.UpdateAppAdvanceRequest{
		GrantTypes:     []string{"authorization_code", "implicit", "refresh_token", "client_credentials", "password"},
		SignAlgo:       "HS256",
		OIDCConformant: false,
	})
	if err != nil {
		log.Fatal(err)
	}
	app, err := testClient.AppInfo(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(app)
}

func TestClient_UpdateAppSecret(t *testing.T) {
	initTestClient()
	data, err := testClient.CreateApp(types.CreateAppRequest{Name: "test123"})
	if err != nil {
		log.Fatal(err)
	}
	err = testClient.UpdateAppSecret(data)
	if err != nil {
		log.Fatal(err)
	}
	app, err := testClient.AppInfo(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(app)
}

func TestClient_QueryResource(t *testing.T) {
	initTestClient()
	data, err := testClient.QueryResource(types.QueryResourceRequest{Name: ""})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

func TestClient_ResourceInfo(t *testing.T) {
	initTestClient()
	data, err := testClient.ResourceInfo("P2OyZVrY1EvqWpXwKv4R")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

func TestClient_CreateResource(t *testing.T) {
	initTestClient()
	data, err := testClient.CreateResource(types.CreateResourceRequest{Name: "test123", Audience: "https://baidu.com"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

func TestClient_DelResource(t *testing.T) {
	initTestClient()
	err := testClient.DelResource("QwnlJYpO1kO15RD0em4G")
	fmt.Println(err)
}

func TestClient_UpdateResource(t *testing.T) {
	initTestClient()
	data, err := testClient.CreateResource(types.CreateResourceRequest{Name: "test123", Audience: "https://baidu.com"})
	if err != nil {
		log.Fatal(err)
	}
	err = testClient.UpdateResource(data, types.UpdateResourceRequest{
		Name:                "test",
		RBACEnable:          false,
		TokenWithPermission: false,
		AllowSkipConsent:    false,
		AllowOfflineAccess:  false,
	})
	if err != nil {
		log.Fatal(err)
	}
	r, err := testClient.ResourceInfo(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}

func TestClient_QueryPermission(t *testing.T) {
	initTestClient()
	data, err := testClient.QueryPermission(types.QueryPermissionRequest{
		ResourceID: "p3WNKkEx1N0dD4wV26PY",
		Scope:      "",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

func TestClient_CreatePermission(t *testing.T) {
	initTestClient()
	data, err := testClient.CreatePermission(types.CreatePermissionRequest{
		ResourceID:  "p3WNKkEx1N0dD4wV26PY",
		Scope:       "read:order",
		Description: "test",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

func TestClient_DelPermission(t *testing.T) {
	initTestClient()

	err := testClient.DelPermission("P2OyZVrY1EvqWpXwKv4R")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(err)
}

func TestClient_QueryResourceApps(t *testing.T) {
	initTestClient()
	data, err := testClient.QueryResourceApps("P2OyZVrY1EvqWpXwKv4R")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

func TestClient_QueryAppResources(t *testing.T) {
	initTestClient()
	data, err := testClient.QueryAppResources("Qv4n6ZwKMVjMOBWypxGz")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

func TestClient_OpenAppResource(t *testing.T) {
	initTestClient()
	err := testClient.OpenAppResource("Qv4n6ZwKMVjMOBWypxGz", "p3WNKkEx1N0dD4wV26PY")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(err)
}

func TestClient_CloseAppResource(t *testing.T) {
	initTestClient()
	err := testClient.CloseAppResource("Qv4n6ZwKMVjMOBWypxGz", "p3WNKkEx1N0dD4wV26PY")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(err)
}

func TestClient_OpenAppIDP(t *testing.T) {
	initTestClient()
	err := testClient.OpenAppIDP("Qv4n6ZwKMVjMOBWypxGz", "github", "github")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(err)
}

func TestClient_CloseAppIDP(t *testing.T) {
	initTestClient()
	err := testClient.CloseAppIDP("Qv4n6ZwKMVjMOBWypxGz", "github", "github")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(err)
}
