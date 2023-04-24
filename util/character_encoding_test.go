package util

import (
	"testing"

	"github.com/bovinae/golang.org/x/text/encoding/simplifiedchinese"
	"github.com/bovinae/golang.org/x/text/transform"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUTF8GBK(t *testing.T) {
	GB18030 := simplifiedchinese.All[0]
	encodeReader := transform.NewReader(nil, GB18030.NewEncoder())
	decodeReader := transform.NewReader(nil, GB18030.NewDecoder())
	Convey("TestUTF8GBK", t, func() {
		Convey("chinese", func() {
			src := "中国 人"
			gbkBytes, err := UTF82GBK(encodeReader, src)
			So(err, ShouldEqual, nil)
			src1, err := GBK2UTF8(decodeReader, gbkBytes)
			So(err, ShouldEqual, nil)
			So(src1, ShouldEqual, src)
		})
		Convey("english", func() {
			src := "hello english"
			gbkBytes, err := UTF82GBK(encodeReader, src)
			So(err, ShouldEqual, nil)
			src1, err := GBK2UTF8(decodeReader, gbkBytes)
			So(err, ShouldEqual, nil)
			So(src1, ShouldEqual, src)
		})
		Convey("number", func() {
			src := "1234 5"
			gbkBytes, err := UTF82GBK(encodeReader, src)
			So(err, ShouldEqual, nil)
			src1, err := GBK2UTF8(decodeReader, gbkBytes)
			So(err, ShouldEqual, nil)
			So(src1, ShouldEqual, src)
		})
		Convey("empty", func() {
			src := ""
			gbkBytes, err := UTF82GBK(encodeReader, src)
			So(err, ShouldEqual, nil)
			src1, err := GBK2UTF8(decodeReader, gbkBytes)
			So(err, ShouldEqual, nil)
			So(src1, ShouldEqual, src)
		})
	})
}
