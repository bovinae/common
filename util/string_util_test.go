package util

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReverseString(t *testing.T) {
	Convey("TestReverseString", t, func() {
		Convey("one char", func() {
			So(ReverseString("1"), ShouldEqual, "1")
		})
		Convey("two char", func() {
			So(ReverseString("12"), ShouldEqual, "21")
		})
		Convey("multi char", func() {
			So(ReverseString("123"), ShouldEqual, "321")
		})
	})
}
