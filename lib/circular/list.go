package circular

// NewList with values.
func NewList(vals ...Value) Buffer {
	l := &list{}
	for _, val := range vals {
		l.InsertAfter(val)
	}
	l.Skip(1) // back to head
	return l
}

// Length of the list.
func (l *list) Length() int {
	return l.len
}

// Skip steps.
func (l *list) Skip(steps Value) Buffer {
	if l.head == nil {
		return l
	}
	for ; steps > 0; steps-- {
		l.head = l.head.next
	}
	for ; steps < 0; steps++ {
		l.head = l.head.prev
	}
	return l
}

// InsertBefore current head.
func (l *list) InsertBefore(val Value) Buffer {
	if l.head == nil {
		l.head = &node{val: val}
		l.head.next = l.head
		l.head.prev = l.head
		l.len = 1
		return l
	}
	l.head = &node{val: val, next: l.head, prev: l.head.prev}
	l.head.prev.next = l.head
	l.head.next.prev = l.head
	l.len++
	return l
}

// InsertAfter current head, and select it (== Skip(1) in array).
func (l *list) InsertAfter(val Value) Buffer {
	if l.head == nil {
		l.head = &node{val: val}
		l.head.next = l.head
		l.head.prev = l.head
		l.len = 1
		return l
	}
	l.head = &node{val: val, next: l.head.next, prev: l.head}
	l.head.prev.next = l.head
	l.head.next.prev = l.head
	l.len++
	return l
}

// Get head value.
func (l *list) Get() (val Value) {
	if l.head == nil {
		return
	}
	return l.head.val
}

// Set head value.
func (l *list) Set(val Value) {
	l.head.val = val // panic if list is empty
}

type list struct {
	head *node
	len  int
}

type node struct {
	val        Value
	next, prev *node
}
