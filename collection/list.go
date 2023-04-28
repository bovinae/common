package collection

import "fmt"

// Reference Golang Doubly Linked List

// Element is an element of a linked list.
type Element struct {
	next *Element

	// The value stored with this element.
	Value any
}

// Next returns the next list element or nil.
func (e *Element) Next() *Element {
	return e.next
}

// List represents a single linked list.
// The zero value for List is an empty list ready to use.
type List struct {
	head *Element
	tail *Element
	len  int // current list length excluding (this) sentinel element
}

func NewList() *List {
	list := &List{
		head: &Element{},
	}
	list.tail = list.head
	return list
}

// Len returns the number of elements of list l.
// The complexity is O(1).
func (l *List) Len() int { return l.len }

// Front returns the first element of list l or nil if the list is empty.
func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.head.next
}

// insert inserts e after at, increments l.len, and returns e.
func (l *List) insert(e, at *Element) *Element {
	e.next = at.next
	at.next = e
	l.len++
	if e.next == nil {
		l.tail = e
	}
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *List) insertValue(v any, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

// removeNext removes e.next from its list, decrements l.len
func (l *List) RemoveNext(e *Element) {
	if e.next == nil {
		return
	}
	e.next = e.next.next
	l.len--
	if e.next == nil {
		l.tail = e
	}
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *List) PushFront(v any) *Element {
	return l.insertValue(v, l.head)
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *List) PushBack(v any) *Element {
	return l.insertValue(v, l.tail)
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List) InsertAfter(v any, mark *Element) *Element {
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark)
}

func (l *List) dump() {
	if l.len == 0 {
		return
	}
	fmt.Print("head")
	curr := l.head.next
	for curr != nil {
		fmt.Printf("->%v", curr.Value)
		curr = curr.next
	}
	fmt.Println("")
}
