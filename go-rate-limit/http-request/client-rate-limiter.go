package ratelimiter

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type RHTTPClient struct {
	client      *http.Client
	RateLimiter *rate.Limiter
}

func (c *RHTTPClient) Do(req *http.Request) (*http.Response, error) {
	ctx := context.Background()
	err := c.RateLimiter.Wait(ctx)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func NewClient(rl *rate.Limiter) *RHTTPClient {
	c := &RHTTPClient{
		client:      http.DefaultClient,
		RateLimiter: rl,
	}

	return c
}

func RunClient() {
	// rl := rate.NewLimiter(rate.Every(10*time.Second), 50)
	rl := rate.NewLimiter(rate.Every(1*time.Second), 15)
	c := NewClient(rl)

	reqURL := "http://localhost:3000/"
	req, _ := http.NewRequest("GET", reqURL, nil)

	defer fmt.Println("Finish")

	for i := 0; i < 300; i++ {
		response, err := c.Do(req)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(response.StatusCode)
			return
		}

		if response.StatusCode == 429 {
			fmt.Printf("Rate limit reached after %d requests\n", i)
			return
		}

		defer response.Body.Close()

		buffer, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Response ERROR: ", err)
			continue
		}

		body := string(buffer)
		fmt.Println("Request Body: ", body, i)
	}
}
