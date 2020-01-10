package leetcode

func reverse(x int) int {
	var answer, limit = 0, 1 << 31

	for x != 0 {
		answer = answer*10 + x%10
		x /= 10
	}

	if answer < -limit || answer > limit-1 {
		return 0
	}

	return answer
}

/*
func reverse(x int) int {
	var answer = 0

	for x != 0 {
		answer = answer*10 + x%10
		x /= 10
	}

	if answer < -2147483648 || answer > 2147483647 {
		return 0
	}

	return answer
}
*/

// -5 % 10 == -5

// 2 ^ 31 - 1 == 2147483647 == 0x7FFFFFFF
// 2 ^ 31 == 2147483648 == 0x80000000
// 2 ^ 32 == 4294967296
// 2 ^ 64 == 18446744073709551616

// math.MaxInt32
// math.MinInt32
