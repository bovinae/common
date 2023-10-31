package collection

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewSegmentTree(t *testing.T) {
	Convey("TestNewSegmentTree", t, func() {
		Convey("TestNewSegmentTree", func() {
			a := []int{1, 2, 3, 4, 5}
			st := NewSegmentTree(a)
			So(st.tree, ShouldResemble, []int{15, 6, 9, 3, 3, 4, 5, 1, 2, 0, 0, 0, 0, 0, 0, 0})
		})
	})
}

func TestUpdateSegmentTree(t *testing.T) {
	Convey("TestUpdateSegmentTree", t, func() {
		Convey("TestUpdateSegmentTree", func() {
			a := []int{1, 2, 3, 4, 5}
			st := NewSegmentTree(a)
			st.Update(Interval[int]{0, 3}, 1)
			So(st.tree, ShouldResemble, []int{19, 9, 10, 3, 3, 5, 5, 1, 2, 0, 0, 0, 0, 0, 0, 0})
		})
	})
}

func TestQuerySegmentTree(t *testing.T) {
	Convey("TestQuerySegmentTree", t, func() {
		Convey("TestQuerySegmentTree", func() {
			a := []int{1, 2, 3, 4, 5}
			st := NewSegmentTree(a)
			So(st.Query(Interval[int]{0, 3}), ShouldEqual, 10)
		})
	})
}
