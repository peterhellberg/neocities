package api

import "net/http"

type Authenticator interface {
	Authenticate(*http.Request)
}

// Credentials contains the username and password
type Credentials struct {
	User string
	Pass string
	Key  string
}

func (c Credentials) Authenticate(r *http.Request) {
	if r == nil {
		return
	}

	if c.Key != "" {
		r.Header.Set("Authorization", "Bearer: "+c.Key)
		return
	}

	r.SetBasicAuth(c.User, c.Pass)
}
