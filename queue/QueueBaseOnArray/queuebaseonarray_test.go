package QueueBaseOnArray

import (
	"fmt"
	"testing"
)

func TestQueue_Enque(t *testing.T) {
	queue := NewQueue(10)

	for i := 0; i < 10; i++{
		queue.Enque(i)
	}

	queue.Print()

	fmt.Println()

	queue.Deque()
	queue.Deque()
	queue.Deque()
	queue.Print()
}
