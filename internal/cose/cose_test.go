// Package cose provides ...
package cose

import (
	"crypto/ed25519"
	"encoding/hex"
	"testing"
)

func TestEdDSAPublicKey_Verify(t *testing.T) {
	data, err := hex.DecodeString("d75a980182b10ab7d54bfed3c964073a0ee172f3daa62325af021a68f707511a")
	if err != nil {
		t.Fatal(err)
	}
	sig, err := hex.DecodeString("e5564300c360ac729086e2cc806e828a84877f1eb8e5d974d873e065224901555fb8821590a33bacc61e39701cf9b46bd25bf5f0595bbe24655141438e7a100b")
	if err != nil {
		t.Fatal(err)
	}
	pubKey := ed25519.PublicKey(data)
	ok := ed25519.Verify(pubKey, nil, sig)
	t.Log(ok)
}
