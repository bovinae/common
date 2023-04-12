package util

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCompareAny(t *testing.T) {
	Convey("TestCompareAny", t, func() {
		Convey("number", func() {
			So(CompareAny(1, 7), ShouldEqual, LESS)
			So(CompareAny(7, 7), ShouldEqual, EQUAL)
			So(CompareAny(9, 7), ShouldEqual, GREATER)
		})
		Convey("string", func() {
			So(CompareAny("1", "7"), ShouldEqual, LESS)
			So(CompareAny("7", "7"), ShouldEqual, EQUAL)
			So(CompareAny("9", "7"), ShouldEqual, GREATER)
		})
		Convey("bool", func() {
			So(CompareAny(false, true), ShouldEqual, LESS)
			So(CompareAny(true, true), ShouldEqual, EQUAL)
			So(CompareAny(true, false), ShouldEqual, GREATER)
		})
	})
}
