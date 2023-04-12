package algorithm

import (
	"testing"

	"github.com/bovinae/common/util"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBinarySearch(t *testing.T) {
	ids := []int{1, 3, 5, 7, 9}
	compare := func(i int, target any) int {
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
			pos := BinarySearch(ids, 7, compare)
			So(pos, ShouldEqual, 3)
			pos = BinarySearch(ids, 0, compare)
			So(pos, ShouldEqual, -1)
			pos = BinarySearch(ids, 8, compare)
			So(pos, ShouldEqual, -1)
			pos = BinarySearch(ids, 10, compare)
			So(pos, ShouldEqual, -1)
		})
	})
}
