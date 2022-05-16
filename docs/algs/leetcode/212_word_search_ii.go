package leetcode

/*
Given an m x n board of characters and a list of strings words, return all words on the board.

Each word must be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.
*/

func findWords(board [][]byte, words []string) []string {
	var res []string
	size := len(words)
	dict := make(map[string]bool, size)
	for _, word := range words {
		dict[word] = false
	}
	m, n := len(board), len(board[0])
	for _, word := range words {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if dict[word] == false && findWordsTraversal(board, word, i, j, m, n) {
					dict[word] = true
					res = append(res, word)
				}
			}
		}
	}

	return res
}

func findWordsTraversal(board [][]byte, word string, si, sj, m, n int) bool {
	if len(word) == 0 {
		return true
	}
	if si < 0 || sj < 0 || si >= m || sj >= n || board[si][sj] == byte('.') {
		return false
	}

	if word[0] == board[si][sj] {
		// 这里很巧妙，避免了开辟额外的 visited 数组来记录访问状态
		board[si][sj] = byte('.')
		res := findWordsTraversal(board, word[1:], si+1, sj, m, n) ||
			findWordsTraversal(board, word[1:], si-1, sj, m, n) ||
			findWordsTraversal(board, word[1:], si, sj+1, m, n) ||
			findWordsTraversal(board, word[1:], si, sj-1, m, n)
		board[si][sj] = word[0]
		return res
	}

	return false
}

// for test
func FindWordsII(board [][]byte, words []string) []string {
	return findWords(board, words)
}

// var cords [][]int = [][]int{
// 	{0, 1},
// 	{0, -1},
// 	{1, 0},
// 	{-1, 0},
// }

// func findWords(board [][]byte, words []string) []string {
// 	var res []string
// 	set := make(map[string]bool)
// 	m, n := len(board), len(board[0])
// 	dict := make(map[byte][][]int)
// 	buildMap(board, &dict, m, n)
// 	for _, word := range words {
// 		points, ok := dict[word[0]]
// 		if !ok {
// 			continue
// 		}

// 		for _, point := range points {
// 			if _, ok := set[word]; ok {
// 				continue
// 			}
// 			visited := make([][]bool, m)
// 			for idx, _ := range visited {
// 				visited[idx] = make([]bool, n)
// 			}
// 			if searchBoard(board, word, &visited, point[0], point[1], m, n) {

// 				res = append(res, word)
// 				set[word] = true
// 			}
// 		}

// 	}
// 	return res
// }

// func searchBoard(board [][]byte, word string, visited *[][]bool, x, y, m, n int) bool {
// 	size := len(word)
// 	if size == 0 {
// 		return true
// 	}
// 	if x < 0 || x >= m || y < 0 || y >= n || word[0] != board[x][y] || (*visited)[x][y] {
// 		return false
// 	}

// 	var res bool
// 	for _, cord := range cords {
// 		(*visited)[x][y] = true
// 		res = res || searchBoard(board, word[1:], visited, x+cord[0], y+cord[1], m, n)
// 		(*visited)[x][y] = false
// 	}
// 	return res
// 	// 不能像下面这么写，因为四个搜索之间应该是相互独立的，但是因为 visited 数组赋值在前面，会发生前一次搜索的 visited 状态保留到了下一次搜索，影响最终结果
// 	// (*visited)[x][y] = true
// 	// return searchBoard(board, word[1:], visited, x+1, y, m, n) || searchBoard(board, word[1:], visited,  x-1, y, m, n) || searchBoard(board, word[1:], visited, x, y+1, m, n) || searchBoard(board, word[1:], visited, x, y-1, m, n)
// }

// func buildMap(board [][]byte, dict *map[byte][][]int, m, n int) {
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			(*dict)[board[i][j]] = append((*dict)[board[i][j]], []int{i, j})
// 		}
// 	}
// }
