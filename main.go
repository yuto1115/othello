package main

import (
	"fmt"

	"othello/makeCpu"
	"othello/simulate"
)

func main() {
	fmt.Println("choose a mode\ntwo people -> 'vs' one person -> 'cpu' make cpu -> 'make'")
	var s string
	for {
		_, err := fmt.Scanf("%s", &s)
		if err != nil {
			fmt.Println("invalid input; please try again")
			continue
		}

		if s == "vs" {
			simulate.Simulate(false)
			break
		} else if s == "cpu" {
			simulate.Simulate(true)
			break
		} else if s == "make" {
			makeCpu.MakeCpu()
			break
		} else {
			fmt.Println("invalid input; please try again")
			continue
		}
	}
}
