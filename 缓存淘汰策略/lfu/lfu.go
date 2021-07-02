package lfu

type LFUCache struct {
	keyNode   map[int]*valueNode
	visitList map[int]*valueList
	cap       int
	len       int
	min       int
}

type valueNode struct {
	prev  *valueNode
	next  *valueNode
	key   int
	value int
	visit int
}

type valueList struct {
	handle *valueNode
	size   int
}

func newList() *valueList {
	list := &valueList{}
	list.handle = &valueNode{}
	list.handle.next = list.handle
	list.handle.prev = list.handle
	return list
}

func (l *valueList) push(node *valueNode) {
	node.next = l.handle
	node.prev = l.handle.prev
	node.next.prev = node
	node.prev.next = node
	l.size++
}
func (l *valueList) pop(node *valueNode) {
	node.next.prev = node.prev
	node.prev.next = node.next
	node.next = nil
	node.prev = nil
	l.size--
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		keyNode:   make(map[int]*valueNode, capacity),
		visitList: make(map[int]*valueList),
		cap:       capacity,
		len:       0,
		min:       0,
	}
}

func (c *LFUCache) Get(key int) int {
	if c.len == 0 {
		return -1
	}
	if node, ok := c.keyNode[key]; ok {
		c.update(node)
		return node.value
	}
	return -1
}

func (c *LFUCache) Put(key int, value int) {
	if c.cap == 0 {
		return
	}
	if node, ok := c.keyNode[key]; ok {
		node.value = value
		c.update(node)
		return
	}
	for c.len >= c.cap { // 用 for 不用 if 可支持扩容缩容
		c.pop()
	}
	node := &valueNode{key: key, value: value, visit: 1}
	c.insert(node)
}

func (c *LFUCache) update(node *valueNode) {
	list := c.visitList[node.visit]
	list.pop(node)
	if list.size == 0 && node.visit == c.min {
		c.min++
	}
	node.visit++
	list, ok := c.visitList[node.visit]
	if !ok {
		list = newList()
		c.visitList[node.visit] = list
	}
	list.push(node)
}

func (c *LFUCache) insert(node *valueNode) {
	visit := node.visit
	list, ok := c.visitList[visit]
	if !ok {
		list = newList()
		c.visitList[visit] = list
	}
	list.push(node)
	c.keyNode[node.key] = node
	c.len++
	c.min = visit
}

func (c *LFUCache) pop() *valueNode {
	if c.len == 0 {
		return nil
	}
	// 弹出之后会立即插入新值, 所以这里不需要更新 min 值.
	visit := c.min
	list := c.visitList[visit]
	node := list.handle.next
	list.pop(node)
	delete(c.keyNode, node.key)
	c.len--
	return node
}
