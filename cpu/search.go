package cpu

import (
	"fmt"
	"github.com/yuto1115/othello/tools"
	"os"
)

func lastEval(b *tools.Board, turn int, firstPlayer tools.Player, EVAL *[]int) int {
	//fmt.Printf("turn %d\n",turn)
	if turn >= 4 && b.GetPlayer() == firstPlayer {
		//fmt.Println("toutatu")
		return Eval(b, EVAL)
	}

	choice := b.EnumAllChoices()
	flag := false
	var val int

	for _, pos := range *choice {
		//fmt.Printf("turn %d ",turn)
		//fmt.Printf("%d %dに置いてみたよ\n",pos.I+1,pos.J+1)
		var histories = make([]tools.History, 0, 32)

		err := b.Place(pos, &histories)
		if err != nil {
			fmt.Println("1")
			fmt.Println(err)
			os.Exit(1)
		}
		//b.Display()

		b.ChangePlayer()
		str := b.Judge()
		SKIP := false
		if str != "" {
			if str == "Skip" {
				SKIP = true
			}
			if firstPlayer == tools.Black {
				if str == "Black" {
					return 1e18
				} else {
					return -1e18
				}
			} else {
				if str == "White" {
					return 1e18
				} else {
					return -1e18
				}
			}
		}

		b.ChangePlayer()
		b.Skip()
		if SKIP {
			b.Skip()
		}

		res := lastEval(b, turn+1, firstPlayer, EVAL)
		if flag {
			if SKIP {
				if val < res {
					val = res
				}
			} else {
				if val > res {
					val = res
				}
			}
		} else {
			flag = true
			val = res
		}

		b.Reverse(&histories)
		b.ReverseSkip()
		if SKIP {
			b.ReverseSkip()
		}
	}

	if !flag {
		fmt.Println("NO CHOICE")
		os.Exit(1)
	}

	return val
}

func SearchNextChoice(b *tools.Board, EVAL *[]int) (int, int) {
	choice := b.EnumAllChoices()
	flag := false
	var mx int
	var POS tools.Position
	finished := make(chan bool)
	cnt := 0
	for _, pos := range *choice {
		cnt++
		//fmt.Printf("%d %dに置いてみるよ\n",pos.I+1,pos.J+1)
		var histories = make([]tools.History, 0, 32)
		err := b.Place(pos, &histories)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//b.Display()

		go func(nb tools.Board) {
			nb.Skip()
			pl := nb.GetPlayer()
			res := lastEval(&nb, 1, pl.NextPlayer(), EVAL)
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
		b.Reverse(&histories)
	}

	for i := 0; i < cnt; i++ {
		<-finished
	}

	if !flag {
		fmt.Println("NO CHOICE")
		os.Exit(1)
	}

	return POS.I, POS.J
}
