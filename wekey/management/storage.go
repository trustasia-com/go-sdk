package management

import (
	"crypto/rsa"
	"time"
)

// Storage save your access token you can save it into redis/db
// 请注意,通常情况下，您在调用sdk时应当自己实现一个storage
// storage 是用来存储access token和public key的
type Storage interface {
	SetToken(token string) error
	GetToken(clientID, clientSecret, endpoint string) (string, error)
	RefreshToken(clientID, clientSecret, endpoint string) (string, error)
	SetPublicKey(publicKey *rsa.PublicKey) error
	GetPublicKey() (*rsa.PublicKey, error)
}

// MemoryStorage test in memory
// 这是一个用于测试的storage，通常情况下，您不应该使用它替代实际的storage
type MemoryStorage struct {
	Token      string
	PublicKey  *rsa.PublicKey
	ExpireTime time.Time
}

// SetToken save token in memory
func (s *MemoryStorage) SetToken(token string) error {
	s.Token = token
	s.ExpireTime = time.Now().Add(1209600 * time.Second)
	return nil
}

// GetToken get token
func (s *MemoryStorage) GetToken(clientID, clientSecret, endpoint string) (string, error) {
	if s.ExpireTime.After(time.Now()) {
		return s.Token, nil
	}
	return s.RefreshToken(clientID, clientSecret, endpoint)
}

// RefreshToken refresh token when token is expired
func (s *MemoryStorage) RefreshToken(clientID, clientSecret, endpoint string) (string, error) {
	token, err := GetAccessToken(clientID, clientSecret, endpoint)
	if err != nil {
		return "", err
	}
	return token, s.SetToken(token)
}

// SetPublicKey save zone's public key
func (s *MemoryStorage) SetPublicKey(publicKey *rsa.PublicKey) error {
	s.PublicKey = publicKey
	return nil
}

// GetPublicKey get zone' public key
func (s *MemoryStorage) GetPublicKey() (*rsa.PublicKey, error) {
	return s.PublicKey, nil
}
