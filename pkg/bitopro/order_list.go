package bitopro

import (
	"encoding/json"
	"fmt"

	"github.com/Eden-Sun/bitopro-api-go/internal"
)

// OrderList struct
type OrderList struct {
	Data       []OrderInfo `json:"data,omitempty"`
	Page       int         `json:"page"`
	TotalPages int         `json:"totalPages"`
	StatusCode
}

// GetOrderList Ref. https://developer.bitopro.com/docs#operation/getOrders
func (api *AuthAPI) GetOrderList(pair string, active bool, page uint) *OrderList {
	var data OrderList

	code, res := internal.ReqWithoutBody(api.identity, api.key, api.secret, "GET", fmt.Sprintf("%s/%s?page=%d&active=%v", "v2/orders", pair, page, active))

	if err := json.Unmarshal([]byte(res), &data); err != nil {
		data.Error = res
	}

	data.Code = code

	return &data
}
