package heap

import "fmt"

type heap struct {
	data []int
	size int
}

//func (h *heap) New(cap int) {
//	h.data = make([]int, 10, cap)
//	h.size = 0
//	fmt.Println(h.data)
//}
func (h *heap) New() {
	h.data = []int{}
}
func (h *heap) Push(v int) {
	h.data = append(h.data, v)
	i := len(h.data) - 1
	//log.Println(i)
	for i > 0 {
		p := (i - 1) / 2
		if h.data[p] <= v {
			break
		}
		h.data[i] = h.data[p]
		i = p
	}
	h.data[i] = v
}
func (h *heap) Pop() (ret int, err error) {
	if len(h.data) == 0 {
		return -1, fmt.Errorf("can not pop back")
	}
	ret = h.data[0]
	size := len(h.data) - 1
	//log.Println("size", size)
	x := h.data[size]

	i := 0
	for i*2+1 < size {
		lchild, rchild := i*2+1, i*2+2
		if rchild < size && h.data[rchild] < h.data[lchild] {
			lchild = rchild
		}
		if h.data[lchild] >= x {
			break
		}
		h.data[i] = h.data[lchild]
		i = lchild
	}
	h.data[i] = x
	h.data = h.data[:size]
	return ret, nil
}

//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//type heap struct {
//	arr []*ListNode
//	// len int
//}
//
//func (h *heap) Init() {
//	// h.len = 0
//	h.arr = []*ListNode{}
//}
//func (h *heap) Push(v *ListNode) {
//	h.arr = append(h.arr, v)
//	Len := len(h.arr)
//	i := Len - 1
//	for i > 0 {
//		p := (i - 1) / 2
//		if h.arr[p].Val < v.Val {
//			break
//		}
//		h.arr[i] = h.arr[p]
//		i = p
//	}
//	h.arr[i] = v
//	return
//}
//func (h *heap) Pop() interface{} {
//	if len(h.arr) < 0 {
//		return nil
//	}
//	x := h.arr[0]
//	Len := len(h.arr) - 1
//	v := h.arr[Len]
//	i := 0
//	for 2*i+1 < Len-1 {
//		lchild, rchild := 2*i+1, 2*i+2
//		if rchild < Len && h.arr[rchild].Val < h.arr[lchild].Val {
//			lchild = rchild
//		}
//		if h.arr[lchild].Val >= x.Val {
//			break
//		}
//		h.arr[i] = h.arr[lchild]
//		i = lchild
//	}
//	h.arr[i] = v
//	h.arr = h.arr[:Len]
//	return x
//}
//func mergeKLists(lists []*ListNode) *ListNode {
//	h := new(heap)
//	h.Init()
//	dummy := &ListNode{}
//	tail := dummy
//	for _, list := range lists {
//		if list != nil {
//			h.Push(list)
//		}
//	}
//	// fmt.Println(h.Pop())
//	for len(h.arr) != 0 {
//		Min := h.Pop().(*ListNode)
//		if Min.Next != nil {
//			h.Push(Min.Next)
//		}
//		tail.Next = Min
//		tail = tail.Next
//
//	}
//	return dummy.Next
//}
