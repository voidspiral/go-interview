package list

import (
	"fmt"
	"testing"
)

func (l *SingleList) Print() {
	p := l.head.next
	for p != nil {
		fmt.Printf("%d ", p.val)
		p = p.next
	}
	fmt.Printf("\n")
}
func TestCrud(t *testing.T) {
	var v interface{}
	v = 1
	fmt.Println(v)

	l := Init()
	l.PushBack(1)
	l.PushBack(2)
	l.Print()
	l.PushFront(3)
	l.PushFront(4)
	l.Print()
	//l.PopBack()
	//l.Print()
	////fmt.Println(l.Back().val)
	//l.PopBack()
	//l.Print()
	////fmt.Println(l.Back().val)
	//l.PopBack()
	////fmt.Println(l.Back().val)
	//l.PopBack()
	//l.Print()
	//l.PopBack()
	////应该要报错
	//l.Print()

	l.PopFront()
	l.Print()
	l.PopFront()
	l.Print()
	l.Len()
	l.PopFront()
	l.Print()
	l.PopFront()
	l.Print()
	//应该要报错
	//l.PopFront()
	//l.Print()
}

func TestList(t *testing.T) {
	type list struct {
		val  int
		next *list
	}
	head := &list{}
	for i := 0; i < 10; i++ {
		head.next = &list{val: i, next: head.next}
	}
	//tail := head
	//for i := 0; i < 10; i++ {
	//	tail.next = &list{val: i, next: nil}
	//	tail = tail.next
	//
	p := head.next
	for i := 0; i < 4; i++ {
		p = p.next
	}
	p.next = p.next.next
	p = head.next
	for p != nil {
		fmt.Println(p.val)
		p = p.next
	}
}
