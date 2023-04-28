package algorithm

import (
	"testing"

	"github.com/bovinae/common/util"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBinarySearch(t *testing.T) {
	ids := []int{1, 3, 5, 7, 9}
	var target any
	compare := func(i int) int {
		if target.(int) < ids[i] {
			return util.LESS
		}
		if target == ids[i] {
			return util.EQUAL
		}
		return util.GREATER
	}
	Convey("TestBinarySearch", t, func() {
		Convey("TestBinarySearch", func() {
			target = 7
			pos := BinarySearch(ids, target, compare)
			So(pos, ShouldEqual, 3)
			target = 0
			pos = BinarySearch(ids, target, compare)
			So(pos, ShouldEqual, -1)
			target = 8
			pos = BinarySearch(ids, target, compare)
			So(pos, ShouldEqual, -1)
			target = 10
			pos = BinarySearch(ids, target, compare)
			So(pos, ShouldEqual, -1)
		})
	})
}

func TestLowerBound(t *testing.T) {
	ids := []int{1, 3, 5, 7, 9}
	var target any
	compare := func(i int) int {
		if target.(int) < ids[i] {
			return util.LESS
		}
		if target == ids[i] {
			return util.EQUAL
		}
		return util.GREATER
	}
	Convey("TestLowerBound", t, func() {
		Convey("total number is odd", func() {
			target = 0
			pos := LowerBound(ids, target, compare)
			So(pos, ShouldEqual, -1)
			target = 2
			pos = LowerBound(ids, target, compare)
			So(pos, ShouldEqual, 0)
			target = 5
			pos = LowerBound(ids, target, compare)
			So(pos, ShouldEqual, 2)
			target = 10
			pos = LowerBound(ids, target, compare)
			So(pos, ShouldEqual, 4)
		})
		Convey("total number is even", func() {
			ids = []int{1, 3, 5, 7, 9, 11}
			target = 0
			pos := LowerBound(ids, target, compare)
			So(pos, ShouldEqual, -1)
			target = 5
			pos = LowerBound(ids, target, compare)
			So(pos, ShouldEqual, 2)
			target = 6
			pos = LowerBound(ids, target, compare)
			So(pos, ShouldEqual, 2)
			target = 12
			pos = LowerBound(ids, target, compare)
			So(pos, ShouldEqual, 5)
		})
	})
}
