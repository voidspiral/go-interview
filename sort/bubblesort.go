package sort

//func BubbleSort(arr []int) {
//
//	if arr == nil || len(arr) < 2 {
//		fmt.Println("数组不满足要求")
//		return
//	}
//
//	// 外层循环：确定扫描的次数
//	for i := 1; i <= len(arr) - 1; i++ {
//		// 内层循环：一轮扫描内，两两比较，进行交换
//		for j := 0; j <= len(arr) - 1 - i; j++ {	 // - i 的原因是后面的元素已经被排序过
//			if arr[j] > arr[j + 1] {
//				temp := arr[j]
//				arr[j] = arr[j + 1]
//				arr[j + 1] = temp
//			}
//		}
//	}
//}
func BubbleSort(arr []int) {
	//arr := *ar
	if len(arr) == 0 {
		return
	}
	isSorted := true
	for i := len(arr) - 1; i > 0; i--{
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
				isSorted = false
			}
		}
		if isSorted{
			break
		}
	}
}
