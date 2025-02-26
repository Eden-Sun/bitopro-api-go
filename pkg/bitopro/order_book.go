package bitopro

import (
	"encoding/json"
	"fmt"

	"github.com/Eden-Sun/bitopro-api-go/internal"
)

// OrderBookInfo struct
type OrderBookInfo struct {
	Amount string `json:"amount"`
	Price  string `json:"price"`
	Count  int    `json:"count"`
	Total  string `json:"total"`
}

// OrderBook struct
type OrderBook struct {
	Bids []OrderBookInfo `json:"bids"`
	Asks []OrderBookInfo `json:"asks"`
	StatusCode
}

func getOrderBook(pair string, limit int) *OrderBook {
	var data OrderBook

	code, res := internal.ReqPublic(fmt.Sprintf("%s/%s?limit=%d", "v2/order-book", pair, limit))

	if err := json.Unmarshal([]byte(res), &data); err != nil {
		data.Error = res
	}

	data.Code = code

	return &data
}

// GetOrderBook Ref. https://developer.bitopro.com/docs#operation/getOrderBookByPair
func (*PubAPI) GetOrderBook(pair string) *OrderBook {
	return getOrderBook(pair, 5)
}

// GetOrderBook Ref. https://developer.bitopro.com/docs#operation/getOrderBookByPair
func (*AuthAPI) GetOrderBook(pair string) *OrderBook {
	return getOrderBook(pair, 5)
}

// GetOrderBookWithLimit Ref. https://developer.bitopro.com/docs#operation/getOrderBookByPair
func (*PubAPI) GetOrderBookWithLimit(pair string, limit int) *OrderBook {
	return getOrderBook(pair, limit)
}

// GetOrderBookWithLimit Ref. https://developer.bitopro.com/docs#operation/getOrderBookByPair
func (*AuthAPI) GetOrderBookWithLimit(pair string, limit int) *OrderBook {
	return getOrderBook(pair, limit)
}

// GetOrderBookThroughProxy func
func (*PubAPI) GetOrderBookThroughProxy(pair string, limit int) *OrderBook {
	var data OrderBook
	code, res := internal.ReqProxyPublic(fmt.Sprintf("%s/%s?limit=%d", "v3/order-book", pair, limit))
	if code != 200 {
		data.Code = code
		data.Error = res
		return &data
	}
	if err := json.Unmarshal([]byte(res), &data); err != nil {
		data.Error = res
	}

	data.Code = code

	return &data
}
