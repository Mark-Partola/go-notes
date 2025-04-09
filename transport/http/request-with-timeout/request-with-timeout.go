package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	res, err := request("https://example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

func request(address string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2000*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, address, nil)
	if err != nil {
		return "", err
	}

	res, err := New().Request(req)
	if err != nil {
		return "", err
	}

	return res, nil
}

type HTTPClient struct {
	client *http.Client
}

func New() *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *HTTPClient) Request(req *http.Request) (string, error) {
	res, err := c.client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", nil
	}

	return string(bytes), nil
}
