package msort

func ShellSort(arr []int) {
	size := len(arr)

	for gap := size / 2; gap > 0; gap /= 2 {
		for i := gap; i < size; i++ {
			j := i
			// 这里可以使用基于交换的排序，也可以使用基于插入的排序
			// 另外需要注意的一点是，其实这里的目的并不是要子数组完全有序，只需要部分有序即可。
			// 从算法的实现也可以看出来，是做不到完全有序的。
			for j-gap >= 0 && arr[j] < arr[j-gap] {
				arr[j], arr[j-gap] = arr[j-gap], arr[j]
				j = j - gap
			}
		}
	}
}
