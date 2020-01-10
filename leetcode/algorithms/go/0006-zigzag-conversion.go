package leetcode

import "strings"

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	temp := make([]string, numRows)
	row, step := 0, 1

	for _, r := range s {
		temp[row] += string(r)
		row += step
		if row == numRows {
			row = numRows - 2
			step = -1
		}
		if row < 0 {
			row = 1
			step = 1
		}
	}

	return strings.Join(temp[:], "")
}
