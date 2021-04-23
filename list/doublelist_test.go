package list

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
)

func (d *DoubleList) Print() {
	p := d.head.next
	for p != d.tail {
		fmt.Printf("%v ", p.val)
		p = p.next
	}
	fmt.Printf("\n")
}
func (d *DoubleList) Len() int {
	return d.len
}
func TestInsert(t *testing.T) {
	var l DoubleList
	l.New()
	for i := 0; i < 5; i++ {
		n := &DNode{val: i}
		l.Insert(n, l.head)
	}
	l.Print()
	l.New()
	for i := 0; i < 5; i++ {
		n := &DNode{val: i}
		l.Insert(n, l.tail.pre)
	}
	l.Print()
	//对空链删除问题不大
	l.Remove(l.Front())
	l.Remove(l.Back())
	l.Print()
}
func TestDoubleList_Crud(t *testing.T) {
	rand.Seed(42)
	var dl DoubleList
	dl.New()
	n := rand.Intn(10)
	log.Printf("random list len:%v ", n)
	for i := 0; i < n; i++ {
		dl.PushBack(i)
	}
	dl.Print()
	dl.New()
	log.Printf("random list len:%v ", n)
	for i := 0; i < n; i++ {
		dl.PushFront(i)
	}
	dl.Print()
	for dl.Len() != 0 {
		dl.Remove(dl.Front())
		dl.Print()
	}
	//空链表删除不会出错
	dl.Remove(dl.Front())
	dl.Remove(dl.Back())
	dl.Print()
}
