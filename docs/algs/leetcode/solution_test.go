package leetcode_test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"blog.lightsinger.top/interview/leetcode"
)

// leetcode 1
func TestTwoSum(t *testing.T) {
	nums := []int{1, 7, 2, 4}
	target := 3

	leetcode.ExplicitTwoSum(nums, target)
}

// leetcode 5
func TestLongestPalindromic(t *testing.T) {
	strs := []string{"cbbd"}
	results := []string{"bb"}

	for index, s := range strs {
		res := leetcode.LongestPalindromeSubstring(s)
		if res != results[index] {
			t.Fail()
		}
	}
}

// leetcode 7
func TestReverseInteger(t *testing.T) {
	nums := []int{1534236469}
	results := []int{0}

	for index, value := range nums {
		res := leetcode.Reverse(value)
		if res != results[index] {
			t.Fail()
		}
	}
}

// leetcode 121
func TestMaxProfit(t *testing.T) {
	prices := [][]int{
		{3, 2, 6, 5, 0, 3},
		{7, 1, 5, 3, 6, 4},
		{7, 6, 4, 3, 1},
		{1, 2},
	}

	ans := []int{
		4,
		5,
		0,
		1,
	}

	for index, price := range prices {
		p := leetcode.MaxProfit(price)
		if p != ans[index] {
			t.Fail()
		}
	}
}

// leetcode 238
func TestProductExceptSelf(t *testing.T) {
	nums_arr := [][]int{
		{1, 2, 3, 4},
		{-1, 1, 0, -3, 3},
	}

	ans_arr := [][]int{
		{24, 12, 8, 6},
		{0, 0, 9, 0, 0},
	}

	for index, arr := range nums_arr {
		ans := leetcode.ProductExceptSelf(arr)
		if !reflect.DeepEqual(ans, ans_arr[index]) {
			t.Fatalf("fail in index %d, want %v, get %v", index, ans_arr[index], ans)
		}
	}
}

// leetcode 153
func TestFindMin(t *testing.T) {
	numsArr := [][]int{
		{3, 4, 5, 1, 2},
		{4, 5, 6, 7, 0, 1, 2},
		{11, 13, 15, 17},
	}

	ans := []int{1, 0, 11}
	for index, nums := range numsArr {
		r := leetcode.FindMin(nums)
		if r != ans[index] {
			t.Fatalf("fail in index %d, want %v, get %v", index, ans[index], r)
		}
	}
}

// leetcode 33
func TestSearch(t *testing.T) {
	numsArr := [][]int{
		{4, 5, 6, 7, 0, 1, 2},
		{4, 5, 6, 7, 0, 1, 2},
		{1},
	}

	targets := []int{0, 3, 0}
	ans := []int{4, -1, -1}
	for index, nums := range numsArr {
		r := leetcode.Search(nums, targets[index])
		if r != ans[index] {
			t.Errorf("fail in index %d, want %v, get %v", index, ans[index], r)
		}
	}
}

// leetcode 15
func TestThreeSum(t *testing.T) {
	numArr := [][]int{
		{1, -1, -1, 0},
		{-1, 0, 1, 2, -1, -4},
		{},
		{0},
	}

	for _, value := range numArr {
		ans := leetcode.ThreeSum(value)
		fmt.Printf("%v", ans)
	}
}

//leetcode 516
func TestLongestPalindromicSubseq(t *testing.T) {
	strs := []string{
		"bbbab",
		"cbbd",
	}

	ans := []int{4, 2}

	for index, s := range strs {
		r := leetcode.LongestPalindromeSubseq(s)
		if r != ans[index] {
			t.Fail()
		}
	}
}

