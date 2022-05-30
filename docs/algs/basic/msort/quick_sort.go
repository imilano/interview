package msort

func QuickSort(arr []int, left, right int) {
	if left < right {
		pivot := partition(arr, left, right)
		QuickSort(arr, left, pivot -1)
		QuickSort(arr, pivot + 1, right)
	}
}


// 快速选择算法
func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1
	// 把比 pivot 小的元素都放在左边，也就是说，除了第一个 pivot，index 左边的元素都比 pivot 小（不包括 index）
	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[index] = arr[index], arr[i]
			index += 1
		}
	}

	// 把 index 的前一个（比 pivot 小）和 pivot 交换，这样 pivot 左边的元素都比 pivot 要小了
	arr[pivot], arr[index -1] = arr[index-1], arr[pivot]
	return index -1 
}
