package collection

import "testing"

func TestList(t *testing.T) {
	l := NewList()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)
	l.dump()
}
