package leetcode

/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
*/

// 注意下面这种思路其实是错的，为什么呢？因为下面的考量其实基于一个前提，那就是 s 内部是不可再分的，也就是说，你不能往 s 内部去添加（）以构建新的有效的括号对。
// 这个假设是错误的，当 n >= 4 时，下面的代码就会得出错误的结果。
// 任意一个有效的()字符串 s ，要放置一对 () 让其有效的话，只有三种放法：
// - 全放在最后边，也就是 "s()"
// - 全放在最左边，也就是 "()s"
// - 左右各放一个，也就是"(s)"
// 那么对于任意一个字符串，枚举上面三种情况。需要注意的是，由于第一种和第二种有可能重复，所以需要用一个 map 来去重
// func generateParenthesis(n int) []string {
//     var res []string
//     dict := make(map[string]bool)
//     if n == 0 {
//         return res
//     }   
    
//     generateParenthesisHelper(n, "", &dict)
//     for key, _ := range dict {
//         res = append(res, key)
//     }
    
//     return res
// }

// func generateParenthesisHelper(n int, str string, dict *map[string]bool) {
//     if n == 0 {
//         if _, ok := (*dict)[str]; !ok {
//             (*dict)[str] = true
//         }
        
//         return
//     }
    
//     for i := 0; i< 3; i++ {
//         if i == 0 {
//             generateParenthesisHelper(n-1, str+"()", dict)
//         }
//         if i == 1 {
//             generateParenthesisHelper(n-1, "()"+str, dict)
//         }
//         if i == 2 {
//             generateParenthesisHelper(n-1, "("+str+")", dict)
//         }
//     }
// }


// 参见： https://leetcode-cn.com/problems/generate-parentheses/solution/hui-su-suan-fa-by-liweiwei1419/
// 使用两个变量 left 和 right 来表示「左括号使用了集合」和「右括号使用了几个」，通过分析可以得到以下结论：
// - 只有当left 和 right 都比 n 小的时候，才产生分支。
// - 产生左分支的时候，只看当前还有左括号可用
// - 产生右分支的时候，还受到左括号数量的限制，left 要大于 right 的时候，才可以产生分支。否则就是无效分支，可以直接返回。
// - 在左边和右边的括号数都等于 n 的时候结算。
func generateParenthesis(n int) []string {
    var res []string
	if n == 0 {
		return res
	}

	generateParenthesisHelper(n, 0,0,"", &res)
	return res
}

func generateParenthesisHelper(n, left, right int, cur string, res *[]string) {
	if left == n && right == n {
		*res = append(*res, cur)
		return
	}

	// 剪枝
	if left < right {
		return
	}

	if left < n {
		generateParenthesisHelper(n, left+1, right, cur + "(", res)
	}
	if right < n {
		generateParenthesisHelper(n, left, right+1, cur + ")", res)
	}
}
