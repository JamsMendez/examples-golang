package handler

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

type client struct {
	baseURL    string
	apiKey     string
	clientHTTP *http.Client
}

func NewClient(baseURL, apiKey string) *client {
	return &client{
		baseURL: baseURL,
		apiKey:  apiKey,
		clientHTTP: &http.Client{
			Timeout: 3 * time.Second,
		},
	}
}

func (c *client) makeRequest(method, p string, body io.Reader) (*http.Response, error) {
	u, err := c.getURL(p)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u, body)
	if err != nil {
		return nil, err
	}

	if c.apiKey != "" {
		req.Header.Add("Access-Token", c.apiKey)
	}

	return c.clientHTTP.Do(req)
}

func (c *client) getURL(p string) (string, error) {
	return url.JoinPath(c.baseURL, p)
}
