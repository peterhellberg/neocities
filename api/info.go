package api

import "net/http"

// SiteInfo returns a site info response
func SiteInfo(a Authenticator, site string) (Response, error) {
	req, err := newInfoRequest(a, site)
	if err != nil {
		return Response{}, err
	}

	return performHTTPRequest(req)
}

// Create a new info request
func newInfoRequest(a Authenticator, site string) (*http.Request, error) {
	endpoint := "info"

	if site != "" {
		endpoint = endpoint + "?sitename=" + site
	}

	req, err := http.NewRequest("GET", apiURL+endpoint, nil)
	if err != nil {
		return req, err
	}

	if a != nil {
		a.Authenticate(req)
	}

	return req, nil
}
