package commands

import (
	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func TestUpload(t *testing.T) {
	Convey("Upload", t, func() {
		u := cmdUpload

		Convey("Usage", func() {
			usage := "upload <filename> [<another filename>]"
			So(u.Usage, ShouldEqual, usage)
		})

		Convey("Short", func() {
			short := "Upload files to Neocities"
			So(u.Short, ShouldEqual, short)
		})

		Convey("Long", func() {
			long := "Upload files to your Neocities account"
			So(u.Long, ShouldEqual, long)
		})
	})
}
