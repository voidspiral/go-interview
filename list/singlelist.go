package list

import (
	"fmt"
	"log"
)

type Node struct {
	val  int
	next *Node
}

//头尾两个哨兵结点，实现单链表
type SingleList struct {
	head, tail *Node
	len        int
}

func Init() *SingleList {
	l := new(SingleList)
	l.head = new(Node)
	l.tail = new(Node)
	l.tail = l.head
	//l.head = l.tail
	return l
}

//尾插法
func (l *SingleList) PushBack(val int) {
	node := &Node{val: val, next: nil}
	l.tail.next = node
	l.tail = node
	l.len++
}

//头插法
func (l *SingleList) PushFront(val int) {
	node := &Node{val: val, next: nil}
	node.next = l.head.next
	l.head.next = node
	l.len++
}

//删除最后一个元素
func (l *SingleList) PopBack() (*Node, error) {
	if l.len <= 0 {
		log.Fatal("fatal error : empty list can not pop back")
	}
	e, err := l.Remove()
	if err != nil {
		return nil, fmt.Errorf("unkown error can not pop back: %v", err)
	}
	l.len--
	return e, nil
}

//删除队首元素
func (l *SingleList) PopFront() (*Node, error) {
	if l.len <= 0 {
		log.Fatal("fatal error : empty list can not pop font")
	}
	p := l.head.next
	l.head.next = p.next
	l.len--
	return p, nil
}

//单链表时间复杂度O(n)
func (l *SingleList) Remove() (n *Node, err error) {
	p := l.head.next
	pre := l.head
	for p.next != nil {
		pre = p
		p = p.next
	}
	pre.next = pre.next.next
	//利用go的垃圾回收
	pre.next = nil
	return n, nil
}
func (l *SingleList) Len() int {
	log.Printf("list len :%v", l.len)
	return l.len
}
