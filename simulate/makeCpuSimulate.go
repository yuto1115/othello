package simulate

import (
	"log"

	"othello/cpu"
	"othello/tools"
)

// 0 -> Black Wins
// 1 -> Draw
// 2 -> White Wins
func MakeCpuSimulate(BlackEVAL []int, WhiteEVAL []int) int {
	if len(BlackEVAL) != cpu.Pow3[16] {
		log.Fatal("Size of BlackEval is wrong")
	}
	if len(WhiteEVAL) != cpu.Pow3[16] {
		log.Fatal("Size of WhiteEval is wrong")
	}
	board := tools.NewBoard()

	for {
		var i, j int

		if board.Player == tools.Black {
			i, j = cpu.SearchNextChoice(board, BlackEVAL)
		} else {
			i, j = cpu.SearchNextChoice(board, WhiteEVAL)
		}

		pos := tools.Position{I: i, J: j}
		var histories = make([]tools.History, 0, 32)
		err := board.Place(pos, &histories)
		if err != nil {
			log.Fatal(err)
		}

		board.Skip()
		s := board.Judge()
		if s == "Black" {
			return 0
		} else if s == "White" {
			return 2
		} else if s == "Draw" {
			return 1
		} else if s == "Skip" {
			board.Skip()
		}
	}
}
