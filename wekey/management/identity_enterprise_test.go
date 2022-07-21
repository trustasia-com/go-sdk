package management

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/trustasia-com/go-sdk/wekey/management/types"
)

func TestClient_AddIdpLDAP(t *testing.T) {
	initTestClient()
	err := testClient.AddIdpLDAP(types.AddIdpLDAPRequest{
		Name:        "testldap",
		Identifier:  "test-ldap",
		SyncProfile: true,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func TestClient_GetIdpLDAPConfig(t *testing.T) {
	initTestClient()
	data, err := testClient.GetIdpLDAPConfig("test-ldap")
	if err != nil {
		log.Fatal(err)
	}
	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

func TestClient_DelIdpLDAP(t *testing.T) {
	initTestClient()
	err := testClient.DelIdpLDAP("test-ldap")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("del success")
}

func TestClient_UpdateIdpLDAP(t *testing.T) {
	initTestClient()
	err := testClient.UpdateIdpLDAP("test-ldap", types.UpdateIdpLDAP{
		Name:        "testldap",
		SyncProfile: true,
	})
	if err != nil {
		log.Fatal(err)
	}
}
