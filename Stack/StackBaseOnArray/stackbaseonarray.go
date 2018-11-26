package StackBaseOnArray

import "fmt"

type Stack struct {
	data []interface{}
	top  int
}


func NewStack() *Stack{
	return &Stack{
		data:make([]interface{},0,32),
		top:-1,
	}
}

func (this Stack) isEmpty()bool{
	if this.top < 0{
		return true
	}

	return false
}

func (this *Stack) Push(value interface{}){
	this.top++
	this.data = append(this.data,value)
}

func (this *Stack) Pop()interface{}{
	if this.isEmpty(){
		return nil
	}

	value := this.data[this.top]

	this.top--

	return value
}

func (this *Stack) Top()interface{}{
	if this.isEmpty(){
		return nil
	}

	return this.data[this.top]
}

func (this Stack) Print(){
	if this.isEmpty(){
		fmt.Println("empty Stack")
	}

	for i := this.top; i >= 0; i--{
		fmt.Printf("%v\n",this.data[i])
	}
}
