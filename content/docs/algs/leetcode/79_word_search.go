package leetcode

/*
Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.
*/

// 回溯遍历即可
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				visited := make([][]bool, m)
				for i, _ := range visited {
					visited[i] = make([]bool, n)
				}
				if existBacktrace(board, word, i, j, m, n, &visited) {
					return true
				}
			}
		}
	}
	return false
}

func existBacktrace(board [][]byte, word string, si, sj, m, n int, visited *[][]bool) bool {
	if len(word) == 0 {
		return true
	}
	if si < 0 || sj < 0 || si >= m || sj >= n || (*visited)[si][sj] || board[si][sj] != word[0] {
		return false
	}
	if board[si][sj] == word[0] {
		// 不用重新建立一个 visited 数组,因为下面的语句会把 visited 数组还原
		(*visited)[si][sj] = true
		if existBacktrace(board, word[1:], si+1, sj, m, n, visited) ||
			existBacktrace(board, word[1:], si-1, sj, m, n, visited) ||
			existBacktrace(board, word[1:], si, sj+1, m, n, visited) ||
			existBacktrace(board, word[1:], si, sj-1, m, n, visited) {
			return true
		}
		(*visited)[si][sj] = false
	}

	return false
}

func copyArray(original [][]bool, m, n int) *[][]bool {
	arr := make([][]bool, m)
	for i, _ := range arr {
		arr[i] = make([]bool, n)
		copy(arr[i], original[i])
	}

	return &arr

}

// for test
func WordSearch(board [][]byte, word string) bool {
	return exist(board, word)
}
