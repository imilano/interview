package msort_test

import (
	"fmt"
	"sort"
	"testing"

	"blog.lightsinger.top/interview/leetcode/basic/msort"
)

var testCases [][]int

func init() {
	testCases = [][]int{
		{},
		{3, 2, 1},
		{44, 22, 99},
		{4, 2, 1, 3},
		{89, 31, 10, 56, 4, 96, 35, 28, 7, 89, 4, 84, 12, 39, 55, 32, 84, 96, 90, 92, 93, 91, 81, 70, 1, 10, 78, 19, 34, 10, 33, 81, 31, 67, 65, 93, 48, 7, 36, 74, 43, 92, 5, 9, 41, 100, 36, 23, 61, 70},
	}
}

func check(t *testing.T, arr []int, idx int) {
	if !sort.IntsAreSorted(arr) {
		t.Errorf("index %d, sorted: %v\n", idx, arr)
	}
	fmt.Printf("index %d pass the test\n", idx)
}

func TestBubbleSort(t *testing.T) {
	for idx, arr := range testCases {
		msort.BubbleSort(arr)
		check(t, arr, idx)
	}

}

func TestSelectSort(t *testing.T) {
	for idx, arr := range testCases {
		msort.SelectSort(arr)
		check(t, arr, idx)
	}
}

func TestInsertSort(t *testing.T) {
	for idx, arr := range testCases {
		msort.InsertSort(arr)
		check(t, arr, idx)
	}
}

func TestMergeSort(t *testing.T) {
	for idx, arr := range testCases {
		size := len(arr)
		msort.MergeSort(arr, 0, size-1)
		check(t, arr, idx)
	}
}

func TestShellSort(t *testing.T) {
	for idx, arr := range testCases {
		msort.ShellSort(arr)
		check(t, arr, idx)
	}
}

func TestQuickSort(t *testing.T) {
	for idx, arr := range testCases {
		size := len(arr)
		msort.QuickSort(arr, 0, size - 1)
		check(t, arr, idx)
	}
}


func TestHeapSort(t *testing.T) {
	for idx, arr := range testCases {
		msort.HeapSort(arr)
		check(t, arr, idx)
	}
}
