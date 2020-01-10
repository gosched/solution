package leetcode

func minWindow(s string, t string) string {
	need := make([]int, 128, 128)
	for i := range t {
		need[t[i]]++
	}

	find, total := 0, len(t)
	start, length := -1, len(s)+1

	i := 0
	for j := 0; j < len(s); j++ {
		need[s[j]]--
		if need[s[j]] >= 0 {
			find++
		}
		for find == total {
			if j-i+1 < length {
				start = i
				length = j - i + 1
			}
			need[s[i]]++
			if need[s[i]] > 0 {
				find--
			}
			i++
		}
	}

	if start == -1 {
		return ""
	}

	return s[start : start+length]
}

// func minWindow(s string, t string) string {
// 	need := map[byte]int{}
// 	for i := range t {
// 		need[t[i]]++
// 	}
// 	find, total := 0, len(t)
// 	start, length := -1, len(s)+1
// 	i := 0
// 	for j := 0; j < len(s); j++ {
// 		need[s[j]]--
// 		if need[s[j]] >= 0 {
// 			find++
// 		}
// 		for find == total {
// 			if j-i+1 < length {
// 				start = i
// 				length = j - i + 1
// 			}
// 			need[s[i]]++
// 			if need[s[i]] > 0 {
// 				find--
// 			}
// 			i++
// 		}
// 	}
// 	if start == -1 {
// 		return ""
// 	}
// 	return s[start : start+length]
// }

// s[start:start+length] == window

// i == window's left index
// j == window's right index
// j-i+1 == window's length

// start == minimum window's left index
// length == minimum window's length

// if find == total == len(t)
// then window contain all the characters in T

// A 65 0 1
// D 68 0 0
// O 79 0 0
// B 66 0 1
// E 69 0 0
// C 67 0 1
// O 79 1 0
// D 68 1 0
// E 69 1 0
// B 66 1 1
// A 65 1 1
// N 78 0 0
// C 67 1 1
