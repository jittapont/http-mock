package roundtripper

import "net/http"

//go:generate mockgen -destination=./mock_roundtripper.go -package=roundtripper http-mock/roundtripper RoundTripper
type RoundTripper interface {
	RoundTrip(*http.Request) (*http.Response, error)
}
