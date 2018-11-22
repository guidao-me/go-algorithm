package array

import (
	"fmt"
	"github.com/youtube/vitess/go/vt/log"
	"testing"
)

func TestArray_Insert(t *testing.T) {
	cap := 10
	arr := NewArray(uint(cap))


	for i := 0; i < cap;i++{
		err := arr.Insert(uint(i),i+100)
		if err != nil {
			log.Fatal(err)
		}
	}

	arr.Print()
}

func TestArray_Delete(t *testing.T) {
	cap := 10
	arry := NewArray(uint(cap))

	for i := 0; i < cap; i++{
		err := arry.Insert(uint(i),i+100)
		if err != nil {
			log.Fatal(err)
		}
	}


	arry.Print()

	fmt.Println()

	arry.Delete(9)
	arry.Delete(5)

	arry.Print()
}

func TestArray_Find(t *testing.T) {
	cap := 10
	arry := NewArray(uint(cap))

	for i := 0; i < cap; i++{
		err := arry.Insert(uint(i),i+100)
		if err != nil {
			log.Fatal(err)
		}
	}

	arry.Print()

	fmt.Println()

	value, err := arry.Find(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("array[%d]=%v\n",1,value)
}