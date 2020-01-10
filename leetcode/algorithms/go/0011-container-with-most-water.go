package leetcode

func maxArea(height []int) int {
	var result, i, j = 0, 0, len(height) - 1
	for i < j {
		h := height[i]
		if height[j] < h {
			h = height[j]
		}
		if result < h*(j-i) {
			result = h * (j - i)
		}
		if h == height[i] {
			i++
		} else {
			j--
		}
	}
	return result
}

// func maxArea(height []int) int {
// 	max := func(a, b int) int {
// 		if a > b {
// 			return a
// 		}
// 		return b
// 	}
// 	min := func(a, b int) int {
// 		if a < b {
// 			return a
// 		}
// 		return b
// 	}
// 	result, i, j := 0, 0, len(height)-1
// 	for i < j {
// 		result = max(result, (min(height[i], height[j]) * (j - i)))
// 		if height[i] < height[j] {
// 			i++
// 		} else {
// 			j--
// 		}
// 	}
// 	return result
// }

// math.Max()
// math.Min()
