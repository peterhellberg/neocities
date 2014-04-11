package client

import (
	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func TestArgs(t *testing.T) {
	Convey("Args", t, func() {
		cmd := "foo"
		params := []string{"bar", "baz"}

		Convey("should contain a Command and Params", func() {
			args := &Args{Command: cmd, Params: params}

			So(args.Command, ShouldEqual, cmd)
			So(args.Params, ShouldResemble, params)
		})

		Convey("newArgs", func() {
			Convey("with no arguments", func() {
				args := newArgs([]string{})

				So(args.Command, ShouldEqual, "")
				So(args.ParamsSize(), ShouldEqual, 0)
			})

			Convey("with command ", func() {
				args := newArgs([]string{cmd})

				So(args.Command, ShouldEqual, cmd)
				So(args.ParamsSize(), ShouldEqual, 0)
			})

			Convey("with command and params", func() {
				args := newArgs([]string{cmd, "param1", "param2"})

				So(args.Command, ShouldEqual, cmd)
				So(args.ParamsSize(), ShouldEqual, 2)
			})
		})

		Convey("IsParamsEmpty", func() {
			Convey("with no params", func() {
				args := newArgs([]string{cmd})
				So(args.IsParamsEmpty(), ShouldBeTrue)
			})

			Convey("with params", func() {
				args := newArgs([]string{cmd, "bar", "baz"})
				So(args.IsParamsEmpty(), ShouldBeFalse)
			})
		})
	})
}
