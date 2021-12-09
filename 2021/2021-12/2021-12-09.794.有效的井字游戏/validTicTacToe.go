package main

func main() {
	board := []string{
		"O  ", "   ", "   ",
	}
	validTicTacToe(board)
}

func validTicTacToe(board []string) bool {
	// 1. x == o+1  or  x == o
	// 2. x 和 o 不能同时赢,横竖斜
	// 3. o赢的时候 x == o
	// 4. x赢的时候 x == o+1
	var (
		xWin, oWin       bool
		xCounts, oCounts int
	)

	// 统计个数
	for _, b := range board {
		for _, str := range b {
			if str == 'X' {
				xCounts++
			} else if str == 'O' {
				oCounts++
			}
		}
	}

	xWin, oWin = whoWins(board)
	if xWin && oWin {
		return false
	}

	if oWin && oCounts != xCounts {
		return false
	}

	if xWin && xCounts != oCounts+1 {
		return false
	}

	if !(xCounts == oCounts || xCounts == oCounts+1) {
		return false
	}

	return true
}

func whoWins(board []string) (bool, bool) {
	var xWin, oWin bool
	for i, b := range board {

		// 横
		if b == "XXX" {
			xWin = true
		}

		if b == "OOO" {
			oWin = true
		}

		// 竖
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] && board[2][i] == 'X' {
			xWin = true
		}

		if board[0][i] == board[1][i] && board[1][i] == board[2][i] && board[2][i] == 'O' {
			oWin = true
		}

		// 斜
		if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[0][0] == 'X' {
			xWin = true
		}
		if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[0][0] == 'O' {
			oWin = true
		}

		if board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[0][2] == 'X' {
			xWin = true
		}

		if board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[0][2] == 'O' {
			oWin = true
		}
	}
	return xWin, oWin
}
