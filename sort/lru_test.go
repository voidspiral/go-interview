package sort

import (
	"container/list"
	"log"
	"testing"
)

var l list.List
l = l.New()
func TestLru(t *testing.T) {
	h := new(lru)
	h.New()
	h.len = 2
	node := &Node{key: 1, val: 2}

	node.next = h.head.next
	h.head.next.pre = node
	node.pre = h.head
	h.head.next = node
	h.Print()
	//del
	node.next.pre = node.pre
	node.pre.next = node.next
	log.Println("del")
	h.Print()
	t.Log("go")
	h.Put(1, 1)
	h.Put(2, 2)
	log.Println(h.Get(1))
	h.Put(3, 3)
	h.Print()
}
