package leetcode

/*
You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
*/

// 左右对折翻转一次，然后再按副对角线对折翻转一次
func rotate(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])

	// 左右对折
	for i := 0; i < m; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-1-j], matrix[i][j]
		}
	}

	// 副对角线对折
	for i := 0; i < m; i++ {
		for j := 0; j < n-i-1; j++ {
			matrix[i][j], matrix[m-j-1][m-i-1] = matrix[m-j-1][m-i-1], matrix[i][j]
		}
	}
}
