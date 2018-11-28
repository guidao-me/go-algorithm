package CircularQueue

import "errors"

type CircularQueue struct {
	data []interface{}
	capacity int
	head     int
	tail     int
}

func NewCircularQueue(size int)*CircularQueue{
	return &CircularQueue{
		data:make([]interface{},size),
		capacity:size,
		head:0,
		tail:0,
	}
}

func (this CircularQueue)isEmpty()bool{
	if this.tail == this.head{
		return true
	}
	return false
}

func (this CircularQueue)isFull()bool{
	if (this.tail+1)%this.capacity == this.head{
		return true
	}

	return false
}

func (this *CircularQueue)EnQueue(v interface{})error{
	if this.isFull(){
		return errors.New("CircularQueue Full")
	}

	this.data[this.tail] = v
	this.tail = (this.tail + 1)%this.capacity

	return nil
}

func (this *CircularQueue)Deque(v interface{})error{
	if this.isEmpty(){

	}
}


