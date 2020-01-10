package leetcode

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	mark := make(map[[26]byte]int)

	for _, str := range strs {
		key := [26]byte{}
		for _, r := range str {
			key[r-'a']++
		}

		i, exist := mark[key]
		if exist {
			result[i] = append(result[i], str)
		} else {
			result = append(result, []string{str})
			mark[key] = len(mark)
		}
	}

	return result
}

//  mark := map[[26]byte]int{}
//  mark := make(map[[26]byte]int)

// func groupAnagrams(strs []string) [][]string {
//     var result = [][]string{}
//     var m = make(map[[26]byte]int)
//     for _, str := range strs {
//         key := [26]byte{}
//         for _, r := range str {
//             key[r-'a']++
//         }
//         index, exist := m[key]
//         if exist {
//             result[index] = append(result[index], str)
//         } else {
//             result = append(result, []string{str})
//             m[key] = len(result) - 1
//         }
//     }
//     return result
// }

// func groupAnagrams(strs []string) [][]string {
// 	result := [][]string{}
// 	mark := make(map[string]int)
// 	for _, str := range strs {
// 		temp := [26]byte{}
// 		for _, r := range str {
// 			temp[r-'a']++
// 		}
// 		key := fmt.Sprintf("%x", temp)
// 		i, ok := mark[key]
// 		if ok {
// 			result[i] = append(result[i], str)
// 		} else {
// 			result = append(result, []string{str})
// 			mark[key] = len(mark)
// 		}
// 	}
// 	return result
// }

// func groupAnagrams(strs []string) [][]string {
// 	result := [][]string{}
// 	mark := make(map[string]int)
// 	for _, str := range strs {
// 		var slice = strings.Split(str, "")
// 		sort.Strings(slice)
// 		var key = strings.Join(slice, "")
// 		i, ok := mark[key]
// 		if ok {
// 			result[i] = append(result[i], str)
// 		} else {
// 			result = append(result, []string{str})
// 			mark[key] = len(mark)
// 		}
// 	}
// 	return result
// }
