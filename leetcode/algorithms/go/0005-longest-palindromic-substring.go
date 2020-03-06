package leetcode

func longestPalindrome(s string) string {
	var start, length int
	var helper = func(i, j int, start, length *int) {
		for -1 < i && j < len(s) && s[i] == s[j] {
			i--
			j++
		}
		if *length < j-i-1 {
			*start = i + 1
			*length = j - i - 1
		}
	}
	for i := range s {
		helper(i, i, &start, &length)
		helper(i, i+1, &start, &length)
	}
	return s[start : start+length]
}

// func longestPalindrome(s string) string {
// 	var start, length int
// 	var helper = func(i, j, start, length int) (int, int) {
// 		for -1 < i && j < len(s) && s[i] == s[j] {
// 			i--
// 			j++
// 		}
// 		if length < j-i-1 {
// 			start = i + 1
// 			length = j - i - 1
// 		}
// 		return start, length
// 	}
// 	for i := range s {
// 		start, length = helper(i, i, start, length)
// 		start, length = helper(i, i+1, start, length)
// 	}
// 	return s[start : start+length]
// }
