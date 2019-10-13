package bitopro

import (
	"github.com/Eden-Sun/bitopro-api-go/internal"
)

// PubAPI struct
type PubAPI struct{}

// AuthAPI struct
type AuthAPI struct {
	identity string
	key      string
	secret   string
}

// GetPubClient func
func GetPubClient() *PubAPI {
	return &PubAPI{}
}

// GetAuthClient func
func GetAuthClient(identity, key, secret string) *AuthAPI {
	return &AuthAPI{
		identity: identity,
		key:      key,
		secret:   secret,
	}
}

// StatusCode struct
type StatusCode struct {
	Code  int    `json:"code"`
	Error string `json:"error,omitempty"`
}

// GetReqCounter get total counter
func GetReqCounter() int {
	return internal.GetReqCounter()
}

// SetupProxyIPs func
func SetupProxyIPs(proxyIPs []string) {
	internal.SetupProxyIPs(proxyIPs)
}
