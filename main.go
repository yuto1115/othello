package main

import (
	"fmt"
	"github.com/yuto1115/othello/tools"
	"os"
	"strconv"
)

func main() {
	board := tools.NewBoard()
	flag := true
	for {
		if flag {
			board.Display()
			fmt.Println("Enter two numbers '[row number][column number]'")
		}

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
			flag = false
			continue
		}

		if len(s) != 2 {
			fmt.Println("invalid input; please try again")
			flag = false
			continue
		}

		i, e1 := strconv.Atoi(s[0:1])
		if e1 != nil {
			fmt.Println("invalid input; please try again")
			flag = false
			continue
		}

		j, e2 := strconv.Atoi(s[1:2])
		if e2 != nil {
			fmt.Println("invalid input; please try again")
			flag = false
			continue
		}

		pos := tools.Position{I: i - 1, J: j - 1}
		err := board.Place(pos)
		if err != nil {
			fmt.Println(err)
			flag = false
			continue
		}

		flag = true
	}
}
