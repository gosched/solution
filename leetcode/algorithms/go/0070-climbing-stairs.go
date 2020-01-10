package leetcode

func climbStairs(n int) int {
	var x, y = 1, 1
	for n > 1 {
		x, y = y, x+y
		n--
	}
	return y
}

// func climbStairs(n int) int {
//     var dp = make([]int, n+1, n+1)
//     dp[0] = 1
//     dp[1] = 1
//     for i := 2; i <= n; i++ {
//         dp[i] = dp[i-1] + dp[i-2]
//     }
//     return dp[n]
// }

// n == 0 -> 1
// n == 1 -> 1
// n == 2 -> 1 + 1 -> 2
// n == 3 -> 1 + 2 -> 3
