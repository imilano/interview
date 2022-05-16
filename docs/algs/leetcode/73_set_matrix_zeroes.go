package leetcode

/*
Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.

You must do it in place.
*/

func setZeroes(matrix [][]int) {
	setZeroesSolution2(matrix)
}

// 问题在于扫描时候，会把原来1的位置设置为0，但是要注意，新设置的0并不是原有的0，新位置的0并不代表要把该行该列也要设置为0
// 最简单的办法，先扫描一遍，把要改变的位置记录下来
func setZeroesSolution1(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	var dots [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				dots = append(dots, []int{i, j})
			}
		}
	}

	for _, dot := range dots {
		for i := 0; i < n; i++ {
			matrix[dot[0]][i] = 0
		}

		for j := 0; j < m; j++ {
			matrix[j][dot[1]] = 0
		}
	}
}

// 因为要求常数级的空间复杂度，所以这里我们可以直接复用原数组，使用数组的第一行第一列来记录各行各列是否有0。
// 首先，需要先扫描第一行和第一列，记录一下这一行/列本身是否存在0
// 然后扫描除第一行第一列的整个数组，如果有0，则将对应的第一行和第一列的数字赋0。
// 再次遍历除去第一行第一列的整个数组，如果对应的第一行或者第一列的数字有一个为0，则将当前值赋0
// 最后，根据第一行和第一列的扫描结果来更新第一行和第一列。
func setZeroesSolution2(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])

	// 记录第一行第一列中是否出现0
	var rowZero, colZero bool
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			colZero = true
			break
		}
	}

	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			rowZero = true
			break
		}
	}

	//  统计其余的行和列是否出现0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 根据第一行第一列统计的值设置其余行/列
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 根据第一行第一列的扫描结果设置第一行第一列
	if rowZero {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}

	if colZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
