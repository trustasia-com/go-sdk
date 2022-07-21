package management

import (
	"fmt"
	"log"
	"testing"
)

const (
	testClientID     = "98ZF3qR80ajFGbaT11hCvFvylQbcecO7"
	testClientSecret = "Brk46xSq0rtjFcyBifPtLGcKFJhCPlsGrDiZkFTo"
	testEndpoint     = "https://goll.auth.wekey.cn"
)

var testClient = &Client{}

func initTestClient() {
	s := &MemoryStorage{}
	c, err := NewClient(testClientID, testClientSecret, testEndpoint, s)
	if err != nil {
		log.Fatal(err)
	}
	testClient = c
}

func TestNewClient(t *testing.T) {
	s := &MemoryStorage{}
	c, err := NewClient(testClientID, testClientSecret, testEndpoint, s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c.ClientID)
}
