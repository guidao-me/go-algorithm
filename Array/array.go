package array

import (
	"errors"
	"fmt"
)

type Array struct {
	data  []interface{}				//存储数组数据
	length   uint					//数组当前长度
}

func NewArray(capacity uint)*Array{
	if capacity == 0 {
		return nil
	}

	return &Array{
		data:make([]interface{},capacity,capacity),
		length:0,
	}
}

func (this Array) Len()uint{
	return this.length
}
//通过下标查找数组元素
func (this Array) Find(index uint)(interface{},error){
	if index < 0 || index >= this.length{
		return 0,errors.New("out of index range")
	}

	return this.data[index],nil
}

//插入数值到索引index上
func (this *Array) Insert(index uint,value interface{}) error{
	if this.Len() == uint(cap(this.data)){
		return errors.New("full array")
	}

	if index < 0 || index > this.length{
		return errors.New("out of index range")
	}

	for i := this.length; i > index; i-- {
		this.data[i] = this.data[i-1]
	}

	this.data[index] = value
	this.length++

	return nil
}

func (this *Array) Delete(index uint) (interface{},error){
	if index < 0 || index >= this.length{
		return 0,errors.New("out of index range")
	}

	v := this.data[index]
	for i := index; i < this.length-1; i++{
		this.data[i] = this.data[i+1]
	}
	this.length--

	return v,nil
}


func (this Array)Print(){
	for i := uint(0); i < this.length; i++ {
		fmt.Printf("data[%d] = %v\n",i,this.data[i])
	}
}