// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtl

import "testing"

func checkListLen[T any](t *testing.T, l List[T], len int) bool {
	if n := l.Len(); n != len {
		t.Errorf("l.Len() = %d, want %d", n, len)
		return false
	}
	return true
}

func checkListPointers[T any](t *testing.T, l List[T], es []*ListElement[T]) {
	root := l.Front()

	if !checkListLen(t, l, len(es)) {
		return
	}

	// zero length lists must be the zero value or properly initialized (sentinel circle)
	if len(es) == 0 && root != nil {
		t.Errorf("l.Front() = %p, want nil", root)
		return
	}
	// len(es) > 0

	// check internal and external prev/next connections
	for i, e := range es {
		prev := root
		Prev := (*ListElement[T])(nil)
		if i > 0 {
			prev = es[i-1]
			Prev = prev
		}
		if p := l.Prev(e); p != Prev {
			t.Errorf("elt[%d](%p).Prev() = %p, want %p", i, e, p, Prev)
		}

		next := root
		Next := (*ListElement[T])(nil)
		if i < len(es)-1 {
			next = es[i+1]
			Next = next
		}
		if n := l.Next(e); n != Next {
			t.Errorf("elt[%d](%p).Next() = %p, want %p", i, e, n, Next)
		}
	}
}

func TestList(t *testing.T) {
	l := NewLinkedList[int]()
	checkListPointers(t, l, []*ListElement[int]{})

	// Single element list
	e := l.PushFront(1)
	checkListPointers(t, l, []*ListElement[int]{e})
	l.Remove(e)
	checkListPointers(t, l, []*ListElement[int]{})
	e = l.PushFront(1)
	checkListPointers(t, l, []*ListElement[int]{e})
	l.Remove(e)
	e = l.PushBack(1)
	checkListPointers(t, l, []*ListElement[int]{e})

	// Bigger list
	l = NewLinkedList[int]()
	e2 := l.PushFront(2)
	e1 := l.PushFront(1)
	e3 := l.PushBack(3)
	e4 := l.PushBack(4)
	checkListPointers(t, l, []*ListElement[int]{e1, e2, e3, e4})

	l.Remove(e2)
	checkListPointers(t, l, []*ListElement[int]{e1, e3, e4})

	// move from middle
	l.Remove(e3)
	e3 = l.PushFront(e3.Value())
	checkListPointers(t, l, []*ListElement[int]{e3, e1, e4})

	l.Remove(e1)
	e1 = l.PushFront(e1.Value())
	l.Remove(e3)
	e3 = l.PushBack(e3.Value())
	checkListPointers(t, l, []*ListElement[int]{e1, e4, e3})

	// move from back
	l.Remove(e3)
	e3 = l.PushFront(e3.Value())
	checkListPointers(t, l, []*ListElement[int]{e3, e1, e4})

	// Check standard iteration.
	sum := 0
	for e := l.Front(); e != nil; e = l.Next(e) {
		sum += e.Value()
	}
	if sum != 8 {
		t.Errorf("sum over l = %d, want 4", sum)
	}

	// Clear all elements by iterating
	var next *ListElement[int]
	for e := l.Front(); e != nil; e = next {
		next = l.Next(e)
		l.Remove(e)
	}
	checkListPointers(t, l, []*ListElement[int]{})
}

func checkList(t *testing.T, l List[int], es []int) {
	if !checkListLen(t, l, len(es)) {
		return
	}

	i := 0
	for e := l.Front(); e != nil; e = l.Next(e) {
		le := e.Value()
		if le != es[i] {
			t.Errorf("elt[%d].Value = %v, want %v", i, le, es[i])
		}
		i++
	}
}

func TestExtending(t *testing.T) {
	l1 := NewLinkedList[int]()
	l2 := NewLinkedList[int]()
	d1 := []int{1, 2, 3}
	d2 := []int{4, 5}

	for _, d := range d1 {
		l1.PushBack(d)
	}
	for _, d := range d2 {
		l2.PushBack(d)
	}

	checkList(t, l1, d1)
	checkList(t, l2, d2)

	l3 := NewLinkedList[int]()
	for _, d := range d1 {
		l3.PushBack(d)
	}
	checkList(t, l3, d1)
	for _, d := range d2 {
		l3.PushBack(d)
	}
	checkList(t, l3, []int{1, 2, 3, 4, 5})
}

func TestRemove(t *testing.T) {
	l := NewLinkedList[int]()
	e1 := l.PushBack(1)
	e2 := l.PushBack(2)
	checkListPointers(t, l, []*ListElement[int]{e1, e2})
	e := l.Front()
	l.Remove(e)
	checkListPointers(t, l, []*ListElement[int]{e2})
	l.Remove(e)
	checkListPointers(t, l, []*ListElement[int]{e2})
}

// Test PushFront, PushBack with uninitialized List
func TestZeroList(t *testing.T) {
	var l1 = NewLinkedList[int]()
	l1.PushFront(1)
	checkList(t, l1, []int{1})

	var l2 = NewLinkedList[int]()
	l2.PushBack(1)
	checkList(t, l2, []int{1})
}
