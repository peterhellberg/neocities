package api

import (
	"encoding/json"
	"net/http"
)

type KeyResponse struct {
	Result    string `json:"result"`
	ErrorType string `json:"error_type"`
	Message   string `json:"message"`
	APIKey    string `json:"api_key"`
}

func Key(a Authenticator) (*KeyResponse, error) {
	req, err := newKeyRequest(a)
	if err != nil {
		return nil, err
	}

	resp, err := sendHTTPRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var kr KeyResponse

	return &kr, json.NewDecoder(resp.Body).Decode(&kr)
}

func newKeyRequest(a Authenticator) (*http.Request, error) {
	endpoint := "key"

	req, err := http.NewRequest("GET", apiURL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	if a != nil {
		a.Authenticate(req)
	}

	return req, nil
}
