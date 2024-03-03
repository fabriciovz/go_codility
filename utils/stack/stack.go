package stack

type Stack interface {
	Push(i int)
	Pop() int
	Top() int
	IsEmpty() bool
}

type stack struct {
	items []int
}

// push will add a value at the end (top)
func (s *stack) Push(i int) {
	s.items = append(s.items, i)
}

// push will remove a value at the end (top)
func (s *stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func (s *stack) Top() int {
	return s.items[len(s.items)-1]
}

func (s *stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}
