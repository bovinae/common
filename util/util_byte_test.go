package util

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetPrefix(t *testing.T) {
	Convey("TestGetPrefix", t, func() {
		Convey("chinese", func() {
			So(GetPrefix([]byte("找第一个1非空元素"), []byte("找第一个123")), ShouldResemble, []byte("找第一个1"))
		})
		Convey("english", func() {
			So(GetPrefix([]byte("english"), []byte("engilish")), ShouldResemble, []byte("eng"))
		})
	})
}

func TestGetSuffix(t *testing.T) {
	Convey("TestGetSuffix", t, func() {
		Convey("chinese", func() {
			So(GetSuffix([]byte("找第一个1非空元素"), []byte("找第一个123")), ShouldResemble, []byte(""))
		})
		Convey("english", func() {
			So(GetSuffix([]byte("english"), []byte("engilish")), ShouldResemble, []byte("lish"))
		})
	})
}
