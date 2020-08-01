package main

import (
	"fmt"
	"github.com/yuto1115/othello/tools"
	"os"
	"strconv"
)

func main() {
	board := tools.NewBoard()
	for {
		board.Display()
		fmt.Println("Enter two numbers '[row number][column number]'")

	LABEL:
		var s string
		fmt.Scanf("%s", &s)

		if s == "exit" {
			os.Exit(0)
		}
		if s == "enum" {
			choice := board.EnumAllChoices()
			for _, pos := range *choice {
				fmt.Printf("%d %d\n", pos.I+1, pos.J+1)
			}
			goto LABEL
		}

		if len(s) != 2 {
			fmt.Println("invalid input; please try again")
			goto LABEL
		}

		i, e1 := strconv.Atoi(s[0:1])
		if e1 != nil {
			fmt.Println("invalid input; please try again")
			goto LABEL
		}

		j, e2 := strconv.Atoi(s[1:2])
		if e2 != nil {
			fmt.Println("invalid input; please try again")
			goto LABEL
		}

		pos := tools.Position{I: i - 1, J: j - 1}
		if pos.OutOfRange() {
			fmt.Println("invalid input; please try again")
			goto LABEL
		}

		err := board.Place(pos)
		if err != nil {
			fmt.Println("You can't place your piece there; please try again")
			goto LABEL
		}
	}
}
