package cpu

import (
	"othello/tools"
)

var Pow3 = [17]int{1, 3, 9, 27, 81, 243, 729, 2187, 6561, 19683, 59049, 177147, 531441, 1594323, 4782969, 14348907, 43046721}

func Eval(b *tools.Board, EVAL []int) int {
	str := b.Judge()
	nextPlayer := b.Player.NextPlayer()
	if str == "Draw" {
		return 0
	} else if str == b.Player.ToString() {
		return 1e18
	} else if str == nextPlayer.ToString() {
		return -1e18
	}

	res1 := 0
	res2 := 0
	res3 := 0
	res4 := 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			var val int
			board := &b.Board
			if board[i][j] == tools.Empty {
				val = 0
			} else if board[i][j] == b.Player.ToPiece() {
				val = 1
			} else {
				val = 2
			}
			if i < 4 && j < 4 {
				res1 += Pow3[i*4+j] * val
			} else if i < 4 {
				res2 += Pow3[i*4+(7-j)] * val
			} else if j < 4 {
				res3 += Pow3[(7-i)*4+j] * val
			} else {
				res4 += Pow3[(7-i)*4+(7-j)] * val
			}
		}
	}
	res1 = EVAL[res1]
	res2 = EVAL[res2]
	res3 = EVAL[res3]
	res4 = EVAL[res4]
	return res1 + res2 + res3 + res4
}
