package api

import (
	"net/http"
	"net/url"
	"strings"
)

// DeleteFiles deletes the given filenames
func DeleteFiles(a Authenticator, filenames []string) (Response, error) {
	req, err := newDeleteRequest(a, filenames)
	if err != nil {
		return Response{}, err
	}

	return performHTTPRequest(req)
}

// Create a new delete request
func newDeleteRequest(a Authenticator, filenames []string) (*http.Request, error) {
	data := url.Values{}

	for _, file := range filenames {
		data.Add("filenames[]", file)
	}

	req, err := http.NewRequest("POST", apiURL+"delete", strings.NewReader(data.Encode()))
	if err != nil {
		return req, err
	}

	if a != nil {
		a.Authenticate(req)
	}

	// Set the content type of the form data
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}
