package leetcode

func longestValidParentheses(s string) int {
	length := 0
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if length < left*2 {
				length = left * 2
			}
		} else if left < right {
			left = 0
			right = 0
		}
	}
	left, right = 0, 0
	for j := len(s) - 1; -1 < j; j-- {
		if s[j] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if length < left*2 {
				length = left * 2
			}
		} else if right < left {
			left = 0
			right = 0
		}
	}
	return length
}

// incomprehension
// func longestValidParentheses(s string) int {
// 	var result, stack = 0, []int{}
// 	var max = func(x, y int) int {
// 		if x < y {
// 			return y
// 		}
// 		return x
// 	}
// 	stack = append(stack, -1)
// 	for i, r := range s {
// 		if r == '(' {
// 			stack = append(stack, i)
// 		} else {
// 			stack = stack[:len(stack)-1]
// 			if len(stack) == 0 {
// 				stack = append(stack, i)
// 			} else {
// 				result = max(result, i-stack[len(stack)-1])
// 			}
// 		}
// 	}
// 	return result
// }
