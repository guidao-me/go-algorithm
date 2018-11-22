package dlist

import (
	"fmt"
	"github.com/pkg/errors"
)

type DNode struct{
	data interface{}
	prev *DNode
	next *DNode
}

type DList struct {
	size  uint
	head *DNode
}

func NewDList()*DList{
	return &DList{
		size:0,
		head:&DNode{0,nil,nil},
	}
}

func NewDNode(value interface{})*DNode{
	return &DNode{
		value,
		nil,
		nil,
	}
}

func (this DList)GetHead()*DNode{
	return this.head
}

func (this DNode)GetValue()interface{}{
	return this.data
}

func (this DNode)GetNext()*DNode{
	return this.next
}

func (this DNode)GetPrev()*DNode{
	return this.prev
}

func (this *DList)Insert(i uint,node *DNode) error{
	if i < 0 || i > this.size {
		return errors.New("index out of range")
	}

	if node == nil {
		return errors.New("node is nil")
	}

	preItem := this.head
	for index := uint(0); index < i; index++ {
		preItem = preItem.next
	}

	node.next = preItem.next
	node.prev = preItem
	preItem.next = node

	this.size++

	return nil
}

func (this *DList)Delete(node *DNode) error{
	if nil == node {
		return errors.New("node is nil")
	}

	cur := this.GetHead().next

	for nil != cur{
		if cur.data == node.data {
			break
		}

		cur = cur.next
	}

	if cur == nil {
		return errors.New("not find node")
	}

	pre := cur.prev
	pre.next = cur.next
	if nil != cur.next{
		cur.next.prev = pre
	}

	return nil
}

func (this DList)FindByIndex(index uint)(*DNode,error){
	if index < 0 || index > this.size {
		return nil,errors.New("index out of range")
	}

	preItem := this.GetHead().next
	for i := uint(0); i < index; i++{
		preItem = preItem.next
	}

	return preItem,nil
}

func (this DList)Traverse(){
	pre := this.GetHead().next

	for nil != pre {
		fmt.Printf("%v\n",pre.GetValue())

		pre = pre.next
	}
}
