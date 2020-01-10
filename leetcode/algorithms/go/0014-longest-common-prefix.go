package leetcode

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			if len(prefix) == 0 {
				return ""
			}
		}
	}

	return prefix
}

// func longestCommonPrefix(strs []string) string {
// 	var result string

//     if len(strs) == 0 {
//         return result
//     }

//     str := strs[0]
//     for i := 0; i < len(str); i++ {
//         char := string(str[i])
//         for j := 1; j < len(strs); j++ {
//             if i == len(strs[j]) || char != string(strs[j][i]) {
//                 return result
//             }
//         }
//         result += char
//     }

//     return result
// }
