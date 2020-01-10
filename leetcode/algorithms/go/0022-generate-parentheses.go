package leetcode

func generateParenthesis(n int) []string {
	var result = []string{}
	var helper func(left, right int, output string)
	helper = func(left, right int, output string) {
		if left < 0 || right < 0 || left > right {
			return
		}
		if left == 0 && right == 0 {
			result = append(result, output)
			return
		}
		helper(left-1, right, output+"(")
		helper(left, right-1, output+")")
	}
	helper(n, n, "")
	return result
}
