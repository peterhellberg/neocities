package client

import "testing"

func TestInfo(t *testing.T) {
	t.Run("cmdInfo", func(t *testing.T) {
		i := cmdInfo

		if got, want := i.Usage, "info [sitename]"; got != want {
			t.Fatalf("i.Usage = %q, want %q", got, want)
		}

		if got, want := i.Short, "Info about Neocities websites"; got != want {
			t.Fatalf("i.Short = %q, want %q", got, want)
		}

		if got, want := i.Long, "Info about your Neocities website, or somebody elses"; got != want {
			t.Fatalf("i.Long = %q, want %q", got, want)
		}
	})
}
