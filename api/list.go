package api

import (
	"encoding/json"
	"net/http"
)

type ListResponse struct {
	Result string `json:"result"`
	Files  []File `json:"files"`
}

type File struct {
	Path        string `json:"path"`
	IsDirectory bool   `json:"is_directory"`
	Size        int    `json:"size,omitempty"`
	UpdatedAt   string `json:"updated_at"`
	Sha1Hash    string `json:"sha1_hash,omitempty"`
}

func List(a Authenticator) (*ListResponse, error) {
	req, err := newListRequest(a)
	if err != nil {
		return nil, err
	}

	resp, err := sendHTTPRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var lr ListResponse

	return &lr, json.NewDecoder(resp.Body).Decode(&lr)
}

func newListRequest(a Authenticator) (*http.Request, error) {
	endpoint := "list"

	req, err := http.NewRequest("GET", apiURL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	if a != nil {
		a.Authenticate(req)
	}

	return req, nil
}
