package commands

import (
	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func TestVersion(t *testing.T) {
	Convey("Version", t, func() {
		So(Version, ShouldEqual, "0.0.1")

		v := cmdVersion

		Convey("Usage", func() {
			So(v.Usage, ShouldEqual, "version")
		})

		Convey("Short", func() {
			So(v.Short, ShouldEqual, "Show neocities version")
		})

		Convey("Long", func() {
			So(v.Long, ShouldEqual, "Show the version number of the neocities client")
		})
	})
}
