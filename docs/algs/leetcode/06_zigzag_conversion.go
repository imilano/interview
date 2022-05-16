package leetcode

/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
*/
func convert(s string, numRows int) string {
	size := len(s)
	if size <= numRows || numRows == 1 {
		return s
	}

	curRow, down := 0, false
	arr := make([][]rune, numRows)
	for _, r := range s {
		arr[curRow] = append(arr[curRow], r)
		if curRow == 0 || curRow == numRows-1 {
			down = !down
		}

		if down {
			curRow++
		} else {
			curRow--
		}
	}

	var res string
	for _, row := range arr {
		res += string(row)
	}

	return res
}
