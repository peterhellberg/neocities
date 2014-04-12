package client

import (
	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func TestDelete(t *testing.T) {
	Convey("Delete", t, func() {
		d := cmdDelete

		Convey("Usage", func() {
			usage := "delete <filename> [<another filename>]"
			So(d.Usage, ShouldEqual, usage)
		})

		Convey("Short", func() {
			short := "Delete files from Neocities"
			So(d.Short, ShouldEqual, short)
		})

		Convey("Long", func() {
			long := "Delete files from your Neocities website"
			So(d.Long, ShouldEqual, long)
		})
	})
}
