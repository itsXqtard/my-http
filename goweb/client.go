package goweb

import (
	"net/http"
	"net/url"
)

type Header map[string][]string
type Request struct {
	Method string
	URL    *url.URL
	Header Header
}

type Client struct{}

var defaultClient = &Client{}

func (c *Client) Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}

func GET(url string) (resp *http.Response, err error) {
	return defaultClient.Get(url)
}
