package sort

func SelectionSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	for i := len(arr) - 1; i > 0; i-- {
		pos := i
		for j := i - 1; j >= 0 ; j--{
			if arr[pos] < arr[j] {
				pos = j
			}
		}
		if pos != i {
			arr[pos], arr[i] = arr[i], arr[pos]
		}
	}

}
