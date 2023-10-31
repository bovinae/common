package collection

import (
	"fmt"

	"github.com/bovinae/common/algorithm"
	"github.com/bovinae/common/types"
	"github.com/bovinae/common/util"
)

// Left closed and right open: [a, b).
type Interval[T types.Integer] struct {
	Left  T
	Right T
}

func (i Interval[T]) String() string {
	return fmt.Sprintf("[%v, %v)", i.Left, i.Right)
}

// Sorted in ascending order by left endpoint and without intersections.
type SortedInterval[T types.Integer] struct {
	Intervals  []Interval[T]
	TotalPoint int
}

func NewSortedInterval[T types.Integer]() *SortedInterval[T] {
	return &SortedInterval[T]{}
}

// Append point must larger than all interval.
func (si *SortedInterval[T]) AppendPoint(point ...T) {
	for _, each := range point {
		si.appendPoint(each)
	}
}

func (si *SortedInterval[T]) appendPoint(point T) {
	si.TotalPoint++
	length := len(si.Intervals)
	if length == 0 || point > si.Intervals[length-1].Right {
		si.Intervals = append(si.Intervals, Interval[T]{point, point + 1})
		return
	}
	si.Intervals[length-1].Right++
}

// Append interval must larger than all interval.
func (si *SortedInterval[T]) AppendInterval(interval ...Interval[T]) {
	for _, each := range interval {
		si.appendInterval(each)
	}
}

func (si *SortedInterval[T]) appendInterval(interval Interval[T]) {
	si.TotalPoint += int(interval.Right - interval.Left)
	length := len(si.Intervals)
	if length == 0 || interval.Left > si.Intervals[length-1].Right {
		si.Intervals = append(si.Intervals, interval)
		return
	}
	si.Intervals[length-1].Right = interval.Right
}

func (si *SortedInterval[T]) GetTotalPoint() int {
	return si.TotalPoint
}

// return lower bound position
func (si *SortedInterval[T]) LowerBound(left T) int {
	return algorithm.LowerBound(si.Intervals, left, func(i int) int {
		if left < si.Intervals[i].Left {
			return util.LESS
		}
		if left > si.Intervals[i].Left {
			return util.GREATER
		}
		return util.EQUAL
	})
}

func (si *SortedInterval[T]) IntersectionElementNum(si1 *SortedInterval[T]) T {
	if len(si.Intervals) == 0 || len(si1.Intervals) == 0 {
		return 0
	}
	pos := si.LowerBound(si1.Intervals[0].Left)
	if pos < 0 {
		pos = 0
	}
	i, j := pos, 0
	var count T
	for i < len(si.Intervals) && j < len(si1.Intervals) {
		if si.Intervals[i].Right <= si1.Intervals[j].Left {
			i++
			continue
		}
		if si.Intervals[i].Left >= si1.Intervals[j].Right {
			j++
			continue
		}
		leftMax := util.Max(si.Intervals[i].Left, si1.Intervals[j].Left).(T)
		rightMin := util.Min(si.Intervals[i].Right, si1.Intervals[j].Right).(T)
		count += rightMin - leftMax
		i++
		j++
	}
	return count
}

func (si *SortedInterval[T]) dump() {
	for i := 0; i < len(si.Intervals); i++ {
		fmt.Println(si.Intervals[i].String())
	}
}
