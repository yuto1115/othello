package main

import (
	"fmt"
	"github.com/yuto1115/othello/simulate"
)

func main() {
	fmt.Println("choose a mode\ntwo person -> 'vs' one people -> 'cpu'")
	var s string
	for {
		_, e := fmt.Scanf("%s", &s)
		if e != nil {
			fmt.Println("invalid input; please try again")
			continue
		}

		if s == "vs" {
			simulate.Simulate(false)
			break
		} else if s == "cpu" {
			simulate.Simulate(true)
			break
		} else {
			fmt.Println("invalid input; please try again")
			continue
		}
	}
}
