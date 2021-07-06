package sort

import (
	"math/rand"
	"sort"
	"testing"
)

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}
func TestMergeSort(t *testing.T) {
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
		MergeSort(test[i], 0, len(test[i])-1)
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
		}
	}

}
