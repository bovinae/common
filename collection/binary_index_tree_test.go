package collection

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewBinaryIndexTree(t *testing.T) {
	Convey("TestNewBinaryIndexTree", t, func() {
		Convey("TestNewBinaryIndexTree", func() {
			a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
			expect := []int{1, 3, 3, 10, 5, 11, 7, 36, 9, 19, 11, 42, 13, 27, 15, 136}
			tree := NewBinaryIndexTree(a)
			So(tree.data, ShouldResemble, expect)
		})
	})
}

func BenchmarkNewBinaryIndexTree(b *testing.B) {
	a := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		a[i] = i + 1
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewBinaryIndexTree(a)
	}
}

func TestPrefixSum(t *testing.T) {
	Convey("TestPrefixSum", t, func() {
		Convey("TestPrefixSum", func() {
			a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
			tree := NewBinaryIndexTree(a)
			So(tree.PrefixSum(11), ShouldEqual, 78)
		})
	})
}

func TestUpdate(t *testing.T) {
	Convey("TestPrefixSum", t, func() {
		Convey("TestPrefixSum", func() {
			a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
			tree := NewBinaryIndexTree(a)
			tree.Update(12, 15)
			So(tree.data[12], ShouldEqual, 15)
			So(tree.data[13], ShouldEqual, 29)
			So(tree.data[15], ShouldEqual, 138)
		})
	})
}
