package QueueBaseOnArray

import (
	"fmt"
	"github.com/pkg/errors"
)

type Queue struct {
	data 		[]interface{}
	capacity    int
	head        int
	tail        int
}

func NewQueue(size int)*Queue{
	return &Queue{
		data:make([]interface{},size),
		capacity:size,
		head:0,
		tail:0,
	}
}

func (this *Queue)Enque(v interface{})error{
	if this.tail == this.capacity{
		return errors.New("Queue full")
	}

	this.data[this.tail] = v
	this.tail++

	return nil
}

func (this *Queue)Deque()(interface{},error){
	if this.tail == this.head{
		return 0,errors.New("Queue Empty")
	}

	v := this.data[this.head]
	this.head++
	return v,nil
}

func (this *Queue)Print(){
	for i := this.head; i < this.tail; i++{
		fmt.Println(this.data[i])
	}
}


