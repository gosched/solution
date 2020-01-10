package leetcode

func isValidSudoku(board [][]byte) bool {

	for i := 0; i < 9; i++ {
		row := [9]byte{}
		column := [9]byte{}
		box := [9]byte{}

		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				row[board[i][j]-49]++
			}

			if board[j][i] != '.' {
				column[board[j][i]-49]++
			}

			rowIndex := 3*(i/3) + j/3
			columnIndex := 3*(i%3) + j%3

			if board[rowIndex][columnIndex] != '.' {
				box[board[rowIndex][columnIndex]-49]++
			}
		}

		for _, v := range row {
			if v > 1 {
				return false
			}
		}

		for _, v := range column {
			if v > 1 {
				return false
			}
		}

		for _, v := range box {
			if v > 1 {
				return false
			}
		}
	}

	return true
}

// box1 -> box2 -> box3
// box4 -> box5 -> box6
// box7 -> box8 -> box9
