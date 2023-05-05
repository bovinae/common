package util

import "math"

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
