package leetcode

func minDistance(word1 string, word2 string) int {
	var min = func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	var m = len(word1)
	var n = len(word2)

	var dp = make([][]int, m+1, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}

	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}

	return dp[m][n]
}

// 從 horse 換成 ros, 令 $ 表示空字串
//   $ r o s
// $ 0 1 2 3
// h 1 1 2 3
// o 2 2 1 2
// r 3 2 2 2
// s 4 3 3 2
// e 5 4 4 3

// 當前字符相同的情況
// 如 ho -> ro, 不考慮相等的 o, 操作次數取決於 h -> r
// if word1[i-1] == word2[j-1], then dp[i][j] == dp[i-1][j-1]

// 當前字符不同的情況
// (hor -> ro) == (ho -> r + 1) or (ho -> ro + 1) or (hor -> r + 1)
// 2           == (2 + 1)       or (1 + 1)        or (2 + 1)
// min(dp[i-1][j-1] + 1, dp[i-1][j] + 1, dp[i][j-1] + 1)
// min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1

// dp[i][j]  代表 從 word1 的前 i 個字 轉成 word2 的前 j 個字 最少操作次數
// dp[m]][n] 代表 從 word1 的前 m 個字 轉成 word2 的前 n 個字 最少操作次數
