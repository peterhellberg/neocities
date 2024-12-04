package api

import (
	"os"
	"testing"
)

func TestUpload(t *testing.T) {
	cred := Credentials{User: "foo", Pass: "bar", Key: "baz"}

	t.Run("newUploadDataRequest", func(t *testing.T) {
		testContent, err := os.ReadFile("../LICENSE")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		testData := []UploadData{
			{
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
