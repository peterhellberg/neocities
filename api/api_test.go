package api

import (
	"io/ioutil"
	"testing"
)

func TestAPI(t *testing.T) {
	cred := Credentials{User: "foo", Pass: "bar", Key: "baz"}

	t.Run("expected constants", func(t *testing.T) {
		if got, want := apiURL, "https://neocities.org/api/"; got != want {
			t.Fatalf("apiURL = %q, want %q", got, want)
		}

		if got, want := userAgent, "neocities (Go package using net/http)"; got != want {
			t.Fatalf("userAgent = %q, want %q", got, want)
		}
	})

	t.Run("newInfoRequest", func(t *testing.T) {
		req, err := newInfoRequest(cred, "foo")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got, want := req.Method, "GET"; got != want {
			t.Fatalf("req.Method = %q, want %q", got, want)
		}

		if got, want := req.URL.String(), "https://neocities.org/api/info?sitename=foo"; got != want {
			t.Fatalf("req.URL.String() = %q, want %q", got, want)
		}
	})

	t.Run("newDeleteRequest", func(t *testing.T) {
		req, err := newDeleteRequest(cred, []string{"foo", "bar"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got, want := req.Method, "POST"; got != want {
			t.Fatalf("req.Method = %q, want %q", got, want)
		}

		if got, want := req.URL.String(), "https://neocities.org/api/delete"; got != want {
			t.Fatalf("req.URL.String() = %q, want %q", got, want)
		}
	})

	t.Run("newUploadRequest", func(t *testing.T) {
		req, err := newUploadRequest(cred, []string{"../LICENSE"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got, want := req.Method, "POST"; got != want {
			t.Fatalf("req.Method = %q, want %q", got, want)
		}

		if got, want := req.URL.String(), "https://neocities.org/api/upload"; got != want {
			t.Fatalf("req.URL.String() = %q, want %q", got, want)
		}
	})

	t.Run("newUploadDataRequest", func(t *testing.T) {
		testContent, err := ioutil.ReadFile("../LICENSE")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		testData := []UploadData{
			UploadData{
				FileName: "LICENSE_string",
				Content:  testContent,
			},
		}

		req, err := newUploadDataRequest(cred, testData)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got, want := req.Method, "POST"; got != want {
			t.Fatalf("req.Method = %q, want %q", got, want)
		}

		if got, want := req.URL.String(), "https://neocities.org/api/upload"; got != want {
			t.Fatalf("req.URL.String() = %q, want %q", got, want)
		}
	})
}
