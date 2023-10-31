package collection

import "github.com/bovinae/common/types"

type BinaryIndexTree[T types.Integer] struct {
	data []T
}

func NewBinaryIndexTree[T types.Integer](a []T) *BinaryIndexTree[T] {
	data := make([]T, len(a))
	// index start from 1
	for i := 1; i <= len(a); i++ {
		data[i-1] += a[i-1]
		for j := i - 1; j > i&(i-1); j = j & (j - 1) {
			data[i-1] += data[j-1]
		}
	}
	return &BinaryIndexTree[T]{
		data: data,
	}
}

func (t *BinaryIndexTree[T]) PrefixSum(i int) T {
	// index start from 1
	i++

	var sum T
	for ; i > 0; i = i & (i - 1) {
		sum += t.data[i-1]
	}
	return sum
}

func (t *BinaryIndexTree[T]) Update(i int, target T) {
	// index start from 1
	i++

	delta := target - t.data[i-1]
	for ; i <= len(t.data); i += i - (i & (i - 1)) {
		t.data[i-1] += delta
	}
}
