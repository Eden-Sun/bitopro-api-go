package bitopro

import (
	"encoding/json"

	"github.com/Eden-Sun/bitopro-api-go/internal"
)

// Balance struct
type Balance struct {
	Currency  string `json:"currency"`
	Amount    string `json:"amount"`
	Available string `json:"available"`
	Stake     string `json:"stake"`
}

// Account struct
type Account struct {
	Data []Balance `json:"data,omitempty"`
	StatusCode
}

// GetAccountBalance Ref. https://developer.bitopro.com/docs#operation/getAccountBalance
func (api *AuthAPI) GetAccountBalance() *Account {
	var data Account

	code, res := internal.ReqWithoutBody(api.identity, api.key, api.secret, "GET", "v2/accounts/balance")

	if err := json.Unmarshal([]byte(res), &data); err != nil {
		data.Error = res
	}

	data.Code = code

	return &data
}
