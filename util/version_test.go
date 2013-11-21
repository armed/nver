package util

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var C = Convey

func TestVersionList(t *testing.T) {
	C("VersionList should", t, func() {
		C("be created empty with func NewVersionList", func() {
			vl := NewVersionList()
			So(vl, ShouldNotBeNil)
		})
		C("have Add func with validator", func() {
			vl := NewVersionList()
			wrongStr := "some wrong string"
			vl.Add(wrongStr)
			So(vl.Count(), ShouldEqual, 0)

			goodStr := "v0.10.22"
			vl.Add(goodStr)
			So(vl.Count(), ShouldEqual, 1)
			So(vl.Vers()[0], ShouldEqual, "v0.10.22")

			C("and ignore npm versions", func() {
				vl := NewVersionList()
				goodStr := "v0.10.22 1.34.3"
				vl.Add(goodStr)
				So(vl.Count(), ShouldEqual, 1)
				So(vl.Vers()[0], ShouldEqual, "v0.10.22")
			})
		})
		C("be created nonempty from []string with func NewVersionListFromSlice", func() {
			// 5 valid versions, 2 invalid
			strs := []string{
				"v0.10.16",
				"v0.10.17",
				"v0.11.7",
				"v0.11.6",
				"v0.11.8 3.44.32-beta-5",
				"",
				".someFolder",
			}
			vl := NewVersionListFromSlice(strs)
			So(vl.Count(), ShouldEqual, 5)
			So(vl.Vers()[4], ShouldEqual, "v0.11.8")
		})
		C("find best matched version from it's list", func() {
			strs := []string{
				"v0.10.20",
				"v0.10.21",
				"v0.10.9",
				"v0.11.7",
				"v0.11.6",
				"v0.11.2",
			}
			vl := NewVersionListFromSlice(strs)
			success, bestMatch := vl.FindBest("0.11")
			So(success, ShouldBeTrue)
			So(bestMatch, ShouldEqual, "v0.11.7")

			success, bestMatch = vl.FindBest("0.10")
			So(success, ShouldBeTrue)
			So(bestMatch, ShouldEqual, "v0.10.21")
		})

	})
}
