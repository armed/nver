package util

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUtil(t *testing.T) {
	Convey("archs map should correctly convert GOARH", t, func() {
		So(archs["amd64"], ShouldEqual, "x64")
		So(archs["386"], ShouldEqual, "x86")
	})
	Convey("makeUrl should create url from version, os and arch", t, func() {
		url := makeUrl("v0.10.21", "darwin", "amd64")
		So(url, ShouldEqual, "http://nodejs.org/dist/v0.10.21/node-v0.10.21-darwin-x64.tar.gz")
	})
}
