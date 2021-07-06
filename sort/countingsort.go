package sort

func CountingSort(arr []int) {
	if arr == nil || len(arr) <= 1 {
		return
	}
	max, min := arr[0], arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	// 建立长度为 range 的数组，下标 0~range-1 对应数字 min~max
	Len := max - min + 1
	bucket := make([]int, Len)
	for _, v := range arr {
		bucket[v-min]++
	}
	result := make([]int, len(arr))
	index := 0
	for k, v := range bucket {
		value := k + min
		for j := v; j > 0; j-- {
			result[index] = value
			index++
		}
	}
	copy(arr, result)
}
