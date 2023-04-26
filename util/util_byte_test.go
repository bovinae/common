package util

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetPrefix(t *testing.T) {
	Convey("TestGetPrefix", t, func() {
		Convey("chinese", func() {
			So(GetPrefix([]rune("找第一个1非空元素"), []rune("找第一个123")), ShouldResemble, []rune("找第一个1"))
		})
		Convey("english", func() {
			So(GetPrefix([]rune("english"), []rune("engilish")), ShouldResemble, []rune("eng"))
		})
	})
}

func TestGetSuffix(t *testing.T) {
	Convey("TestGetSuffix", t, func() {
		Convey("chinese", func() {
			So(GetSuffix([]rune("找第一个1非空元素"), []rune("找第一个123")), ShouldResemble, []rune(""))
		})
		Convey("english", func() {
			So(GetSuffix([]rune("english"), []rune("engilish")), ShouldResemble, []rune("lish"))
		})
	})
}

func TestReverseRuneSlice(t *testing.T) {
	Convey("TestReverseRuneSlice", t, func() {
		Convey("one char", func() {
			rs := []rune("元")
			ReverseSlice(rs)
			So(string(rs), ShouldEqual, "元")
		})
		Convey("two char", func() {
			rs := []rune("元1")
			ReverseSlice(rs)
			So(string(rs), ShouldEqual, "1元")
		})
		Convey("multi char", func() {
			rs := []rune("1找s元素3")
			ReverseSlice(rs)
			So(string(rs), ShouldEqual, "3素元s找1")
		})
	})
}
