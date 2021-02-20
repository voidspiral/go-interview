package sort

import (
	"log"
	"math/rand"
	"sort"
	"testing"
)

func TestHeapSort(t *testing.T) {
	test := make([][]int, 10)
	for i := 0; i < len(test); i++ {
		n := rand.Intn(100)
		test[i] = make([]int, n)
		for j := 0; j < n; j++ {
			test[i][j] = rand.Intn(100)
		}
		//log.Println(test[i])
	}
	elems := make([][]int, len(test))
	for i := 0; i < len(test); i++ {
		elems[i] = make([]int, len(test[i]))
		copy(elems[i], test[i])
		sort.Slice(elems[i], func(a, b int) bool { return elems[i][a] < elems[i][b] })

		//log.Println(elems[i])
	}

	for i := 0; i < len(test); i++ {
		log.Println("before:", test[i])
		HeapSort(test[i])
		isSorted := sort.SliceIsSorted(test[i], func(a int, b int) bool { return test[i][a] < test[i][b] })
		log.Println("isSorted:", isSorted)
		log.Println("after:", test[i])
		log.Println("elems:", elems[i])
		isEqual := false
		if equal(test[i], elems[i]) {
			isEqual = true
		}
		log.Println("isEqual:", isEqual)
		log.Println("")
		if !isSorted && !isEqual {
			log.Println("some case wrong")
		}
	}
}
