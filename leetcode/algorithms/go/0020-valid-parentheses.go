package leetcode

func isValid(s string) bool {
	var stack = []rune{}
	for _, r := range s {
		if r == '(' || r == '[' || r == '{' {
			stack = append(stack, r)
		} else {
			n := len(stack)
			if n == 0 {
				return false
			}
			if stack[n-1] == '(' && r != ')' {
				return false
			}
			if stack[n-1] == '[' && r != ']' {
				return false
			}
			if stack[n-1] == '{' && r != '}' {
				return false
			}
			stack = stack[:n-1]
		}
	}
	return len(stack) == 0
}
