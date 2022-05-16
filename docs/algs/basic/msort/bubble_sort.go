package msort

// Bubble sort sort data by comparation
func BubbleSort(arr []int) {
	size := len(arr)
	for i := size - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
