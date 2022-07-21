package management

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestClient_GetIdpWebauthnConfig(t *testing.T) {
	initTestClient()
	data, err := testClient.GetIdpWebauthnConfig()
	if err != nil {
		log.Fatal(err)
	}
	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

func TestClient_AddIdpWebauthn(t *testing.T) {
	initTestClient()
	err := testClient.AddIdpWebauthn()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("add webauthn success")
}

func TestClient_DelIdpWebauthn(t *testing.T) {
	initTestClient()
	err := testClient.DelIdpWebauthn()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("del webauthn success")
}
