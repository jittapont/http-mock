package client

import (
	"log"
	"net/http"
)

//go:generate mockgen -destination=./mock_client.go -package=client http-mock/client Client
type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

func makeGetRequest(client Client, req *http.Request) (*http.Response, error) {
	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}
	log.Printf("Response status code : %d", resp.StatusCode)
	return resp, nil
}
