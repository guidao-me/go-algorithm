package dlist

import (
	"fmt"
	"testing"
)

func TestDList_Insert(t *testing.T) {
	dlist := NewDList()

	for i := 0; i < 10; i++ {
		dlist.Insert(uint(i),NewDNode(i+10))
	}

	dlist.Traverse()
}

func TestDList_FindByIndex(t *testing.T) {
	dlist := NewDList()

	for i := 0; i < 10; i++ {
		dlist.Insert(uint(i),NewDNode(i+10))
	}

	d, err := dlist.FindByIndex(5)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("value: %v\n",d.GetPrev().GetValue())
}

func TestDList_Delete(t *testing.T) {
	dlist := NewDList()

	for i := 0; i < 10; i++ {
		dlist.Insert(uint(i),NewDNode(i+10))
	}

	dlist.Delete(NewDNode(19))

	dlist.Traverse()

	dlist.Delete(NewDNode(10))

	dlist.Traverse()
}
