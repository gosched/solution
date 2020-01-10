package leetcode

func letterCombinations(digits string) []string {
	var m = map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	var helper func(digits string) []string

	helper = func(digits string) []string {
		var combinations []string

		if len(digits) == 0 {
			return combinations
		}

		if len(digits) == 1 {
			return m[digits[0]]
		}

		for _, letter := range m[digits[0]] {
			for _, suffix := range helper(digits[1:]) {
                combinations = append(combinations, letter + suffix)
			}
		}

		return combinations
	}

	return helper(digits)
}

// func letterCombinations(digits string) []string {
// 	result := []string{}

// 	if len(digits) == 0 {
// 		return result
// 	}

// 	dict := make(map[rune]string)

// 	dict[50] = "abc"
// 	dict[51] = "def"
// 	dict[52] = "ghi"
// 	dict[53] = "jkl"
// 	dict[54] = "mno"
// 	dict[55] = "pqrs"
// 	dict[56] = "tuv"
// 	dict[57] = "wxyz"

// 	result = append(result, "")

// 	for _, digit := range digits {
// 		str := dict[digit]
// 		temp := []string{}
// 		for _, value := range result {
// 			for _, r := range str {
// 				temp = append(temp, value+string(r))
// 			}
// 		}
// 		result = temp
// 	}

// 	return result
// }
