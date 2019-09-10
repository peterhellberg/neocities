package api

import (
	"errors"
	"net/http"
)

const (
	apiURL    = "https://neocities.org/api/"
	userAgent = "neocities (Go package using net/http)"
)

var ErrUnexpectedStatusCode = errors.New("unexpected status code")

func performHTTPRequest(req *http.Request) (Response, error) {
	res, err := sendHTTPRequest(req)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()

	var r Response

	if err := r.PopulateFromHTTPResponse(res); err != nil {
		return r, err
	}

	if res.StatusCode != 200 {
		return r, ErrUnexpectedStatusCode
	}

	return r, nil
}

// Send HTTP request
func sendHTTPRequest(req *http.Request) (*http.Response, error) {
	// Create a HTTP client
	client := &http.Client{}

	// Make sure that the correct User-Agent is set
	req.Header.Add("User-Agent", userAgent)

	// Send the request and return the response
	return client.Do(req)
}
