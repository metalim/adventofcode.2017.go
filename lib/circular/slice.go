package circular

// NewSlice with values
func NewSlice(vals ...Value) Buffer {
	return &slice{data: vals}
}

func (s *slice) Skip(steps int) Buffer {
	if len(s.data) > 0 {
		s.head = (s.head + steps) % len(s.data)
	}
	return s
}

func (s *slice) InsertBefore(val Value) Buffer {
	var empty Value
	s.data = append(s.data, empty)
	copy(s.data[s.head+1:], s.data[s.head:])
	s.data[s.head] = val
	return s
}

func (s *slice) InsertAfter(val Value) Buffer {
	var empty Value
	s.data = append(s.data, empty)
	s.head++
	copy(s.data[s.head+1:], s.data[s.head:])
	s.data[s.head] = val
	return s
}

func (s *slice) Get() (val Value) {
	return s.data[s.head]
}

func (s *slice) Set(val Value) {
	s.data[s.head] = val
}

func (s *slice) Length() int {
	return len(s.data)
}

type slice struct {
	data []Value
	head int
}
