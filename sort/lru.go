package sort

import (
	"log"
)

type Node struct {
	key, val  int
	next, pre *Node
}
type lru struct {
	len        int
	head, tail *Node
	m          map[int]*Node
	n          int
}

func (t *lru) New() {
	t.head = new(Node)
	t.tail = new(Node)
	t.m = make(map[int]*Node)
	t.head.next = t.tail
	t.tail.pre = t.head
	t.n = 0
}
func (t *lru) Print() {
	p := t.head.next
	for p != t.tail {
		log.Printf("%d:%d ", p.key, p.val)
		p = p.next
	}
}
func (t *lru) Get(key int) int {
	if _, ok := t.m[key]; ok {
		node := t.m[key]
		//del
		t.del(node)
		//add
		t.MoveToHead(node)
		return node.val
	}
	return -1
}
func (t *lru) del(node *Node) {
	node.next.pre = node.pre
	node.pre.next = node.next
}
func (t *lru) MoveToHead(node *Node) {
	t.head.next.pre = node
	node.next = t.head.next
	node.pre = t.head
	t.head.next = node
}
func (t *lru) Put(key, val int) {
	if _, ok := t.m[key]; !ok {
		node := &Node{key: key, val: val}
		t.m[key] = node
		t.MoveToHead(node)
		t.n++
		if t.n > t.len {
			node := t.tail.pre
			t.del(node)
			delete(t.m, key)
			t.n--
		}
	} else {
		t.m[key].val = val
		t.MoveToHead(t.m[key])
	}
}
