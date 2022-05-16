package leetcode

/*
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
*/
func climbStairs(n int) int {
	// 注意这里赋值的顺序
	FMinusTwo, FMinusOne := 0, 1
	for i := 1; i <= n; i++ {
		FN := FMinusOne + FMinusTwo
		FMinusTwo = FMinusOne
		FMinusOne = FN
	}

	return FMinusOne
}
