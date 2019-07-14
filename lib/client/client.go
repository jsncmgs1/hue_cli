package client

import (
	"io"
	"log"
	"net/http"
)

func New() *HueClient {
	return &HueClient{&http.Client{}}
}

type HueClient struct {
	*http.Client
}

func (hc *HueClient) Put(url string, data io.Reader) {
	req, err := http.NewRequest(http.MethodPut, url, data)

	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = hc.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
}
