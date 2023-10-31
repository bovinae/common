package util

import (
	"math"
	"unsafe"

	"github.com/bovinae/common/types"
)

func Min(a, b any) any {
	if CompareAny(a, b) == LESS {
		return a
	}
	return b
}

func Max(a, b any) any {
	if CompareAny(a, b) == GREATER {
		return a
	}
	return b
}

func CalcAverage[T float32 | float64](values []T) T {
	if len(values) == 0 {
		return 0
	}

	sum := T(0.0)
	for i := 0; i < len(values); i++ {
		sum += values[i]
	}
	return sum / T(len(values))
}

func CalcSigma[T float32 | float64](average T, values []T) T {
	if len(values) == 0 {
		return 0
	}

	squareSum := T(0.0)
	for i := 0; i < len(values); i++ {
		tmp := values[i] - average
		tmp *= tmp
		squareSum += tmp
	}
	return T(math.Sqrt(float64(squareSum) / float64(len(values))))
}

func NextPow2[T types.Integer](num T) T {
	if num&(num-1) == 0 {
		return num
	}

	bitSize := 8 * unsafe.Sizeof(num)
	cursor := T(1 << (bitSize - 1))
	for cursor != 0 {
		if cursor&num != 0 {
			return cursor << 1
		}
		oldCursor := cursor
		cursor >>= 1
		if cursor < 0 {
			cursor ^= oldCursor
		}
	}
	return cursor << 1
}
