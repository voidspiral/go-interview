package heap

import (
	"log"
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
	//	log.Println(h.Pop())
	//}
	for i := 0; i < n; i++ {
		ans[i], _ = h.Pop()
		//log.Println(ans[i])
	}
	log.Println("correct:", correct)
	log.Println("ans:", ans)
	sort.Ints(correct)
	log.Println(correct)
	log.Println("test:", reflect.DeepEqual(correct, ans))

}
