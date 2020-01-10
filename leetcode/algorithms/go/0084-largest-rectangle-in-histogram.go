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

// 向右遍歷 若非最後一個高度 遇到高度非遞減的情況 不計算面積 直接向右
// 否則視為高度遞減 向左計算面積 更新面積 再繼續向右

// [2, 1, 5, 6, 2, 3]

// 2
// 6, 10, 3, 4
// 3, 4,  6, 8, 5, 6

// incomprehension
// func largestRectangleArea(heights []int) int {
// 	var result = 0
// 	var stack = []int{}
// 	heights = append(heights, 0)
// 	for i := 0; i < len(heights); i++ {
// 		for len(stack) > 0 && heights[i] <= heights[stack[len(stack)-1]] {
// 			mark := stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]
// 			area := 0
// 			if len(stack) == 0 {
// 				area = heights[mark] * i
// 			} else {
// 				area = heights[mark] * (i - stack[len(stack)-1] - 1)
// 			}
// 			if result < area {
// 				result = area
// 			}
// 		}
// 		stack = append(stack, i)
// 	}
// 	return result
// }
