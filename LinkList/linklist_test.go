package main

import (
	"fmt"
	"testing"
)

func TestLinkList_Insert(t *testing.T) {
	n1 := ListNode{1,nil}
	n2 := ListNode{2,nil}
	n3 := ListNode{3,nil}

	list := NewLinkList()

	list.Insert(0,&n1)
	list.Insert(1,&n2)
	list.Insert(2,&n3)

	list.Print()

	n4 := ListNode{10,nil}
	list.Insert(0,&n4)
	list.Print()
}

func TestLinkList_Delete(t *testing.T) {
	n1 := ListNode{1,nil}
	n2 := ListNode{2,nil}
	n3 := ListNode{3,nil}

	list := NewLinkList()

	list.Insert(0,&n1)
	list.Insert(1,&n2)
	list.Insert(2,&n3)

	list.Print()

	fmt.Println()
	list.Delete(&n1)
	list.Delete(&n2)


	list.Print()
}


func TestLinkList_FindByIndex(t *testing.T) {
	n1 := ListNode{1,nil}
	n2 := ListNode{2,nil}
	n3 := ListNode{3,nil}

	list := NewLinkList()

	list.Insert(0,&n1)
	list.Insert(1,&n2)
	list.Insert(2,&n3)

	list.Print()

	fmt.Println()

	node, err := list.FindByIndex(1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("node value:%v\n",node.GetValue())
}