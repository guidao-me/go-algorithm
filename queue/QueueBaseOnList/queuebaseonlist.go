package QueueBaseOnList

import (
	"fmt"
	"github.com/pkg/errors"
)

type ListNode struct {
	data interface{}
	next *ListNode
}

type LinkListNode struct {
	head  	*ListNode
	tail  	*ListNode
	length  int
}

func NewLinkListNode()*LinkListNode{
	return &LinkListNode{
		head:nil,
		tail:nil,
		length:0,
	}
}

func (this *LinkListNode)Enque(v interface{}){
	node := &ListNode{v,nil}

	if this.head == nil {
		this.head = node
		this.tail = node
	}else{
		this.tail.next = node
		this.tail = node
	}

	this.length++
}

func (this *LinkListNode)Deque()(interface{},error){
	if this.head == this.tail{
		return 0,errors.New("Queue Empty")
	}

	v := this.head.data
	this.head = this.head.next
	this.length--

	return v,nil
}

func (this LinkListNode)Print(){
	prev := this.head

	for prev != nil {
		fmt.Println(prev.data)

		prev = prev.next
	}
}
