package main

import (
	. "github.com/smartystreets/goconvey/convey"

	"os"
	"os/exec"
	"testing"
)

func TestNeocitiesCommand(t *testing.T) {
	Convey("neocities", t, func() {
		Convey("outputs usage instructions if no args", func() {
			out, _ := execGo("run", "main.go")

			So(out, ShouldContainSubstring, "neocities <command> [<args>]")
		})

		Convey("can output help for command", func() {
			out, _ := execGo("run", "main.go", "help", "version")

			expected := "Show the version number of the Neocities API client"

			So(out, ShouldContainSubstring, expected)
		})

		Convey("outputs note about missing NEOCITIES_USER variable", func() {
			os.Setenv("NEOCITIES_USER", "")

			out, _ := execGo("run", "main.go", "upload", "LICENSE")

			expected := "Error: Missing environment variable NEOCITIES_USER"

			So(out, ShouldContainSubstring, expected)
		})

		Convey("outputs note about missing NEOCITIES_PASS variable", func() {
			os.Setenv("NEOCITIES_USER", "foo")
			os.Setenv("NEOCITIES_PASS", "")

			out, _ := execGo("run", "main.go", "upload", "LICENSE")

			expected := "Error: Missing environment variable NEOCITIES_PASS"

			So(out, ShouldContainSubstring, expected)
		})
	})
}

func execGo(args ...string) (string, error) {
	out, err := exec.Command("go", args...).CombinedOutput()

	return string(out), err
}
