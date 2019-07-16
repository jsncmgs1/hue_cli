package client

import (
	"io"
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

	res, err := hc.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}
