package management

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/trustasia-com/go-sdk/wekey/management/types"
)

func TestClient_QueryUsers(t *testing.T) {
	initTestClient()

	data, err := testClient.QueryUsers(types.QueryUserReq{
		BaseQueryRequest: types.BaseQueryRequest{Page: 1, Size: 10},
		Search:           "",
	})

	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

func TestClient_UserInfo(t *testing.T) {
	initTestClient()

	data, err := testClient.UserInfo("LQKkOjXPMpVqBzEN5og2")
	if err != nil {
		log.Fatal(err)
	}
	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

func TestClient_CreateUser(t *testing.T) {
	initTestClient()

	id, err := testClient.CreateUser(types.CreateUserReq{
		Username:     "zhangsan",
		PhoneNumber:  "",
		Email:        "",
		Password:     "123456",
		ConnectionID: "",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
}

func TestClient_DelUser(t *testing.T) {
	initTestClient()

	err := testClient.DelUser("p3WNKkEx1N0dD4wV26PY")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("del user success")
}

func TestClient_UpdateUserBasic(t *testing.T) {
	initTestClient()

	err := testClient.UpdateUserBasic("LQKkOjXPMpVqBzEN5og2", types.UpdateUserBasicRequest{
		Email:       "",
		PhoneNumber: "",
		Nickname:    "hello word",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("update user success")
}

func TestClient_UpdateUserPicture(t *testing.T) {
	initTestClient()

	err := testClient.UpdateUserPicture("LQKkOjXPMpVqBzEN5og2", types.UpdateUserPictureRequest{
		Picture: "https://static.wekey.com/web/user/avatar/l8mCfmJEnR87gIy2x8oNMdF2x.png",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("update user picture success")
}

func TestClient_UpdateUserPassword(t *testing.T) {
	initTestClient()

	err := testClient.UpdateUserPassword("LQKkOjXPMpVqBzEN5og2", types.UpdateUserPwdRequest{
		Password: "aaaaaa",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("update user password success")
}

func TestClient_UpdateUserMetadata(t *testing.T) {
	initTestClient()

	err := testClient.UpdateUserMetadata("LQKkOjXPMpVqBzEN5og2", types.UpdateUserMetadataRequest{
		UserMetadata: struct {
			Name string `json:"name"`
		}{
			Name: "张三",
		},
		AppMetadata: struct {
			Name string `json:"name"`
		}{
			Name: "李四",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("update user metadata success")
}

func TestClient_BlockUser(t *testing.T) {
	initTestClient()

	err := testClient.BlockUser("LQKkOjXPMpVqBzEN5og2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("block user success")
}

func TestClient_UnBlockUser(t *testing.T) {
	initTestClient()

	err := testClient.UnBlockUser("LQKkOjXPMpVqBzEN5og2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("unblock user success")
}

func TestClient_GetUserPermissions(t *testing.T) {
	initTestClient()

	data, err := testClient.GetUserPermissions("XKlwGg0zdvRdPyYEx4p9")
	if err != nil {
		log.Fatal(err)
	}
	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

func TestClient_GetUserWithoutPermissions(t *testing.T) {
	initTestClient()

	data, err := testClient.GetPermissionsWithoutUser("XKlwGg0zdvRdPyYEx4p9", types.UserWithoutPermissionRequest{
		ResourceID: "LQKkOjXPMpVqBzEN5og2",
		Scope:      "",
	})
	if err != nil {
		log.Fatal(err)
	}
	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

func TestClient_GetUserRoles(t *testing.T) {
	initTestClient()

	data, err := testClient.GetUserRoles("XKlwGg0zdvRdPyYEx4p9")
	if err != nil {
		log.Fatal(err)
	}
	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

func TestClient_GetUserWithoutRoles(t *testing.T) {
	initTestClient()

	data, err := testClient.GetRolesWithoutUser("XKlwGg0zdvRdPyYEx4p9", types.UserWithoutRoleRequest{})
	if err != nil {
		log.Fatal(err)
	}
	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

func TestClient_GetUserGrantApps(t *testing.T) {
	initTestClient()

	data, err := testClient.GetUserGrantApps("LQKkOjXPMpVqBzEN5og2")
	if err != nil {
		log.Fatal(err)
	}
	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

func TestClient_UserAssignPermissions(t *testing.T) {
	initTestClient()

	err := testClient.UserAssignPermissions("LQKkOjXPMpVqBzEN5og2",
		types.UserAssignPermissionReq{PermissionIDs: []string{"jARyX8kl1WO14BzJVnw6"}})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("assign permission success")
}

func TestClient_UserRemovePermission(t *testing.T) {
	initTestClient()

	err := testClient.UserRemovePermission("LQKkOjXPMpVqBzEN5og2", "jARyX8kl1WO14BzJVnw6")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("remove success")

}

func TestClient_UserAssignRoles(t *testing.T) {
	initTestClient()

	err := testClient.UserAssignRoles("LQKkOjXPMpVqBzEN5og2",
		types.UserAssignRolesReq{RoleIDs: []string{"LQKkOjXPMpVqBzEN5og2"}})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("assign roles success")
}

func TestClient_UserRemoveRole(t *testing.T) {
	initTestClient()

	err := testClient.UserRemoveRole("LQKkOjXPMpVqBzEN5og2", "LQKkOjXPMpVqBzEN5og2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("remove success")

}

func TestClient_UserRemoveApp(t *testing.T) {
	initTestClient()

	err := testClient.UserRemoveApp("XKlwGg0zdvRdPyYEx4p9", "QwnlJYpO1kO15RD0em4G")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("remove success")
}
