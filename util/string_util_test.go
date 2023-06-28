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

func TestContainsSubSequence(t *testing.T) {
	Convey("sub string", t, func() {
		Convey("prefix substr", func() {
			So(ContainsSubSequence("a时间2戳转ab换工10具f", "a时"), ShouldEqual, true)
		})
		Convey("suffix substr", func() {
			So(ContainsSubSequence("a时间2戳转ab换工10具f", "10具f"), ShouldEqual, true)
		})
		Convey("middle substr", func() {
			So(ContainsSubSequence("a时间2戳转ab换工10具f", "戳转"), ShouldEqual, true)
		})
	})
	Convey("sub sequence", t, func() {
		Convey("prefix subseq", func() {
			So(ContainsSubSequence("a时间2戳转ab换工10具f", "a时2"), ShouldEqual, true)
		})
		Convey("suffix subseq", func() {
			So(ContainsSubSequence("a时间2戳转ab换工10具f", "2换f"), ShouldEqual, true)
		})
		Convey("middle subseq", func() {
			So(ContainsSubSequence("a时间2戳转ab换工10具f", "时转工"), ShouldEqual, true)
		})
	})
	Convey("reverse sub sequence", t, func() {
		Convey("middle subseq", func() {
			So(ContainsSubSequence("a时间2戳转ab换工10具f", "转时"), ShouldEqual, false)
		})
	})
}
