package msort

// 自顶向下
func MergeSortUpToDown(arr []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		MergeSortUpToDown(arr, left, mid)
		MergeSortUpToDown(arr, mid+1, right)
		Merge(arr, left, mid, right)
	}
}

// 自底向上
func MergeSortDownToUp(arr []int) {
	size := len(arr)
	if size <= 1 {
		return
	}

	for i := 1; i < size; i *= 2 { // 间隔，就是每个子数组的长度
		for j := 0; j < size; j += 2*i { // 两两合并， j 表示第一个子数组的开始索引
			Merge(arr, j, j+i-1, min(j+2*i-1,  size-1)) // 对于长度不满足 2 的 x 次幂的数组，j+2*i-1 可能会超过数组长度，所以需要控制一下
		}
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
