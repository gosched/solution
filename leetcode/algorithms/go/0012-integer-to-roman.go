package leetcode

import "strings"

type symbol struct {
	char  string
	value int
}

func intToRoman(num int) string {
	var symbols = []symbol{
		{char: "I", value: 1},
		{char: "II", value: 2},
		{char: "III", value: 3},
		{char: "IV", value: 4},
		{char: "V", value: 5},
		{char: "IX", value: 9},
		{char: "X", value: 10},
		{char: "XL", value: 40},
		{char: "L", value: 50},
		{char: "XC", value: 90},
		{char: "C", value: 100},
		{char: "CD", value: 400},
		{char: "D", value: 500},
		{char: "CM", value: 900},
		{char: "M", value: 1000},
	}
	var index = 14
	var builder strings.Builder
	for num > 0 {
		var difference = num - symbols[index].value
		if difference < 0 {
			index--
		} else {
			builder.WriteString(symbols[index].char)
			num = difference
		}
	}
	return builder.String()
}

// func intToRoman(num int) string {
//     var v1 = []string{"", "M", "MM", "MMM"}
//     var v2 = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
//     var v3 = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
//     var v4 = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
//     return v1[num / 1000] + v2[(num % 1000) / 100] + v3[(num % 100) / 10] + v4[num % 10]
// }
