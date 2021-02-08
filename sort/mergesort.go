package sort

//[l, r]
func MergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	MergeSort(arr, l, mid)
	MergeSort(arr, mid+1, r)
	merge(arr, l, mid, r)
}
func merge(arr []int, l, mid int, r int) {
	i, j, k := l, mid+1, 0
	tmp := make([]int, r-l+1)
	//[i,mid], [mid+1, r]
	for ; i <= mid && j <= r; k++ {
		if arr[i] < arr[j] {
			tmp[k] = arr[i]
			i++
		} else {
			tmp[k] = arr[j]
			j++
		}
	}
	for i <= mid {
		tmp[k] = arr[i]
		k++
		i++
	}
	for j <= r {
		tmp[k] = arr[j]
		k++
		j++
	}
	//copy(arr[l:r], tmp)
	x := 0
	for i := l; i <= r; i++ {
		arr[i] = tmp[x]
		x++
	}
}
