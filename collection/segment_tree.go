package collection

import (
	"github.com/bovinae/common/types"
	"github.com/bovinae/common/util"
)

type SegmentTree[T types.Integer] struct {
	data []T
	tree []T
	mark []T
}

func NewSegmentTree[T types.Integer](a []T) *SegmentTree[T] {
	tree := &SegmentTree[T]{
		data: a,
		tree: make([]T, util.NextPow2(2*len(a)-1)),
		mark: make([]T, util.NextPow2(2*len(a)-1)),
	}
	tree.buildTree(a, 0)
	return tree
}

// index p start from 0
func (t *SegmentTree[T]) buildTree(a []T, p int) {
	if len(a) == 1 {
		t.tree[p] = a[0]
		return
	}

	mid := (len(a)-1)/2 + 1
	t.buildTree(a[:mid], 2*p+1)
	t.buildTree(a[mid:], 2*p+2)
	t.tree[p] = t.tree[2*p+1] + t.tree[2*p+2]
}

func (t *SegmentTree[T]) Update(target Interval[int], delta T) {
	t.update(target, delta, 0, Interval[int]{0, len(t.data) - 1})
}

func (t *SegmentTree[T]) update(target Interval[int], delta T, p int, curr Interval[int]) {
	// no intersection
	if curr.Right < target.Left || curr.Left > target.Right {
		return
	}

	// target interval contains current interval
	// target  interval: <------->
	// current interval:   <--->
	if target.Left <= curr.Left && target.Right >= curr.Right {
		t.tree[p] += delta * T(curr.Right-curr.Left+1)
		// have child tree need to update, but we use mark to lazy update child tree
		if curr.Right > curr.Left {
			t.mark[p] += delta
		}
		return
	}

	// partial intersection
	// target  interval:    <--->
	// current interval:  <------->
	// current interval:      <----->
	// current interval: <---->
	t.pushDown(p, curr.Right-curr.Left+1)
	mid := (curr.Left + curr.Right) / 2
	t.update(target, delta, 2*p+1, Interval[int]{curr.Left, mid})
	t.update(target, delta, 2*p+2, Interval[int]{mid + 1, curr.Right})
	t.tree[p] = t.tree[2*p+1] + t.tree[2*p+2]
}

func (t *SegmentTree[T]) pushDown(p, currLen int) {
	// update child tree according to current mark
	t.tree[2*p+1] += t.mark[p] * T(currLen-currLen/2)
	t.tree[2*p+2] += t.mark[p] * T(currLen/2)

	// update child tree's mark
	t.mark[2*p+1] += t.mark[p]
	t.mark[2*p+2] += t.mark[p]

	// clear current mark
	t.mark[p] = 0
}

func (t *SegmentTree[T]) Query(target Interval[int]) T {
	return t.query(target, 0, Interval[int]{0, len(t.data) - 1})
}

func (t *SegmentTree[T]) query(target Interval[int], p int, curr Interval[int]) T {
	// no intersection
	if curr.Right < target.Left || curr.Left > target.Right {
		return 0
	}

	// target interval contains current interval
	if target.Left <= curr.Left && target.Right >= curr.Right {
		return t.tree[p]
	}

	// partial intersection
	t.pushDown(p, curr.Right-curr.Left+1)
	mid := (curr.Left + curr.Right) / 2
	return t.query(target, 2*p+1, Interval[int]{curr.Left, mid}) + t.query(target, 2*p+2, Interval[int]{mid + 1, curr.Right})
}
