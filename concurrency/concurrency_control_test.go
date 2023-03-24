package concurrency

import (
	"fmt"
	"sort"
	"sync/atomic"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConcurrencyControl(t *testing.T) {
	total := 2000
	minTokenSize, maxTokenSize := int64(total), int64(0)
	result := make([]int, 0, total)
	expected := make([]int, 0, total)
	ch := make(chan int, total)
	cc := NewConcurrencyControl(1)
	expectedGoroutine := len(cc.Ch)
	for i := 0; i < total; i++ {
		i := i
		expected = append(expected, i)
		cc.Get()
		Go(func() {
			defer cc.Put()
			getMinMaxTokenSize(cc, &minTokenSize, &maxTokenSize)
			ch <- i
		})
	}
	cc.Get()
	Go(func() {
		defer cc.Put()
		for i := 0; i < total; i++ {
			getMinMaxTokenSize(cc, &minTokenSize, &maxTokenSize)
			result = append(result, <-ch)
		}
	})
	getMinMaxTokenSize(cc, &minTokenSize, &maxTokenSize)
	cc.Wait()
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	Convey("TestConcurrencyControl", t, func() {
		Convey("TestConcurrencyControl", func() {
			So(result, ShouldResemble, expected)
			fmt.Println(minTokenSize, maxTokenSize)
			So(minTokenSize, ShouldBeGreaterThanOrEqualTo, 0)
			So(maxTokenSize, ShouldBeLessThanOrEqualTo, expectedGoroutine)
		})
	})
}

func getMinMaxTokenSize(cc *ConcurrencyControl, minTokenSize, maxTokenSize *int64) {
	currGoroutine := int64(len(cc.Ch))
	tmp := atomic.LoadInt64(minTokenSize)
	if currGoroutine < tmp {
		atomic.CompareAndSwapInt64(minTokenSize, tmp, currGoroutine)
	}
	tmp = atomic.LoadInt64(maxTokenSize)
	if currGoroutine > tmp {
		atomic.CompareAndSwapInt64(maxTokenSize, tmp, currGoroutine)
	}
}
