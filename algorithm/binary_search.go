package algorithm

import (
	"bytes"
	"math"
	"reflect"
	"time"
)

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

func CompareAny(value1, value2 any) int {
	if value1 == nil {
		if value2 != nil {
			return LESS
		}
		return EQUAL
	}
	if value2 == nil {
		return GREATER
	}
	switch value1 := value1.(type) {
	case []byte:
		value2, ok := value2.([]byte)
		if !ok {
			return LESS
		}
		cmp := bytes.Compare(value1, value2)
		if cmp < 0 {
			return LESS
		}
		if cmp == 0 {
			return EQUAL
		}
		return GREATER
	case string:
		value2, ok := value2.(string)
		if !ok {
			return LESS
		}
		if value1 < value2 {
			return LESS
		}
		if value1 == value2 {
			return EQUAL
		}
		return GREATER
	case int64:
		value2, ok := value2.(int64)
		if !ok {
			return LESS
		}
		if value1 < value2 {
			return LESS
		}
		if value1 == value2 {
			return EQUAL
		}
		return GREATER
	case int:
		value2, ok := value2.(int)
		if !ok {
			return LESS
		}
		if value1 < value2 {
			return LESS
		}
		if value1 == value2 {
			return EQUAL
		}
		return GREATER
	case float64:
		value2, ok := value2.(float64)
		if !ok {
			return LESS
		}
		if value1 < value2 {
			return LESS
		}
		if math.Abs(value1-value2) < 0.000001 {
			return EQUAL
		}
		return GREATER
	case time.Time:
		value2, ok := value2.(time.Time)
		if !ok {
			return LESS
		}
		if value1.Before(value2) {
			return LESS
		}
		if value1.Equal(value2) {
			return EQUAL
		}
		return GREATER
	}
	return LESS
}
