package client

import "testing"

func TestHelp(t *testing.T) {
	t.Run("cmdHelp", func(t *testing.T) {
		h := cmdHelp

		if got, want := h.Usage, "help [command]"; got != want {
			t.Fatalf("h.Usage = %q, want %q", got, want)
		}

		if got, want := h.Short, "Show help"; got != want {
			t.Fatalf("h.Short = %q, want %q", got, want)
		}

		if got, want := h.Long, "Show usage instructions for a command"; got != want {
			t.Fatalf("h.Long = %q, want %q", got, want)
		}
	})
}
