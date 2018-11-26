package linklist


//单链表回文串判断
//对比链表中心点左右的字符判断
func (list LinkList)IsPalindrome() bool{
	length := list.GetLength()
	head   := list.GetHead()

	if length <= 0{
		return false
	}

	s := make([]string,0,length/2)
	cur := head

	for i := uint(1); i <= length; i++{
		cur = cur.GetNext()

		//如果是奇数个节点中间节点跳过
		if length % 2 != 0 && ((length + 1)/2==i){
			continue
		}

		if i <= length/2 {
			s = append(s,cur.GetValue().(string))
		}else{
			if s[length-i] != cur.GetValue().(string){
				return false
			}
		}
	}

	return  true
}

//链表反转
func (this *LinkList) Reverse() {
	if nil == this || nil == this.GetHead().GetNext() || nil == this.GetHead().GetNext().GetNext() {
		return
	}

	cur  := this.GetHead().GetNext()
	prev := cur.GetNext()
	head := this.GetHead()

	for prev != nil {
		cur.next = prev.next
		prev.next = head.next
		head.next = prev

		prev = cur.next
	}
}



