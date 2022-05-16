package msort

func MergeSort(arr []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		MergeSort(arr, left, mid)
		MergeSort(arr, mid+1, right)
		Merge(arr, left, mid, right)
	}
}

func Merge(arr []int, left, mid, right int) {
	var tmp []int
	l, r := left, mid+1
	for l <= mid && r <= right {
		if arr[l] < arr[r] {
			tmp = append(tmp, arr[l])
			l++
		} else {
			tmp = append(tmp, arr[r])
			r++
		}
	}

	for l <= mid {
		tmp = append(tmp, arr[l])
		l++
	}

	for r <= right {
		tmp = append(tmp, arr[r])
		r++
	}

	for _, v := range tmp {
		arr[left] = v
		left++
	}
}
