package QueueBaseOnList

import (
	"fmt"
	"testing"
)

func TestLinkListNode_Enque(t *testing.T) {
	queue := NewLinkListNode()

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
