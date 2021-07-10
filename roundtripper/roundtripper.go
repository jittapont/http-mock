package roundtripper

import (
	"log"
	"net/http"
)

//go:generate mockgen -destination=./mock_roundtripper.go -package=roundtripper http-mock/roundtripper RoundTripper
type RoundTripper interface {
	RoundTrip(*http.Request) (*http.Response, error)
}

func makeGetRequest(client *http.Client, req *http.Request) (*http.Response, error) {
	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}
	log.Printf("Response status code : %d", resp.StatusCode)
	return resp, nil
}
