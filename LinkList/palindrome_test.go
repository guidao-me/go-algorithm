package linklist

import (
	"fmt"
	"testing"
)

func TestLinkList_IsPalindrome(t *testing.T) {
	list := NewLinkList()

	for i := 0; i < 10; i++{
		list.Insert(uint(i),NewLinkNode(fmt.Sprintf("%d",i)))
	}

	for i := 10; i < 20; i++{
		list.Insert(uint(i),NewLinkNode(fmt.Sprintf("%d",19-i)))
	}

	list.Print()


	if list.IsPalindrome(){
		fmt.Println("list is Palindrome")
	}else{
		fmt.Println("list is not Palindrome")
	}
}

func TestLinkList_Reverse(t *testing.T) {
	list := NewLinkList()

	for i := 0; i < 10; i++{
		list.Insert(uint(i),NewLinkNode(i))
	}

	list.Print()

	list.Reverse()

	fmt.Println()
	list.Print()

	fmt.Println("pragram exit")
}

