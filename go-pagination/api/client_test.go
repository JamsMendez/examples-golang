package api

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	baseURL    string
	apiKey     string
	clientHTTP *http.Client
}

func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		baseURL: baseURL,
		apiKey:  apiKey,
		clientHTTP: &http.Client{
			Timeout: 3 * time.Second,
		},
	}
}

func (c *Client) makeRequest(method, endpoint string, queries map[string]string, body io.Reader) (*http.Response, error) {
	u, err := c.getURL(endpoint)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u, body)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()

	for key, value := range queries {
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()

	return c.clientHTTP.Do(req)
}

func (c *Client) getURL(p string) (string, error) {
	return url.JoinPath(c.baseURL, p)
}
