package commands

import (
	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func TestHelp(t *testing.T) {
	Convey("Help", t, func() {
		h := cmdHelp

		Convey("Usage", func() {
			usage := "help [command]"
			So(h.Usage, ShouldEqual, usage)
		})

		Convey("Short", func() {
			short := "Show help"
			So(h.Short, ShouldEqual, short)
		})

		Convey("Long", func() {
			long := "Show usage instructions for a command"
			So(h.Long, ShouldEqual, long)
		})
	})
}
