package client

import (
	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func TestInfo(t *testing.T) {
	Convey("Info", t, func() {
		i := cmdInfo

		Convey("Usage", func() {
			usage := "info [sitename]"
			So(i.Usage, ShouldEqual, usage)
		})

		Convey("Short", func() {
			short := "Info about Neocities websites"
			So(i.Short, ShouldEqual, short)
		})

		Convey("Long", func() {
			long := "Info about your Neocities website, or somebody elses"
			So(i.Long, ShouldEqual, long)
		})
	})
}
