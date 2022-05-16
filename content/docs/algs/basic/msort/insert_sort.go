package msort

// InsertSort sort by comparation
func InsertSort(arr []int) {
	size := len(arr)
	for i := 1; i < size; i++ {
		tmp := arr[i]
		j := i - 1

		// 注意比较条件
		// 注意不能是 arr[i] < arr[j]， 这样会出问题
		for j >= 0 && tmp < arr[j] {
			arr[j+1] = arr[j]
			j--
		}

		// 注意填充位置
		arr[j+1] = tmp
	}
}
