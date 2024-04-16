package httpClient

import (
	"bytes"
	"net/http"
)

type Client interface {
	Get(url string) (*http.Response, error)
	Post(url string, body []byte) (*http.Response, error)
}

type DefaultClient struct{}

func (c *DefaultClient) Get(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *DefaultClient) Post(url string, body []byte) (*http.Response, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	return resp, nil
}
