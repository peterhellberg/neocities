package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestNeocitiesCommand(t *testing.T) {
	t.Run("outputs usage instructions if no args", func(t *testing.T) {
		out := execGo("run", "main.go")
		expected := "neocities <command> [<args>]"

		if !strings.Contains(out, expected) {
			t.Fatalf("%q does not contain %q", out, expected)
		}
	})

	t.Run("can output help for command", func(t *testing.T) {
		out := execGo("run", "main.go", "help", "version")
		expected := "Show the version number of the neocities client"

		if !strings.Contains(out, expected) {
			t.Fatalf("%q does not contain %q", out, expected)
		}
	})

	t.Run("outputs note about missing NEOCITIES_USER variable", func(t *testing.T) {
		os.Setenv("NEOCITIES_USER", "")

		out := execGo("run", "main.go", "upload", "LICENSE")
		expected := "Error: Missing environment variable NEOCITIES_USER or NEOCITIES_KEY"

		if !strings.Contains(out, expected) {
			t.Fatalf("%q does not contain %q", out, expected)
		}
	})

	t.Run("outputs note about missing NEOCITIES_PASS variable", func(t *testing.T) {
		os.Setenv("NEOCITIES_USER", "foo")
		os.Setenv("NEOCITIES_PASS", "")

		out := execGo("run", "main.go", "upload", "LICENSE")
		expected := "Error: Missing environment variable NEOCITIES_PASS"

		if !strings.Contains(out, expected) {
			t.Fatalf("%q does not contain %q", out, expected)
		}
	})
}

func execGo(args ...string) string {
	out, _ := exec.Command("go", args...).CombinedOutput()

	return string(out)
}
