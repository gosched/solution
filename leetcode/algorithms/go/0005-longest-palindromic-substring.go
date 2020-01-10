package leetcode

func longestPalindrome(s string) string {
	var start, length = 0, 0

	helper := func(i, j int, start, length *int) {
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
