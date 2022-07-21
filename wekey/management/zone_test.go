package management

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/trustasia-com/go-sdk/wekey/management/types"
)

func TestClient_GetZoneBasicInfo(t *testing.T) {
	initTestClient()

	data, err := testClient.GetZoneBasicInfo()
	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

func TestClient_GetZoneMembersInfo(t *testing.T) {
	initTestClient()

	data, err := testClient.GetZoneMembersInfo()
	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

func TestClient_GetZoneSecretsInfo(t *testing.T) {
	initTestClient()

	data, err := testClient.GetZoneSecretsInfo()
	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

func TestClient_GetZoneAdvanceInfo(t *testing.T) {
	initTestClient()

	data, err := testClient.GetZoneAdvanceInfo()
	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

func TestClient_UpdateZoneBasicInfo(t *testing.T) {
	initTestClient()

	err := testClient.UpdateZoneBasicInfo(types.UpdateZoneBasicInfoReq{
		Name:         "testsaaa",
		Description:  "oasjaklsdmsnzxm",
		Logo:         "https://static.wekey.com/web/wekey-logo.png",
		EmailSupport: "goll.wang@trustasia.com",
		SupportUrl:   "https://wekey.com",
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("update success")
}

func TestClient_UpdateZoneMembersInfo(t *testing.T) {
	initTestClient()

	err := testClient.UpdateZoneMembersInfo("jARyX8kl1WO14BzJVnw6",
		types.UpdateZoneMembersReq{Roles: []string{"asa"}})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("update success")
}

func TestClient_DeleteZoneMembers(t *testing.T) {
	initTestClient()

	err := testClient.DeleteZoneMembers("jARyX8kl1WO14BzJVnw6")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("delete success")
}

func TestClient_DeleteZoneSecrets(t *testing.T) {
	initTestClient()

	err := testClient.DeleteZoneSecrets()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("delete success")
}
