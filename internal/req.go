package internal

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/parnurzeal/gorequest"
)

const apiURL = "https://api.bitopro.com"

var reqCounter = 0
var counterLocker = sync.RWMutex{}

var pubCounter = 0
var proxyIps []string

// GetReqCounter yields request counters
func GetReqCounter() int {
	return reqCounter
}

// SetupProxyIPs func
func SetupProxyIPs(proxyServerS []string) {
	proxyIps = proxyServerS
}

func upCount() {
	counterLocker.Lock()
	reqCounter++
	defer counterLocker.Unlock()
}

// ReqProxyPublic func
func ReqProxyPublic(api string) (int, string) {
	go upCount()
	pubCounter++
	proxyCounter := len(proxyIps)
	rpmAPIUrl := apiURL
	if proxyCounter > 0 {
		index := pubCounter % proxyCounter
		rpmAPIUrl = proxyIps[index]
	}
	req := gorequest.New().Timeout(500 * time.Millisecond).Get(fmt.Sprintf("%s/%s", rpmAPIUrl, api))
	req.Set("X-BITOPRO-API", "golang")
	res, body, _ := req.End()
	if res == nil {
		return -1, rpmAPIUrl
	}
	return res.StatusCode, body
}

// ReqPublic func
func ReqPublic(api string) (int, string) {
	req := gorequest.New().Get(fmt.Sprintf("%s/%s", apiURL, api))

	req.Set("X-BITOPRO-API", "golang")

	res, body, _ := req.End()

	return res.StatusCode, body
}

// ReqWithoutBody func
func ReqWithoutBody(identity, apiKey, apiSecret, method, endpoint string) (int, string) {
	go upCount()
	payload := getNonPostPayload(identity, GetTimestamp())
	sig := getSig(apiSecret, payload)
	req := gorequest.New()
	url := fmt.Sprintf("%s/%s", apiURL, endpoint)

	switch strings.ToUpper(method) {
	case "GET":
		req = req.Get(url)
	case "DELETE":
		req = req.Delete(url)
	default:
		return http.StatusMethodNotAllowed, fmt.Sprintf("Method Not Allowed: %s", method)
	}

	req.Set("X-BITOPRO-APIKEY", apiKey)
	req.Set("X-BITOPRO-PAYLOAD", payload)
	req.Set("X-BITOPRO-SIGNATURE", sig)
	req.Set("X-BITOPRO-API", "golang")

	res, body, _ := req.End()

	return res.StatusCode, body
}

// ReqWithBody func
func ReqWithBody(identity, apiKey, apiSecret, endpoint string, param map[string]interface{}) (int, string) {
	go upCount()
	body, payload := getPostPayload(param)
	sig := getSig(apiSecret, payload)
	url := fmt.Sprintf("%s/%s", apiURL, endpoint)
	req := gorequest.New().Post(url)

	req.Set("X-BITOPRO-APIKEY", apiKey)
	req.Set("X-BITOPRO-PAYLOAD", payload)
	req.Set("X-BITOPRO-SIGNATURE", sig)
	req.Set("X-BITOPRO-API", "golang")
	req.Send(body)

	res, body, _ := req.End()

	return res.StatusCode, body
}
