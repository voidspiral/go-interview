package heap

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func Test_heap_new(t *testing.T) {
	h := heap{}
	h.New()
	n := 10000
	correct := make([]int, n)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		correct[i] = rand.Intn(10)
		h.Push(correct[i])
	}
	//for i := 4; i >= 0; i-- {
	//	h.Push(i)
	//}
	//for i := 0; i < 5; i++ {
	//	t.Log(h.Pop())
	//}
	for i := 0; i < n; i++ {
		ans[i], _ = h.Pop()
		//t.Log(ans[i])
	}
	t.Log("correct:", correct)
	t.Log("ans:", ans)
	sort.Ints(correct)
	t.Log(correct)
	t.Log("test:", reflect.DeepEqual(correct, ans))
	if ok := reflect.DeepEqual(correct, ans); !ok {
		t.Fail()
	}

}
