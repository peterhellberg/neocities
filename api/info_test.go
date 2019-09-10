package api

import "testing"

func TestInfo(t *testing.T) {
	cred := Credentials{User: "foo", Pass: "bar", Key: "baz"}

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
}