// leetcode 121
func TestWordBreak(t *testing.T) {
	strs := []string{
		"leetcode",
		"applepenapple",
		"catsandog",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
	}

	wordDcits := [][]string{
		{"leet", "code"},
		{"apple", "pen"},
		{"cats", "dog", "sand", "and", "cat"},
		{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
	}

	ans := []bool{
		true,
		true,
		false,
		false,
	}

	for index, s := range strs {
		r := leetcode.WordBreak(s, wordDcits[index])
		if r != ans[index] {
			t.Fatalf("index %d, want %v, get %v", index, ans[index], r)
		}
	}
}

// leetcode 39
func TestCombinationSum(t *testing.T) {
	candidates := [][]int{
		{2, 3, 6, 7},
		{2, 3, 5},
		{2},
	}

	targets := []int{
		7, 8, 1,
	}

	for index, candidate := range candidates {
		ans := leetcode.CombinationSum(candidate, targets[index])
		fmt.Println(ans)
	}
}

//leetcode 40
func TestCombinationSum2(t *testing.T) {
	candidates := [][]int{
		{1, 2},
		{10, 1, 2, 7, 6, 1, 5},
		{2, 5, 2, 1, 2},
		{1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}

	targets := []int{
		3,
		8,
		5,
		1,
		27,
	}

	for index, candcandidate := range candidates {
		ans := leetcode.CombinationSum2(candcandidate, targets[index])
		fmt.Println(ans)
	}
}

//leetcode 216
func TestCombinationSum3(t *testing.T) {
	ks := []int{3, 3, 4}
	ns := []int{7, 9, 1}

	for index, k := range ks {
		ans := leetcode.CombinationSum3(k, ns[index])
		fmt.Println(ans)
	}
}

// leetcode 377
func TestCombinationSum4(t *testing.T) {
	nums := [][]int{
		{1, 2, 3},
		{9},
	}

	targets := []int{4, 0}
	ans := []int{7, 0}
	for index, arr := range nums {
		r := leetcode.CombinationSum4(arr, targets[index])
		if r != ans[index] {
			t.Fatalf("index %v, want %v, get %v", index, ans[index], r)
		}
	}
}

// leetcode 77
func TestCombine(t *testing.T) {
	ns := []int{4, 1}
	ks := []int{2, 1}
	for index, n := range ns {
		r := leetcode.Combine(n, ks[index])
		fmt.Println(r)
	}
}

//leetcode 91
func TestNumDecodings(t *testing.T) {
	strs := []string{"12", "226", "06"}
	ans := []int{2, 3, 0}
	for index, str := range strs {
		r := leetcode.NumDecodings(str)
		if r != ans[index] {
			t.Errorf("index %d, want %v, get %v", index, ans[index], r)
		}
	}
}

// leetcode 62
func TestUniquePath(t *testing.T) {
	m, n := 100, 1
	fmt.Println(leetcode.ComputePermutation(m-1, m+n-2))
}

// leetcode 252
func TestCanAttendMeetngs(t *testing.T) {
	nums := [][][]int{
		{
			{0, 30},
			{5, 10},
			{15, 20},
		},
		{
			{7, 10},
			{2, 4},
		},
	}

	ans := []bool{false, true}

	for index, value := range nums {
		r := leetcode.CanAttendMeetings(value)
		if r != ans[index] {
			t.Errorf("index %d, want %v, get %v", index, ans[index], r)
		}
	}
}

// leetcode 253
func TestMinMeetingRoom(t *testing.T) {
	intervals := [][][]int{
		{
			{0, 30},
			{5, 10},
			{15, 20},
		},
		{
			{7, 10},
			{2, 4},
		},
	}

	ans := []int{2, 1}
	for index, interval := range intervals {
		r := leetcode.MinMeetingRoom(interval)
		if r != ans[index] {
			t.Errorf("index %d, want %v, get %v", index, ans[index], r)
		}
	}
}

// leetcode 79 word search
func TestWordSearch(t *testing.T) {
	boards := [][][]byte{
		{
			{'A', 'A', 'A', 'A', 'A', 'A'},
			{'A', 'A', 'A', 'A', 'A', 'A'},
			{'A', 'A', 'A', 'A', 'A', 'A'},
			{'A', 'A', 'A', 'A', 'A', 'A'},
			{'A', 'A', 'A', 'A', 'A', 'A'},
		},
	}

	words := []string{
		"AAAAAAAAAAAABAA",
	}
	ans := []bool{
		false,
	}
	for index, board := range boards {
		r := leetcode.WordSearch(board, words[index])
		if r != ans[index] {
			t.Errorf("index %d , want %v, get %v", index, ans[index], r)
		}
	}
}

// 20. valid parentheses
func TestValidParentheses(t *testing.T) {
	inputs := []string{
		"()",
		"()[]{}",
		"(]",
	}

	ans := []bool{true, true, false}

	for index, input := range inputs {
		r := leetcode.ValidParentheses(input)
		if r != ans[index] {
			t.Errorf("index %d, want %v, get %v", index, ans[index], r)
		}
	}
}

// test type of element for string
func TestString(t *testing.T) {
	str := "a"
	for i, s := range str {
		fmt.Printf("type for str[i]: %T\n", str[i])
		fmt.Printf("type for 'a': %T\n", 'a')
		fmt.Printf("type for s: %T\n", s)
	}
}

// leetcode 8, string to integer
func TestMyAtoI(t *testing.T) {
	strs := []string{"42", "4193 with words", "-91283472332"}
	res := []int{42, 4193, -2147483648}

	size := len(strs)
	for i := 0; i < size; i++ {
		r := leetcode.MyAtoi(strs[i])
		if r != res[i] {
			t.Errorf("index %d, want %d, get %d", i, res[i], r)
		}
	}
}

func TestCountComponents(t *testing.T) {
	ns := []int{5, 5}
	edges := [][][]int{
		{
			{0, 1},
			{1, 2},
			{3, 4},
		},
		{
			{0, 1},
			{1, 2},
			{2, 3},
			{3, 4},
		},
	}
	ans := []int{2, 1}

	for idx, _ := range ns {
		r := leetcode.CountComponents(ns[idx], edges[idx])
		if r != ans[idx] {
			t.Errorf("index %d, want %d, get %d", idx, ans[idx], r)
		}
	}

}

func TestFindWords(t *testing.T) {
	board := [][]byte{
		{'a', 'b', 'c'},
		{'a', 'e', 'd'},
		{'a', 'f', 'g'},
	}

	words := []string{"eaabcdgfa"}
	leetcode.FindWordsII(board, words)
}


func TestIOTA(t *testing.T) {
	const (
		a = 1 << iota
		b
		c
		d = iota
	)

	fmt.Println(a, b, c, d)
}


func TestSort(t *testing.T) {
	nums := [][]int{
		{1,3},
		{2,3},
		{1,2},
		{2,4},
	}

	sort.Slice(nums, func(i, j int) bool {
		if nums[i][0] == nums[j][0] {
			return nums[i][1] < nums[j][1]
		} else {
			return nums[i][0] < nums[j][0]
		}
	})

	fmt.Println(nums)
}

// test if array could be used as key in map
// arrar could not be used as key in map
// func TestMap(t *testing.T) {
// 	m := make(map[[]int]bool)
// }

func TestStringSort(t *testing.T) {
	strs := []string{"bc", "bca"}
	sort.Strings(strs)
	fmt.Println(strs)
}
