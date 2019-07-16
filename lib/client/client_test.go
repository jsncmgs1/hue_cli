package client_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	httpclient "github.com/jsncmgs1/hue_cli/lib/client"
)

func TestPut(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(201)
	}))

	client := httpclient.New()
	defer ts.Close()

	u, _ := url.Parse(ts.URL)
	res, err := client.Put(u.String(), strings.NewReader(`{"on": true}`))

	if err != nil {
		t.Fatalf("fatal error: %s", err)
	}

	if res.StatusCode != 201 {
		t.Errorf("exepcted 201 created, got %d", res.StatusCode)
	}
}
