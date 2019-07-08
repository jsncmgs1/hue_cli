package httpclient

import (
  "bytes"
  "net/http"
  "github.com/imroc/req"
)

type Client struct {
}

func(client *Client) Get(url string) (*http.Response, error) {
  resp, err := http.Get(url)
  return resp, err
}

func(client *Client) Put(url string, body string) (*req.Resp, error) {
  jsonStr := []byte(body)
  req, err:= req.Put(url, bytes.NewBuffer(jsonStr))
  return req, err
}
