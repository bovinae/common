package algorithm

import "reflect"

const (
	NOT_FOUND = -1
)

const (
	LESS = iota - 1
	EQUAL
	GREATER
)

func BinarySearch(sortedSlice any, target any, compare func(i int, target any) int) int {
	if reflect.TypeOf(sortedSlice).Kind() != reflect.Slice {
		return NOT_FOUND
	}
	left, right := 0, reflect.ValueOf(sortedSlice).Len()-1
	for left <= right {
		mid := left + (right-left)>>1
		cmp := compare(mid, target)
		if cmp == EQUAL {
			return mid
		}
		if cmp == LESS {
			right = mid - 1
			continue
		}
		left = mid + 1
	}
	return NOT_FOUND
}
