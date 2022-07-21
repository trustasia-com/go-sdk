package authentication

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/trustasia-com/go-sdk/wekey/authentication/types"
)

const (
	testClientID = "Szxty4O5J5oY0enQUFMvgfTroiDKrUbX"
	testEndpoint = "https://goll.auth.wekey.com"
)

func TestAuthenticationClient_LoginByPassword(t *testing.T) {

	c, err := NewClient(testClientID, testEndpoint)
	if err != nil {
		t.Fatal(err)
	}
	req := types.LoginByPasswordRequest{
		Target:   "goll",
		Password: "123456",
	}
	user, err := c.LoginByPassword(req)
	if err != nil {
		t.Fatal(err)
	}
	info, _ := json.Marshal(user)
	fmt.Println(string(info))
}

func TestAuthenticationClient_SendPasscode(t *testing.T) {
	c, err := NewClient(testClientID, testEndpoint)
	if err != nil {
		t.Fatal(err)
	}
	req := types.SendPasscodeRequest{
		Target: "18302842579",
		Action: "login",
	}
	err = c.SendPasscode(req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAuthenticationClient_LoginByPasscode(t *testing.T) {
	c, err := NewClient(testClientID, testEndpoint)
	if err != nil {
		t.Fatal(err)
	}
	req := types.LoginByPasscodeRequest{
		Target:   "18302842579",
		Passcode: "329061",
	}
	user, err := c.LoginByPasscode(req)
	if err != nil {
		t.Fatal(err)
	}
	info, _ := json.Marshal(user)
	fmt.Println(string(info))
}
