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
			// 5 valid entries, 2 invalid
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
		C("find", func() {
			strs := []string{
				"v0.10.20",
				"v0.10.21",
				"v0.10.9",
				"v0.11.7",
				"v0.11.6",
				"v0.11.2",
			}
			C("newest version", func() {
				vl := NewVersionListFromSlice(strs)
				bestMatch, err := vl.FindNewest("0.11")
				So(err, ShouldBeNil)
				So(bestMatch, ShouldEqual, "v0.11.7")
			})
			C("exact version", func() {
				vl := NewVersionListFromSlice(strs)
				bestMatch, err := vl.FindExact("0.11.2")
				So(err, ShouldBeNil)
				So(bestMatch, ShouldEqual, "v0.11.2")
			})
			C("panic FindExact when version is not full", func() {
				vl := NewVersionListFromSlice(strs)
				_, err := vl.FindExact("0.11")
				So(err, ShouldEqual, ErrorFullVersionMustBeSpecified)
			})
		})
	})
}
