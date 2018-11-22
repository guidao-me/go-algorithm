package main

import (
	"errors"
	"fmt"
)

type ListNode struct {
	data interface{}
	next *ListNode
}

type LinkList struct {
	head   *ListNode
	length uint
}

func NewLinkList() *LinkList {
	return &LinkList{
		head:   &ListNode{0, nil},
		length: 0,
	}
}

func (this *ListNode) GetNext() *ListNode {
	return this.next
}

func (this *ListNode) GetValue() interface{} {
	return this.data
}

func (this *LinkList) Insert(i uint, node *ListNode) error {
	if i < 0 || node == nil || i > this.length{
		return errors.New("index out of range or node is nil")
	}

	preItem := (*this).head
	for j := uint(0); j < i; j++ {
		preItem = preItem.next
	}

	node.next = preItem.next
	preItem.next = node

	this.length++

	return nil
}

func (this *LinkList) Delete(node *ListNode)error{
	if nil == node{
		return errors.New("node is nil")
	}

	pre := this.head
	cur := this.head.next

	for nil != cur{
		if cur == node {
			break
		}
		pre = cur
		cur = cur.next
	}

	if nil == cur{
		return errors.New("not find node")
	}

	pre.next = cur.next
	node = nil
	this.length--

	return nil
}

func (this LinkList) FindByIndex(index uint)(*ListNode,error){
	if index < 0 || index >= this.length {
		return nil,errors.New("out of range")
	}

	preItem := this.head.next
	for i := uint(0); i < index; i++ {
		preItem = preItem.next
	}

	return preItem,nil
}

func (this LinkList)Print(){
	pre := this.head.next

	for nil != pre{
		fmt.Printf("%v\n",pre.GetValue())
		pre = pre.next
	}
}