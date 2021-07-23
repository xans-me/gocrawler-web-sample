package http

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var netTransport = &http.Transport{
	Dial: (&net.Dialer{
		Timeout:   10 * time.Second,
		KeepAlive: 30 * time.Second,
	}).Dial,
	TLSHandshakeTimeout: 10 * time.Second,
	MaxIdleConns:        100,
	MaxIdleConnsPerHost: 100,
}
var netClient = &http.Client{
	Timeout:   30 * time.Second,
	Transport: netTransport,
}

// Request function to do http request, default 30 second timeout
func Request(request *http.Request) (*http.Response, error) {
	netClient.Timeout = 30 * time.Second
	resp, err := netClient.Do(request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RequestWithTimeout to do http request with user defined timeout
func RequestWithTimeout(request *http.Request, timeoutInSecond int) (*http.Response, error) {
	netClient.Timeout = time.Duration(timeoutInSecond) * time.Second
	resp, err := netClient.Do(request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ParseResponseBody is a func to parse json response body to map of string to interface
func ParseResponseBody(response *http.Response) (data map[string]interface{}, err error) {
	// Get Response Body
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return data, err
	}

	// decode response data to get the accessToken
	err = json.Unmarshal(responseData, &data)
	return
}
