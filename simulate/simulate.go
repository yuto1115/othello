package simulate

import (
	"fmt"
	"os"
	"strconv"

	"othello/cpu"
	"othello/tools"
)

func Simulate(useCpu bool) {
	board := tools.NewBoard()
	flag := true
	var EVAL = make([]int, cpu.Pow3[16], cpu.Pow3[16])
	for {
		if flag {
			board.Display()
			fmt.Println("Enter two numbers '[row number][column number]'")
		}

		var i, j int

		if !useCpu || board.Player == tools.Black {
			var s string
			_, err := fmt.Scanf("%s", &s)
			if err != nil {
				fmt.Println("invalid input; please try again")
				flag = false
				continue
			}

			if s == "exit" {
				os.Exit(0)
			}
			if s == "enum" {
				choice := board.EnumAllChoices()
				for _, pos := range choice {
					fmt.Printf("%d %d\n", pos.I+1, pos.J+1)
				}
				flag = false
				continue
			}

			if len(s) != 2 {
				fmt.Println("invalid input; please try again")
				flag = false
				continue
			}

			ni, err2 := strconv.Atoi(s[0:1])
			if err2 != nil {
				fmt.Println("invalid input; please try again")
				flag = false
				continue
			}

			nj, err3 := strconv.Atoi(s[1:2])
			if err3 != nil {
				fmt.Println("invalid input; please try again")
				flag = false
				continue
			}

			i = ni
			j = nj
		} else {
			i, j = cpu.SearchNextChoice(board, EVAL)
			i += 1
			j += 1
			fmt.Printf("CPU chose %d%d\n", i, j)
		}

		pos := tools.Position{I: i - 1, J: j - 1}
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
