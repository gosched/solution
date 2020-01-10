package leetcode

func lengthOfLastWord(s string) int {
	count, r := 0, len(s)-1

	for r >= 0 && s[r] == ' ' {
		r--
	}

	for r >= 0 && s[r] != ' ' {
		count++
		r--
	}

	return count
}
