package util

import (
	"fmt"
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
		Convey("nil", func() {
			So(CompareAny(nil, nil), ShouldEqual, EQUAL)
			So(CompareAny(nil, ""), ShouldEqual, LESS)
			So(CompareAny("", nil), ShouldEqual, GREATER)
			So(CompareAny("", 1), ShouldEqual, LESS)
		})
	})
}

func TestIsEmpty(t *testing.T) {
	Convey("TestIsEmpty", t, func() {
		Convey("TestIsEmpty", func() {
			So(IsEmpty(nil), ShouldEqual, true)
			So(IsEmpty(""), ShouldEqual, true)
			So(IsEmpty(" "), ShouldEqual, true)
			So(IsEmpty("  ab c "), ShouldEqual, false)
		})
	})
}

func TestFindFirstNonEmpty(t *testing.T) {
	values := []any{nil, 123, "456", " ", nil, 12.3}

	Convey("TestFindFirstNonEmpty", t, func() {
		Convey("TestFindFirstNonEmpty", func() {
			pos, element := FindFirstNonEmpty(0, values)
			So(pos, ShouldEqual, 1)
			So(element, ShouldResemble, []byte(fmt.Sprint(123)))

			pos, _ = FindFirstNonEmpty(1, values)
			So(pos, ShouldEqual, 1)

			pos, _ = FindFirstNonEmpty(3, values)
			So(pos, ShouldEqual, 5)
		})
	})
}
