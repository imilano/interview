package msort

func HeapSort(arr []int) {
	size := len(arr)
	buildHeap(arr, size)
	for i := size - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		size--
		// 注意，因为经过建堆之后，堆已经是完全有序了
		// 经过交换之后，不需要再次进行建堆，因为最大值还是会在根节点的左右子节点中，所以只需要对从根节点开始进行调整即可。
		heapify(arr, 0, size)
	}
}


func buildHeap(arr []int, size int) {
	for i := size/2; i >= 0; i-- {
		heapify(arr, i, size)
	}
}


func heapify(arr []int, cur, size int) {
	left, right := cur *2 + 1, cur * 2 + 2
	largest := cur
	
	if left < size && arr[left] > arr[largest] {
		largest = left
	}
	if right < size && arr[right] > arr[largest] {
		largest = right
	}

	if largest != cur {
		arr[largest], arr[cur] = arr[cur], arr[largest]
		heapify(arr, largest, size)
	}
} 
