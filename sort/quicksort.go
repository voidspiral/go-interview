package sort

import "math/rand"

func Partition(arr []int, lo, hi int) int {
	arr[lo], arr[lo+int(rand.Int63())%(hi-lo)] = arr[lo+int(rand.Int63())%(hi-lo)], arr[lo]
	hi--
	pivot := arr[lo]
	for lo < hi {
		for lo < hi && pivot <= arr[hi] {
			hi--
		}
		arr[lo] = arr[hi]
		for lo < hi && pivot >= arr[lo] {
			lo++
		}
		arr[hi] = arr[lo]
	}
	arr[lo] = pivot
	return lo
}

//区间是[,)
func QuickSort(arr []int, l, r int) {
	if r-l < 2 {
		return
	}
	mid := Partition(arr, l, r)
	QuickSort(arr, l, mid)
	QuickSort(arr, mid+1, r)
}
