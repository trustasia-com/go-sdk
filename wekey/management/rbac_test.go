package management

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/trustasia-com/go-sdk/wekey/management/types"
)

func TestClient_QueryRoles(t *testing.T) {
	initTestClient()

	data, err := testClient.QueryRoles(types.QueryRoleRequest{Name: ""})
	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

func TestClient_RoleInfo(t *testing.T) {
	initTestClient()

	data, err := testClient.RoleInfo("LQKkOjXPMpVqBzEN5og2")
	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.Marshal(data)
	fmt.Println(string(b))

}

func TestClient_CreateRole(t *testing.T) {
	initTestClient()

	id, err := testClient.CreateRole(types.CreateRoleRequest{
		Name:        "dev admin",
		Description: "test info",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
}

func TestClient_DelRole(t *testing.T) {
	initTestClient()

	err := testClient.DelRole("RxB4Qb82dRQqXVPykEYz")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("del role success")
}

func TestClient_UpdateRole(t *testing.T) {
	initTestClient()

	err := testClient.UpdateRole("LX0wBJg8da31EO67VrGv", types.UpdateRoleRequest{
		Name:        "update role",
		Description: "123321",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("update role success")
}

func TestClient_GetRolePermissions(t *testing.T) {
	initTestClient()

	data, err := testClient.GetRolePermissions("LQKkOjXPMpVqBzEN5og2")
	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

func TestClient_GetRoleUsers(t *testing.T) {
	initTestClient()

	data, err := testClient.GetRoleUsers("LQKkOjXPMpVqBzEN5og2", types.BaseQueryRequest{
		Page: 1,
		Size: 10,
	})
	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

func TestClient_RoleAddPermission(t *testing.T) {
	initTestClient()

	err := testClient.RoleAddPermission("LQKkOjXPMpVqBzEN5og2",
		types.AddRolePermissionRequest{PermissionIDs: []string{"jARyX8kl1WO14BzJVnw6"}})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("add success")
}

func TestClient_RoleAddUsers(t *testing.T) {
	initTestClient()

	err := testClient.RoleAddUsers("LQKkOjXPMpVqBzEN5og2", types.RoleAssignUsersReq{UserIDs: []string{
		"RxB4Qb82dRQqXVPykEYz",
	}})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("add success")
}

func TestClient_RoleRemovePermission(t *testing.T) {
	initTestClient()

	err := testClient.RoleRemovePermission("LQKkOjXPMpVqBzEN5og2", "LQKkOjXPMpVqBzEN5og2")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("remove success")
}

func TestClient_GetPermissionsWithoutRole(t *testing.T) {
	initTestClient()

	data, err := testClient.GetPermissionsWithoutRole("LQKkOjXPMpVqBzEN5og2",
		types.QueryPermissionsWithoutRoleRequest{
			ResourceID: "LQKkOjXPMpVqBzEN5og2",
			Scope:      "",
		})
	if err != nil {
		log.Fatal(err)
	}
	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

func TestClient_GetUsersWithoutRole(t *testing.T) {
	initTestClient()

	data, err := testClient.GetUsersWithoutRole("LQKkOjXPMpVqBzEN5og2", types.QueryUserWithoutRoleRequest{})
	if err != nil {
		log.Fatal(err)
	}
	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}
