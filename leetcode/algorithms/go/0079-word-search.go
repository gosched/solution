package leetcode

func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}

	n := len(board[0])
	if n == 0 {
		return false
	}

	if len(word) == 0 {
		return false
	}

	target := len(word) - 1

	var helper func(index, i, j int) bool
	helper = func(index, i, j int) bool {
		if i < 0 || j < 0 || i >= m || j >= n || board[i][j] != word[index] {
			return false
		}
		if index == target {
			return true
		}
		board[i][j] = '.'
		if helper(index+1, i+1, j) ||
			helper(index+1, i-1, j) ||
			helper(index+1, i, j+1) ||
			helper(index+1, i, j-1) {
			return true
		}
		board[i][j] = word[index]
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if helper(0, i, j) {
				return true
			}
		}
	}

	return false
}
