package sort

func sswap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
func HeSort(arr []int) {
	//构建初始大顶堆
	for i := len(arr)/2 - 1; i >= 0; i-- {
		Heapify(arr, i, len(arr)-1)
	}

	for i := len(arr) - 1; i > 0; i-- {
		sswap(arr, 0, i)
		Heapify(arr, 0, i-1)
	}
}
func Heapify(arr []int, start, end int) {
	dad := start
	son := dad*2 + 1
	for son <= end {
		if son+1 <= end && arr[son] < arr[son+1] {
			son++
		}
		if arr[dad] > arr[son] {
			return
		} else {
			swap(arr, dad, son)
			dad = son
			son = dad*2 + 1
		}
	}
}
