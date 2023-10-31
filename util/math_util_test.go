package util

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCalcAverage(t *testing.T) {
	Convey("TestCalcAverage", t, func() {
		Convey("empty slice", func() {
			So(CalcAverage[float32](nil), ShouldEqual, 0)
		})
		Convey("one element", func() {
			values := []float64{1.2}
			So(CalcAverage(values), ShouldEqual, 1.2)
		})
		Convey("some elements", func() {
			values := []float64{1.2, 3.2, 1.6}
			So(CalcAverage(values), ShouldEqual, 2)
		})
	})
}

func TestCalcSigma(t *testing.T) {
	Convey("TestCalcSigma", t, func() {
		Convey("empty slice", func() {
			So(CalcSigma[float32](0, nil), ShouldEqual, 0)
		})
		Convey("one element", func() {
			values := []float64{1.2}
			So(CalcSigma(values[0], values), ShouldEqual, 0)
		})
		Convey("some elements", func() {
			values := []float64{1.2, 3.2, 1.6}
			So(CalcSigma(2, values), ShouldEqual, 0.8640987597877147)
		})
	})
}

func TestNextPow2(t *testing.T) {
	Convey("TestNextPow2", t, func() {
		Convey("0", func() {
			So(NextPow2(0), ShouldEqual, 0)
		})
		Convey("1", func() {
			So(NextPow2(1), ShouldEqual, 1)
		})
		Convey("2", func() {
			So(NextPow2(2), ShouldEqual, 2)
		})
		Convey("3", func() {
			So(NextPow2(3), ShouldEqual, 4)
		})
		Convey("78", func() {
			So(NextPow2(78), ShouldEqual, 128)
		})
	})
}
