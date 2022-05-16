package leetcode

/*
There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]). The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or right at any point in time.

Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The test cases are generated so that the answer will be less than or equal to 2 * 109.
*/

func uniquePaths(m int, n int) int {
	// 为什么这里要选更小的那个数作为 m（假设 n 在下，m 在上）呢？其实跟我们 computePermutation 的算法有关。
	// 因为我们是除法和乘法同步进行的，而这有可能会导致我们的 res 会变成0，从而得到0这个答案。而如果我们从小的那个数字开始计算，则不会有这个问题。
	// 当然，这里其实也可以更改计算方法，将二者的计算分开进行。
	var smaller int
	if m < n {
		smaller = m
	} else {
		smaller = n
	}
	return computePermutation(smaller-1, m+n-2)
}

// 解法1： 排列组合
func computePermutation(above, under int) int {
	res := 1
	for i := 0; i < above; i++ {
		res *= (under - i)
		res /= (i + 1)
	}

	return res

	// 这里需要注意的是，对于above 和 under 都等于99的情况，下面的 up 和 down 都会严重超出 uint64 所能表示的范围，
	// 截断之后 up 和 down 都会变成0，从而得到错误的结算结果。
	// var (
	// 	up uint64 = 1
	// 	down uint64 = 1
	// )
	// for i := 0; i < above;i++ {
	//     up *= uint64((under-i))
	//     down *= uint64((i+1))
	// }

	// return int(up/down)
}

// 解法2： 动态规划
// dp[i][j] 表示到达[i,j]点的路径个数，因为任意一个点，要么从上方到达，要么从左边到达，所以递推方程就是 dp[i][j] = dp[i-1][j] + dp[i][j-1]
// 对于base case 的情况，可以得到 dp[0][j] = dp[i][0] = 1，因为第一行和第一列的任意一个位置只有一条路可以到达。
func uniquePathsDP(m int, n int) int {
	// initialization
	dp := make([][]int, m)
	for idx, _ := range dp {
		dp[idx] = make([]int, n)
	}

	// base case
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	// two dimensional dp
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

// 进一步的发现，dp[i][j] 的值只会从 dp[i-1][j] 和 dp[i][j-1] 两处得来，也就是说只会从侧和上侧得到，那么我们就不需要使用一个二维数组，
// 只需要使用一个一维数组即可
// func uniquePathsDPOptimized(m int, n int) int {
// 	left, up := make([]int, m), make([]int, n)
// 	for i := 0 ; i < m; i++ {
// 		left[i] = 1
// 	}

// 	for j := 0; j < n ;j++ {
// 		up[j] = 1
// 	}

// 	for i := 1; i < m; i++ {
// 		for j := 1; j < n; j++ {
// 			up[j] = up[j-1] + left[j]
// 		}
// 	}
// }

// for test
func ComputePermutation(above, under int) int {
	return computePermutation(above, under)
}
