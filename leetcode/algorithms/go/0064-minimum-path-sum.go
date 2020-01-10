package leetcode

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	var min = func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	var n = len(grid)
	var m = len(grid[0])

	var dp = make([][]int, n, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m, m)
	}

	dp[0][0] = grid[0][0]

	for i := 1; i < n; i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}

	for j := 1; j < m; j++ {
		dp[0][j] = grid[0][j] + dp[0][j-1]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[n-1][m-1]
}

// import "math"

// func minPathSum(grid [][]int) int {
// 	if len(grid) == 0 || len(grid[0]) == 0 {
// 		return 0
// 	}
// 	var min = func(x, y int) int {
// 		if x < y {
// 			return x
// 		}
// 		return y
// 	}
// 	var n = len(grid)
// 	var m = len(grid[0])
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < m; j++ {
// 			if i == 0 && j == 0 {
// 				continue
// 			}
// 			var up, left = 0, 0
// 			if i == 0 {
// 				up = math.MaxInt64
// 			} else {
// 				up = grid[i-1][j]
// 			}
// 			if j == 0 {
// 				left = math.MaxInt64
// 			} else {
// 				left = grid[i][j-1]
// 			}
// 			grid[i][j] += min(up, left)
// 		}
// 	}
// 	return grid[n-1][m-1]
// }

// grid
// [1,3,1]
// [1,5,1]
// [4,2,1]

// dp
// [1,4,5]
// [2,7,6]
// [6,8,7]

// dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
// dp[n-1][m-1] == 7
