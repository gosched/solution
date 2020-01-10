package leetcode

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
		l := 0
		r := n - 1
		for l < r {
			matrix[i][l], matrix[i][r] = matrix[i][r], matrix[i][l]
			l++
			r--
		}
	}
}

// 1  2  3        1  4  7        7  4  1
// 4  5  6        2  5  8        8  5  2
// 7  8  9        3  6  9        9  6  3

// matrix[0][1] <-> matrix[1][0] (2<->4)
// matrix[0][2] <-> matrix[2][0] (3<->7)
// matrix[1][2] <-> matrix[2][1] (6<->8)

// matrix[0][0] <-> matrix[0][2] (1<->7)
// matrix[1][0] <-> matrix[1][2] (2<->8)
// matrix[2][0] <-> matrix[2][2] (3<->9)
