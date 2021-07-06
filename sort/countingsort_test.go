package sort

import (
	"math/rand"
	"sort"
	"testing"
)

func TestCounting(t *testing.T) {
	test := make([][]int, 10)
	for i := 0; i < len(test); i++ {
		n := rand.Intn(100)
		test[i] = make([]int, n)
		for j := 0; j < n; j++ {
			test[i][j] = rand.Intn(100)
		}
		//t.Log(test[i])
	}
	elems := make([][]int, len(test))
	for i := 0; i < len(test); i++ {
		elems[i] = make([]int, len(test[i]))
		copy(elems[i], test[i])
		sort.Slice(elems[i], func(a, b int) bool { return elems[i][a] < elems[i][b] })

		//t.Log(elems[i])
	}

	for i := 0; i < len(test); i++ {
		t.Log("before:", test[i])
		CountingSort(test[i])
		isSorted := sort.SliceIsSorted(test[i], func(a int, b int) bool { return test[i][a] < test[i][b] })
		t.Log("isSorted:", isSorted)
		t.Log("after:", test[i])
		t.Log("elems:", elems[i])
		isEqual := false
		if equal(test[i], elems[i]) {
			isEqual = true
		}
		t.Log("isEqual:", isEqual)
		t.Log("")
		if !isSorted && !isEqual {
			t.Errorf("some case wrong")
			t.Fail()
		}
	}
}
