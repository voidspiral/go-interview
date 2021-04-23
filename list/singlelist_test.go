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
