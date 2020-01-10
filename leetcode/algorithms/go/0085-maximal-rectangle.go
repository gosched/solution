package leetcode

func largestRectangleArea(heights []int) int {
	var result = 0
	for i := 0; i < len(heights); i++ {
		if i+1 < len(heights) && heights[i] <= heights[i+1] {
			continue
		}
		var minHeight = heights[i]
		for j := i; j >= 0; j-- {
			if heights[j] < minHeight {
				minHeight = heights[j]
			}
			if result < minHeight*(i-j+1) {
				result = minHeight * (i - j + 1)
			}
		}
	}
	return result
}

func maximalRectangle(matrix [][]byte) int {
	var result = 0
	if len(matrix) == 0 {
		return result
	}
	var max = func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	var heights = make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				heights[j] = 0
			} else {
				heights[j]++
			}
		}
		result = max(result, largestRectangleArea(heights))
	}
	return result
}

// ["1","0","1","0","0"]
// ["1","0","1","1","1"]
// ["1","1","1","1","1"]
// ["1","0","0","1","0"]

// [(1),"0",(1),"0","0"]
// largestRectangleArea([1,0,1,0,0])

// [(1),"0",(1),"0","0"]
// [(1),"0",(1),(1),(1)]
// largestRectangleArea([2,0,2,1,1])

// [(1),"0",(1),"0","0"]
// [(1),"0",(1),(1),(1)]
// [(1),(1),(1),(1),(1)]
// largestRectangleArea([3,1,3,2,2])

// [(1),"0","1","0","0"]
// [(1),"0","1",(1),"1"]
// [(1),"1","1",(1),"1"]
// [(1),"0","0",(1),"0"]
// largestRectangleArea([4,0,0,3,0])

// import "math"
// func maximalRectangle(matrix [][]byte) int {
// 	r := len(matrix)
// 	if r == 0 {
// 		return 0
// 	}
// 	c := len(matrix[0])
// 	dp := make([][]int, r, r)
// 	for i := 0; i < r; i++ {
// 		dp[i] = make([]int, c, c)
// 	}
// 	for i := 0; i < r; i++ {
// 		for j := 0; j < c; j++ {
// 			if matrix[i][j] == '1' {
// 				if j == 0 {
// 					dp[i][j] = 1
// 				} else {
// 					dp[i][j] = dp[i][j-1] + 1
// 				}
// 			} else {
// 				dp[i][j] = 0
// 			}
// 		}
// 	}
// 	min := func(x, y int) int {
// 		if x < y {
// 			return x
// 		}
// 		return y
// 	}
// 	max := func(x, y int) int {
// 		if x < y {
// 			return y
// 		}
// 		return x
// 	}
// 	result := 0
// 	for i := 0; i < r; i++ {
// 		for j := 0; j < c; j++ {
// 			length := math.MaxInt64
// 			for k := i; k < r; k++ {
// 				length = min(length, dp[k][j])
// 				if length == 0 {
// 					break
// 				}
// 				result = max(result, length*(k-i+1))
// 			}
// 		}
// 	}
// 	return result
// }

// func maximalRectangle(matrix [][]byte) int {
// 	m := len(matrix)
// 	if m == 0 {
// 		return 0
// 	}
// 	n := len(matrix[0])
// 	if n == 0 {
// 		return 0
// 	}
// 	type level struct {
// 		left, right, height int
// 	}
// 	calculateArea := func(l level) int {
// 		return l.height * (l.right - l.left + 1)
// 	}
// 	maxArea := 0
// 	levels := []level{}
// 	for i := 0; i < m; i++ {
// 		// update height
// 		for j := 0; j < n; j++ {
// 			if j >= len(levels) {
// 				l := level{right: n}
// 				if matrix[i][j] == '1' {
// 					l.height = 1
// 				}
// 				levels = append(levels, l)
// 			} else {
// 				if matrix[i][j] == '1' {
// 					levels[j].height++
// 				} else {
// 					levels[j].height = 0
// 				}
// 			}
// 		}
// 		// update left
// 		left := 0
// 		for j := 0; j < n; j++ {
// 			if matrix[i][j] == '1' {
// 				if left > levels[j].left {
// 					levels[j].left = left
// 				}
// 			} else {
// 				levels[j].left = 0
// 				left = j + 1
// 			}
// 		}
// 		// update right
// 		right := n - 1
// 		for j := n - 1; j >= 0; j-- {
// 			if matrix[i][j] == '1' {
// 				if right < levels[j].right {
// 					levels[j].right = right
// 				}
// 			} else {
// 				levels[j].right = n
// 				right = j - 1
// 			}
// 		}
// 		// update area
// 		for j := 0; j < n; j++ {
// 			area := calculateArea(levels[j])
// 			if area > maxArea {
// 				maxArea = area
// 			}
// 		}
// 	}
// 	return maxArea
// }
