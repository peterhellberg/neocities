package client

import "testing"

func TestDelete(t *testing.T) {
	t.Run("cmdDelete", func(t *testing.T) {
		d := cmdDelete

		if got, want := d.Usage, "delete <filename> [<another filename>]"; got != want {
			t.Fatalf("d.Usage = %q, want %q", got, want)
		}

		if got, want := d.Short, "Delete files from Neocities"; got != want {
			t.Fatalf("d.Short = %q, want %q", got, want)
		}

		if got, want := d.Long, "Delete files from your Neocities website"; got != want {
			t.Fatalf("d.Long = %q, want %q", got, want)
		}
	})
}
