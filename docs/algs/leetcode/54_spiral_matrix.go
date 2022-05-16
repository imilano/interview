package leetcode

/*
Given an m x n matrix, return all elements of the matrix in spiral order.
*/

func spiralOrder(matrix [][]int) []int {
	var res []int
	m, n := len(matrix), len(matrix[0])
	up, down, left, right := 0, m-1, 0, n-1
	for {
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		// 避免上下重叠
		up++
		if up > down {
			break
		}

		for i := up; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		// 避免左右重叠
		right--
		if right < left {
			break
		}
		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}

		// 避免上下重叠
		down--
		if down < up {
			break
		}
		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}

		//避免左右重叠
		left++
		if left > right {
			break
		}
	}

	return res
}

// wrong first sight
func spiralOrderOriginal(matrix [][]int) []int {
	var res []int
	m, n := len(matrix), len(matrix[0])
	lm, ln, rm, rn := 0, 0, m-1, n-1
	for lm <= rm && ln <= rn {
		for i := ln; i <= rn; i++ {
			res = append(res, matrix[lm][i])
		}

		for i := lm + 1; i <= rm; i++ {
			res = append(res, matrix[i][rn])
		}

		for i := rn - 1; i >= ln; i-- {
			res = append(res, matrix[rm][i])
		}
		for i := rm - 1; i >= lm+1; i-- {
			res = append(res, matrix[i][ln])
		}

		lm++
		ln++
		rm--
		rn--
	}

	return res
}
