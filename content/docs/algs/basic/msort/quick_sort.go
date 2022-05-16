package msort

func QuickSort(arr []int, left, right int) {
	if left < right {
		pivot := partition(arr, left, right)
		QuickSort(arr, left, pivot -1)
		QuickSort(arr, pivot + 1, right)
	}
}


func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[index] = arr[index], arr[i]
			index += 1
		}
	}

	arr[pivot], arr[index -1] = arr[index-1], arr[pivot]
	return index -1 
}
