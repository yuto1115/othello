package simulate

import (
	"fmt"
	"log"

	"othello/cpu"
	"othello/tools"
)

// 0 -> Black Wins
// 1 -> Draw
// 2 -> White Wins
func makeCpuSimulate(BlackEVAL []int, WhiteEVAL []int) int {
	if len(BlackEVAL) != cpu.Pow3[16] {
		log.Fatal("Size of BlackEval is wrong")
	}
	if len(WhiteEVAL) != cpu.Pow3[16] {
		log.Fatal("Size of WhiteEval is wrong")
	}
	board := tools.NewBoard()
	flag := true

	for {
		if flag {
			board.Display()
			fmt.Println("Enter two numbers '[row number][column number]'")
		}

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
			fmt.Println(err)
			flag = false
			continue
		}
		board.Proceed()

		flag = true
	}
}
