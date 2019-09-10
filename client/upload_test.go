package client

import "testing"

func TestUpload(t *testing.T) {
	t.Run("cmdUpload", func(t *testing.T) {
		u := cmdUpload

		if got, want := u.Usage, "upload <filename> [<another filename>]"; got != want {
			t.Fatalf("u.Usage = %q, want %q", got, want)
		}

		if got, want := u.Short, "Upload files to Neocities"; got != want {
			t.Fatalf("u.Short = %q, want %q", got, want)
		}

		if got, want := u.Long, "Upload files to your Neocities website"; got != want {
			t.Fatalf("u.Long = %q, want %q", got, want)
		}
	})
}
