package linklist

import (
	"fmt"
	"testing"
)

func TestLinkList_Insert(t *testing.T) {
	list := NewLinkList()

	for i := 0; i < 10; i++{
		list.Insert(uint(i),NewLinkNode(i+10))
	}

	list.Print()

	n4 := ListNode{10,nil}
	list.Insert(0,&n4)
	list.Print()
}

func TestLinkList_Delete(t *testing.T) {
	list := NewLinkList()

	for i := 0; i < 10; i++{
		list.Insert(uint(i),NewLinkNode(i+10))
	}

	list.Print()

	fmt.Println()
	list.Delete(NewLinkNode(11))
	list.Delete(NewLinkNode(10))
	list.Delete(NewLinkNode(19))


	list.Print()
}


func TestLinkList_FindByIndex(t *testing.T) {
	list := NewLinkList()

	for i := 0; i < 10; i++{
		list.Insert(uint(i),NewLinkNode(i+10))
	}

	list.Print()

	fmt.Println()

	node, err := list.FindByIndex(0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("node value:%v\n",node.GetValue())
}