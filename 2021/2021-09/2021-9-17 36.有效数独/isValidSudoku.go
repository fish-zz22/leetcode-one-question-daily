package main

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	isValidSudoku(board)
}

func isValidSudoku(board [][]byte) bool {
	var row, col [9][9]int
	var box [3][3][9]int
	for i, cols := range board {
		for j, v := range cols {
			if v == '.' {
				continue
			}
			ch := v - '1'
			row[i][ch]++
			col[j][ch]++
			box[i/3][j/3][ch]++
			if row[i][ch] > 1 || col[j][ch] > 1 || box[i/3][j/3][ch] > 1 {
				return false
			}
		}
	}
	return true
}
