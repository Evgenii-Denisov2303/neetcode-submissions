func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	columns := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		columns[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			value := board[row][column]

			if value == '.' {
				continue
			}

			boxIndex := (row/3)*3 + column/3

			if rows[row][value] ||
				columns[column][value] ||
				boxes[boxIndex][value] {
				return false
			}

			rows[row][value] = true
			columns[column][value] = true
			boxes[boxIndex][value] = true
		}
	}

	return true
}