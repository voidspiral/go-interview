package lru

import "fmt"

type Node struct {
	key, val   int
	prev, next *Node
}

type LRUCache struct {
	cache      map[int]*Node
	head, tail *Node
	len        int
	cap        int
}

func Constructor(capacity int) LRUCache {
	var l LRUCache
	l.cache = make(map[int]*Node)
	l.head = new(Node)
	l.tail = new(Node)
	l.len = 0
	l.cap = capacity
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (l *LRUCache) addToHead(n *Node) {
	n.next = l.head.next
	l.head.next.prev = n
	l.head.next = n
	n.prev = l.head
}

func (l *LRUCache) del(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (l *LRUCache) moveToHead(n *Node) {
	l.del(n)
	n.next = l.head.next
	l.head.next.prev = n
	l.head.next = n
	n.prev = l.head
}

func (l *LRUCache) Get(key int) int {
	if v, ok := l.cache[key]; ok {
		l.moveToHead(v)
		return v.val
	}
	return -1
}

func (l *LRUCache) Print() {
	p := l.head.next
	fmt.Println("map", l.cache)
	for p != l.tail {
		fmt.Printf("%v %v\n", p.key, p.val)
		p = p.next
	}
}
func (l *LRUCache) Put(key int, value int) {
	// 1, 1, 1
	// 2, 1, 1, 1
	//        , _
	if _, ok := l.cache[key]; !ok {
		n := &Node{
			key: key,
			val: value,
		}
		l.cache[key] = n
		l.addToHead(n)
		l.len++
		if l.len > l.cap {
			d := l.tail.prev
			l.del(d)
			l.len--
			delete(l.cache, d.key)
		}
	} else {
		node := l.cache[key]
		node.val = value
		l.moveToHead(node)
	}
	return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
