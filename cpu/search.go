package cpu

import (
	"fmt"
	"log"
	"os"

	"othello/tools"
)

func lastEval(b *tools.Board, turn int, firstPlayer tools.Player, EVAL []int, nowMax int, existMax bool) int {
	if turn == 6 {
		if b.Player != firstPlayer {
			log.Fatal("turn calc is wrong!!!!!")
		}
		return Eval(b, EVAL)
	}

	choice := b.EnumAllChoices()
	flag := false
	var mx int

	if len(choice) == 0 {
		b.ChangePlayer()
		res := -lastEval(b, turn+1, firstPlayer, EVAL, 0, false)
		b.ChangePlayer()
		return res
	}

	for _, pos := range choice {
		var histories = make([]tools.History, 0, 32)

		err := b.Place(pos, &histories)
		if err != nil {
			fmt.Println("1")
			log.Fatal(err)
		}

		b.ChangePlayer()

		var res int
		if flag {
			res = -lastEval(b, turn+1, firstPlayer, EVAL, mx, true)
		} else {
			res = -lastEval(b, turn+1, firstPlayer, EVAL, 0, false)
		}

		b.ChangePlayer()
		b.Reverse(histories)

		if existMax && res > -nowMax {
			return -nowMax
		}

		if flag {
			if res > mx {
				mx = res
			}
		} else {
			flag = true
			mx = res
		}
	}

	if !flag {
		fmt.Println("NO CHOICE")
		os.Exit(1)
	}

	return mx
}

func SearchNextChoice(b *tools.Board, EVAL []int) (int, int) {
	choice := b.EnumAllChoices()
	flag := false
	var mx int
	var POS tools.Position
	finished := make(chan bool, 20)

	for _, pos := range choice {
		var histories = make([]tools.History, 0, 32)
		err := b.Place(pos, &histories)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		go func(nb tools.Board) {
			nb.ChangePlayer()
			res := -lastEval(&nb, 1, nb.Player.NextPlayer(), EVAL, 0, false)
			if flag {
				if res > mx {
					mx = res
					POS = pos
				}
			} else {
				flag = true
				mx = res
				POS = pos
			}
			finished <- true
		}(*b)
		b.Reverse(histories)
	}

	for i := 0; i < len(choice); i++ {
		<-finished
	}

	if !flag {
		fmt.Println("NO CHOICE")
		os.Exit(1)
	}

	return POS.I, POS.J
}
