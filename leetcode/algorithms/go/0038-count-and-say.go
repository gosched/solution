package leetcode

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	var result strings.Builder
	var temp = []rune(countAndSay(n - 1))

	var count = 1
	var prevChar = temp[0]

	for i := 1; i < len(temp); i++ {
		if temp[i] == prevChar {
			count++
		} else {
			result.WriteString(strconv.Itoa(count))
			result.WriteString(string(prevChar))
			count = 1
			prevChar = temp[i]
		}
	}
	result.WriteString(strconv.Itoa(count))
	result.WriteString(string(prevChar))

	return result.String()
}

// func countAndSay(n int) string {
// 	if n == 1 {
// 		return "1"
// 	}

// 	var result strings.Builder
// 	var temp = []byte(countAndSay(n - 1))

// 	var count = 1
// 	var prevChar = temp[0]

// 	for i := 1; i < len(temp); i++ {
// 		if temp[i] == prevChar {
// 			count++
// 		} else {
// 			result.WriteString(strconv.Itoa(count))
// 			result.WriteString(string(prevChar))
// 			count = 1
// 			prevChar = temp[i]
// 		}
// 	}
// 	result.WriteString(strconv.Itoa(count))
// 	result.WriteString(string(prevChar))

// 	return result.String()
// }

// func countAndSay(n int) string {
// 	if n == 1 {
// 		return "1"
// 	}

// 	temp := countAndSay(n - 1)
// 	length := len(temp)
// 	result := ""

// 	count := 1
// 	prevChar := temp[0] // byte

// 	for i := 1; i < length; i++ {
// 		if temp[i] == prevChar {
// 			count++
// 			continue
// 		}
// 		result = result + strconv.Itoa(count) + string(prevChar)
// 		prevChar = temp[i]
// 		count = 1
// 	}
// 	result = result + strconv.Itoa(count) + string(prevChar)

// 	return result
// }

// 1. 1       -> 11
// 2. 11      -> 21
// 3. 21      -> 1211
// 4. 1211    -> 111221
// 5. 111221  -> 312211
