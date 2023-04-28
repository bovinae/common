package collection

import (
	"github.com/bovinae/common/algorithm"
	"github.com/bovinae/common/util"
)

type integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// Left closed and right open: [a, b).
type Interval[T integer] struct {
	Left  T
	Right T
}

// Sorted in ascending order by left endpoint and without intersections.
type SortedInterval[T integer] []Interval[T]

func NewIntervalSlice[T integer]() *SortedInterval[T] {
	return &SortedInterval[T]{}
}

// Append point must larger than all interval.
func (si *SortedInterval[T]) AppendPoint(point ...T) {
	for _, each := range point {
		si.appendPoint(each)
	}
}

func (si *SortedInterval[T]) appendPoint(point T) {
	length := len(*si)
	if length == 0 || point > (*si)[length-1].Right {
		*si = append(*si, Interval[T]{point, point + 1})
		return
	}
	(*si)[length-1].Right++
}

// Append interval must larger than all interval.
func (si *SortedInterval[T]) AppendInterval(interval ...Interval[T]) {
	for _, each := range interval {
		si.appendInterval(each)
	}
}

func (si *SortedInterval[T]) appendInterval(interval Interval[T]) {
	length := len(*si)
	if length == 0 || interval.Left > (*si)[length-1].Right {
		*si = append(*si, interval)
		return
	}
	(*si)[length-1].Right = interval.Right
}

func (si SortedInterval[T]) BinarySearch(left T) int {
	return algorithm.LowerBound(si, left, func(i int) int {
		if left < si[i].Left {
			return util.LESS
		}
		if left > si[i].Left {
			return util.GREATER
		}
		return util.EQUAL
	})
}

func (si SortedInterval[T]) IntersectionElementNum(si1 SortedInterval[T]) T {
	if len(si) == 0 || len(si1) == 0 {
		return 0
	}
	pos := si.BinarySearch(si1[0].Left)
	if pos < 0 {
		pos = 0
	}
	i, j := pos, 0
	var count T
	for i < len(si) && j < len(si1) {
		if si[i].Right <= si1[j].Left {
			i++
			continue
		}
		if si[i].Left >= si1[j].Right {
			j++
			continue
		}
		leftMax := util.Max(si[i].Left, si1[j].Left).(T)
		rightMin := util.Min(si[i].Right, si1[j].Right).(T)
		count += rightMin - leftMax
	}
	return count
}
