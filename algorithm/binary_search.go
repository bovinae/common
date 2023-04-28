package algorithm

import (
	"reflect"

	"github.com/bovinae/common/util"
)

const (
	NOT_FOUND = -2
)

func BinarySearch(sortedSlice any, target any, compare func(i int) int) int {
	if reflect.TypeOf(sortedSlice).Kind() != reflect.Slice {
		return NOT_FOUND
	}
	left, right := 0, reflect.ValueOf(sortedSlice).Len()-1
	for left <= right {
		mid := left + (right-left)>>1
		cmp := compare(mid)
		if cmp == util.EQUAL {
			return mid
		}
		if cmp == util.LESS {
			right = mid - 1
			continue
		}
		left = mid + 1
	}
	return NOT_FOUND
}

// -1: less than all
func LowerBound(sortedSlice any, target any, compare func(i int) int) int {
	if reflect.TypeOf(sortedSlice).Kind() != reflect.Slice {
		return NOT_FOUND
	}
	left, right := 0, reflect.ValueOf(sortedSlice).Len()-1
	for left <= right {
		mid := left + (right-left)>>1
		cmp := compare(mid)
		if cmp == util.EQUAL {
			return mid
		}
		if cmp == util.LESS {
			right = mid - 1
			continue
		}
		left = mid + 1
	}
	return right
}
