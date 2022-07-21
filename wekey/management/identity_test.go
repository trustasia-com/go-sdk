package management

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/trustasia-com/go-sdk/wekey/management/types"
)

func TestClient_GetIdpConfigs(t *testing.T) {

	initTestClient()

	data, err := testClient.GetIdpConfigs(types.IdpConfigsReq{Type: "social"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

func TestClient_GetIdpApps(t *testing.T) {
	initTestClient()

	data, err := testClient.GetIdpApps(types.IdpAppsReq{
		Provider:   "github",
		Identifier: "github",
	})
	if err != nil {
		log.Fatal(err)
	}
	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}
