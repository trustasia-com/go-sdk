package authentication

import (
	"fmt"
	"testing"
)

func TestNewAuthentication(t *testing.T) {
	c, err := NewClient("EHYt7Kd5rACN16AK0viOiSXsfUT2CP3X", "https://goll.auth.wekey.com")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(c.ClientID)
}
