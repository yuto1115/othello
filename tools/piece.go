package tools

import "fmt"

type Piece int

func (p Piece) display() {
	var s string
	switch p {
	case Black:
		s = "●"
	case White:
		s = "○"
	case Empty:
		s = " "
	}
	fmt.Printf("%s", s)
}
