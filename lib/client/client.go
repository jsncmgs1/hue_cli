package client

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var c *hueClient

// New returns a hueClient, essentially a wrapper for http.Client
// with a Put method.
func New() *hueClient {
	if c == nil {
		c = &hueClient{http.Client{}}
	}
	return c
}

type hueClient struct {
	http.Client
}

func (hc *hueClient) Put(url string, data io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPut, url, data)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return hc.Do(req)
}

type JSONResponse struct {
	Error error
	Body  map[string]map[string]string
}

func (hc *hueClient) GetJSON(url string, resultJSON map[string]map[string]string) JSONResponse {
	response := JSONResponse{}
	resp, err := hc.Get(url)
	if err != nil {
		response.Error = err
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &resultJSON)
	response.Body = resultJSON
	return response
}
