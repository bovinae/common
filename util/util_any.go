package util

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"time"
)

const (
	LESS = iota - 1
	EQUAL
	GREATER
)

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
	case bool:
		value2, ok := value2.(bool)
		if !ok {
			return LESS
		}
		if !value1 && value2 {
			return LESS
		}
		if value1 == value2 {
			return EQUAL
		}
		return GREATER
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

func IsEmpty(value any) bool {
	if value == nil {
		return true
	}
	if val, ok := value.(string); ok && len(strings.TrimSpace(val)) == 0 {
		return true
	}
	return false
}

func FindFirstNonEmpty(pos int, values []any) (int, []byte) {
	var curr []byte
	for ; pos < len(values); pos++ {
		if IsEmpty(values[pos]) {
			continue
		}
		curr = []byte(fmt.Sprint(values[pos]))
		break
	}
	return pos, curr
}
