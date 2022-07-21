// Package valid provides ...
package valid

import "testing"

func TestIsChinaPhoneNo(t *testing.T) {
	tests := []string{
		"15112341234",
	}

	for _, v := range tests {
		t.Log(IsChinaPhoneNo(v))
	}
}

func TestIsASCIIEmail(t *testing.T) {
	tests := []string{
		"12333@163.com",
	}

	for _, v := range tests {
		t.Log(IsASCIIEmail(v))
	}
}
