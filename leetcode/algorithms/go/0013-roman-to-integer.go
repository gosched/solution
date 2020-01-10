package leetcode

func romanToInt(s string) int {
    sum := 0
    
    m := map[byte]int {
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
    }
    
    for i := 0; i < len(s); i++ {
        sum += m[s[i]]
        if i > 0 && m[s[i-1]] < m[s[i]] {
            sum = sum - 2 * m[s[i-1]]
        }
    }
    
    return sum
}