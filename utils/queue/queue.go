package queue

type Queue interface {
	Enqueue(i int)
	Dequeue() int
	Head() int
	Tail() int
	IsEmpty() bool
}

type queue struct {
	items []int
}

// Enqueue: adds a value at the end
func (q *queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue: remove value at the head
func (q *queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func (q *queue) Head() int {
	return q.items[0]
}

func (q *queue) Tail() int {
	return q.items[len(q.items)-1]
}

func (s *queue) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}
