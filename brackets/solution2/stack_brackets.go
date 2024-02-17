package solution2

type Stack interface {
	Push(i byte)
	Pop() byte
	IsEmpty() bool
}

type stack struct {
	items []byte
}

// push will add a value at the end (top)
func (s *stack) Push(i byte) {
	s.items = append(s.items, i)
}

// push will remove a value at the end (top)
func (s *stack) Pop() byte {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func (s *stack) Top() byte {
	return s.items[len(s.items)-1]
}

func (s *stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}
