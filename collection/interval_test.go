package collection

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAppendPoint(t *testing.T) {
	si := NewSortedInterval[int]()
	si.AppendPoint(1, 2, 3, 10)
	si.dump()
	Convey("TestAppendPoint", t, func() {
		Convey("TestAppendPoint", func() {
			So(si[0].Left, ShouldEqual, 1)
			So(si[0].Right, ShouldEqual, 4)
			So(si[1].Left, ShouldEqual, 10)
			So(si[1].Right, ShouldEqual, 11)
		})
	})
}

func TestAppendInterval(t *testing.T) {
	si := NewSortedInterval[int]()
	si.AppendInterval(Interval[int]{1, 4}, Interval[int]{10, 11})
	si.dump()
	Convey("TestAppendInterval", t, func() {
		Convey("TestAppendInterval", func() {
			So(si[0].Left, ShouldEqual, 1)
			So(si[0].Right, ShouldEqual, 4)
			So(si[1].Left, ShouldEqual, 10)
			So(si[1].Right, ShouldEqual, 11)
		})
	})
}

func TestLowerBound(t *testing.T) {
	si := NewSortedInterval[int]()
	si.AppendInterval(Interval[int]{1, 4}, Interval[int]{10, 11})
	si.dump()
	Convey("TestAppendInterval", t, func() {
		Convey("TestAppendInterval", func() {
			So(si.LowerBound(0), ShouldEqual, -1)
			So(si.LowerBound(1), ShouldEqual, 0)
			So(si.LowerBound(5), ShouldEqual, 0)
			So(si.LowerBound(10), ShouldEqual, 1)
			So(si.LowerBound(15), ShouldEqual, 1)
		})
	})
}

func TestIntersectionElementNum(t *testing.T) {
	si := NewSortedInterval[int]()
	si.AppendInterval(Interval[int]{1, 4}, Interval[int]{10, 60}, Interval[int]{100, 1000})
	si.dump()
	fmt.Println("")
	si1 := NewSortedInterval[int]()
	si1.AppendInterval(Interval[int]{10, 80}, Interval[int]{90, 600})
	si1.dump()
	Convey("TestIntersectionElementNum", t, func() {
		Convey("TestIntersectionElementNum", func() {
			So(si.IntersectionElementNum(si1), ShouldEqual, 550)
		})
	})
}
