package gtl

// List is a generic interface for a list.
type List[T any] interface {
	// Front returns the first element in the list.
	Front() *ListElement[T]
	// Back returns the last element in the list.
	Back() *ListElement[T]
	// PushFront adds an element to the front of the list.
	PushFront(value T) *ListElement[T]
	// PushBack adds an element to the back of the list.
	PushBack(value T) *ListElement[T]
	// PopFront removes and returns the first element in the list.
	PopFront() *ListElement[T]
	// PopBack removes and returns the last element in the list.
	PopBack() *ListElement[T]
	// Remove removes the element from the list.
	Remove(e *ListElement[T])
	// Len returns the number of elements in the list.
	Len() int
	// Next returns the next element in the list.
	Next(elem *ListElement[T]) *ListElement[T]
	// Prev returns the previous element in the list.
	Prev(elem *ListElement[T]) *ListElement[T]
}

// ListElement is an element in a list.
type ListElement[T any] struct {
	value T
	next  *ListElement[T]
	prev  *ListElement[T]
	list  List[T]
}

func (e *ListElement[T]) Value() T {
	return e.value
}

// linkedList is a doubly linked list.
type linkedList[T any] struct {
	root ListElement[T]
	len  int
}

// NewLinkedList returns a new linked list.
func NewLinkedList[T any]() List[T] {
	l := linkedList[T]{}
	l.root.next = &l.root
	l.root.prev = &l.root
	return &l
}

func (l *linkedList[T]) Front() *ListElement[T] {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *linkedList[T]) Back() *ListElement[T] {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

func (l *linkedList[T]) insert(e, at *ListElement[T]) *ListElement[T] {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

func (l *linkedList[T]) PushFront(value T) *ListElement[T] {
	e := &ListElement[T]{value: value}
	return l.insert(e, &l.root)
}

func (l *linkedList[T]) PushBack(value T) *ListElement[T] {
	e := &ListElement[T]{value: value}
	return l.insert(e, l.root.prev)
}

func (l *linkedList[T]) PopFront() *ListElement[T] {
	elem := l.Front()
	l.Remove(elem)
	return elem
}

func (l *linkedList[T]) PopBack() *ListElement[T] {
	elem := l.Back()
	l.Remove(elem)
	return elem
}

func (l *linkedList[T]) Remove(elem *ListElement[T]) {
	if elem == nil {
		return
	}
	if elem.list == nil {
		return
	}
	elem.prev.next = elem.next
	elem.next.prev = elem.prev
	elem.next = nil // avoid memory leaks
	elem.prev = nil // avoid memory leaks
	elem.list = nil
	l.len--
}

func (l *linkedList[T]) Len() int {
	return l.len
}

func (l *linkedList[T]) Next(elem *ListElement[T]) *ListElement[T] {
	if elem == nil {
		return nil
	}
	if elem.list == nil {
		return nil
	}
	if elem.next == nil {
		return nil
	}
	if elem.next == &elem.list.(*linkedList[T]).root {
		return nil
	}
	return elem.next
}

func (l *linkedList[T]) Prev(elem *ListElement[T]) *ListElement[T] {
	if elem == nil {
		return nil
	}
	if elem.list == nil {
		return nil
	}
	if elem.prev == nil {
		return nil
	}
	if elem.prev == &elem.list.(*linkedList[T]).root {
		return nil
	}
	return elem.prev
}
