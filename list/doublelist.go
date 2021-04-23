package list

type DNode struct {
	val       interface{}
	pre, next *DNode
}
type DoubleList struct {
	head, tail *DNode
	len        int
}

func (d *DoubleList) New() {
	d.head = new(DNode)
	d.tail = new(DNode)
	d.head.next = d.tail
	d.tail.pre = d.head
	d.len = 0
}
func (d *DoubleList) PushBack(e interface{}) {
	n := &DNode{val: e}
	d.Insert(n, d.head)
}
func (d *DoubleList) PushFront(e interface{}) {
	n := &DNode{val: e}
	d.Insert(n, d.tail.pre)
}

func (d *DoubleList) InsertAfter(e interface{}, at *DNode) {
	n := &DNode{val: e}
	d.Insert(n, at)
}
func (d *DoubleList) InsertBefore(e interface{}, at *DNode) {
	n := &DNode{val: e}
	d.Insert(n, at.pre)
}

//插入到at之后
func (d *DoubleList) Insert(e, at *DNode) {
	at.next.pre = e
	e.next = at.next
	e.pre = at
	at.next = e
	d.len++
}

//删除at之后
func (d *DoubleList) Remove(e *DNode) *DNode {
	if d.len == 0 {
		return nil
	}
	e.pre.next = e.next
	e.next.pre = e.pre
	e.next = nil
	e.pre = nil
	d.len--
	return e
}
func (d *DoubleList) Back() *DNode {
	if d.len == 0 {
		return nil
	}
	return d.tail.pre
}
func (d *DoubleList) Front() *DNode {
	if d.len == 0 {
		return nil
	}
	return d.head.next
}
