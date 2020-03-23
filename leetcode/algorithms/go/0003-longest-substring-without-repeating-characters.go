package leetcode

func lengthOfLongestSubstring(s string) int {
	var length = 0
	var beforeStart = -1
	var m = make(map[rune]int)
	for i, r := range s {
		if j, exist := m[r]; exist && beforeStart < j {
			beforeStart = j
		}
		m[r] = i
		if length < i-beforeStart {
			length = i - beforeStart
		}
	}
	return length
}

// beforeStart == 當前子字串的前一個索引
// m[r] == 當前字符最後一次出現的位置
// exist 曾經出現過
