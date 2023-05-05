package collection

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewBitMapByIds(t *testing.T) {
	Convey("TestNewBitMapByIds", t, func() {
		Convey("TestNewBitMapByIds", func() {
			bm := NewBitMapByIds([]int64{1, 3})
			So(bm, ShouldResemble, BitMap{0b1010})
		})
	})
}

func TestSetAndIsExist(t *testing.T) {
	bm := NewBitMap(8)
	Convey("TestSetAndIsExist", t, func() {
		Convey("TestSetAndIsExist", func() {
			bm.Set(7)
			exist := bm.IsExist(7)
			So(exist, ShouldEqual, true)
			bm.Set(10000)
			exist = bm.IsExist(10000)
			So(exist, ShouldEqual, true)
		})
	})
}

func TestCleanBitmap(t *testing.T) {
	bm := NewBitMap(8)
	Convey("TestCleanBitmap", t, func() {
		Convey("TestCleanBitmap", func() {
			bm.Set(6)
			bm.Set(7)
			exist := bm.IsExist(7)
			So(exist, ShouldEqual, true)
			bm.Clean(7)
			exist = bm.IsExist(7)
			So(exist, ShouldEqual, false)
			exist = bm.IsExist(6)
			So(exist, ShouldEqual, true)
		})
	})
}
