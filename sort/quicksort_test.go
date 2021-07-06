package sort

import (
	"math/rand"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	test := make([][]int, 10)
	for i := 0; i < len(test); i++ {
		n := rand.Intn(100)
		test[i] = make([]int, n)
		for j := 0; j < n; j++ {
			test[i][j] = rand.Intn(1000)
		}
	}

	for i := 0; i < len(test); i++ {
		t.Log("before:", test[i])
		QuickSort(test[i], 0, len(test[i]))
		isSorted := sort.SliceIsSorted(test[i], func(a int, b int) bool { return test[i][a] < test[i][b] })
		t.Log("isSorted:", isSorted)
		t.Log("after:", test[i])
		t.Log("")
		if !isSorted {
			t.Errorf("some case wrong")
		}
	}

}
