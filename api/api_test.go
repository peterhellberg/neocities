package api

import "testing"

func TestAPI(t *testing.T) {
	if got, want := apiURL, "https://neocities.org/api/"; got != want {
		t.Fatalf("apiURL = %q, want %q", got, want)
	}

	if got, want := userAgent, "neocities (Go package using net/http)"; got != want {
		t.Fatalf("userAgent = %q, want %q", got, want)
	}
}
