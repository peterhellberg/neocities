package api

import "testing"

func TestDeleteFiles(t *testing.T) {
	cred := Credentials{User: "foo", Pass: "bar", Key: "baz"}

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
}
