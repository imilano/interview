package msort

// SelectSort sort by comparation
func SelectSort(arr []int) {
	size := len(arr)
	for i := 0; i < size-1; i++ {
		min := i
		for j := i + 1; j < size; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
	}
}
