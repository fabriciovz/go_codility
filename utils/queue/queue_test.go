package queue

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("", func(t *testing.T) {

		myQueue := &queue{}
		fmt.Println(myQueue)
		myQueue.Enqueue(100)
		myQueue.Enqueue(200)
		myQueue.Enqueue(300)
		fmt.Println(myQueue)
		myQueue.Dequeue()
		fmt.Println(myQueue)
		fmt.Println(myQueue.Head())
		fmt.Println(myQueue.Tail())
	})
}
