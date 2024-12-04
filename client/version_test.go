package client

import "testing"

func TestVersion(t *testing.T) {
	if got, want := Version, "0.0.4"; got != want {
		t.Fatalf("Version = %q, want %q", got, want)
	}

	t.Run("cmdVersion", func(t *testing.T) {
		v := cmdVersion

		if got, want := v.Usage, "version"; got != want {
			t.Fatalf("v.Usage = %q, want %q", got, want)
		}

		if got, want := v.Short, "Show neocities version"; got != want {
			t.Fatalf("v.Short = %q, want %q", got, want)
		}

		if got, want := v.Long, "Show the version number of the neocities client"; got != want {
			t.Fatalf("v.Long = %q, want %q", got, want)
		}
	})
}
