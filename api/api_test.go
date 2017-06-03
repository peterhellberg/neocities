package api

import (
	"io/ioutil"

	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func TestAPI(t *testing.T) {
	Convey("API", t, func() {
		cred := &Credentials{"foo", "bar"}

		Convey("expected constants", func() {
			So(apiURL, ShouldEqual, "https://neocities.org/api/")
			So(userAgent, ShouldEqual, "neocities (Go 1.4.2 package http)")
		})

		Convey("newInfoRequest", func() {
			req, err := newInfoRequest(cred, "foo")

			So(err, ShouldBeNil)
			So(req.Method, ShouldEqual, "GET")
			So(req.URL.String(), ShouldEqual, "https://neocities.org/api/info?sitename=foo")
		})

		Convey("newDeleteRequest", func() {
			req, err := newDeleteRequest(cred, []string{"foo", "bar"})

			So(err, ShouldBeNil)
			So(req.Method, ShouldEqual, "POST")
			So(req.URL.String(), ShouldEqual, "https://neocities.org/api/delete")
		})

		Convey("newUploadRequest", func() {
			req, err := newUploadRequest(cred, []string{"../LICENSE"})

			So(err, ShouldBeNil)
			So(req.Method, ShouldEqual, "POST")
			So(req.URL.String(), ShouldEqual, "https://neocities.org/api/upload")
		})

		Convey("newUploadDataRequest", func() {
			testContent, err := ioutil.ReadFile("../LICENSE")
			testData := []UploadData{UploadData{FileName: "LICENSE_string", Content: testContent}}
			req, err := newUploadDataRequest(cred, testData)

			So(err, ShouldBeNil)
			So(req.Method, ShouldEqual, "POST")
			So(req.URL.String(), ShouldEqual, "https://neocities.org/api/upload")
		})
	})
}
