package tools

import (
	"errors"
	"fmt"
	"os"
)

type Position struct {
	I, J int
}

func (pos Position) OutOfRange() bool {
	i, j := pos.I, pos.J
	return i < 0 || i >= 8 || j < 0 || j >= 8
}

type Board struct {
	board  [8][8]Piece
	player Player
	turn   int
}

func NewBoard() *Board {
	var b Board

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			b.board[i][j] = Empty
		}
	}

	b.board[3][3] = White
	b.board[3][4] = Black
	b.board[4][3] = Black
	b.board[4][4] = White
	b.player = Black
	b.turn = 1

	return &b
}

func (b *Board) Display() {
	fmt.Println("| |1|2|3|4|5|6|7|8|")

	for i := 0; i < 8; i++ {
		fmt.Printf("|%d|", i+1)
		for j := 0; j < 8; j++ {
			b.board[i][j].display()
			fmt.Printf("|")
		}
		fmt.Printf("\n")
	}

	fmt.Printf("next player is %s\n", b.player.mark())
	bl, wh := b.count()
	fmt.Printf("black: %d, white: %d\n", bl, wh)
}

func (b *Board) displayOnlyBoard() {
	fmt.Println("| |1|2|3|4|5|6|7|8|")

	for i := 0; i < 8; i++ {
		fmt.Printf("|%d|", i+1)
		for j := 0; j < 8; j++ {
			b.board[i][j].display()
			fmt.Printf("|")
		}
		fmt.Printf("\n")
	}
}

func (b *Board) count() (int, int) {
	var bl, wh int

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if b.board[i][j] == Black {
				bl++
			} else if b.board[i][j] == White {
				wh++
			}
		}
	}

	return bl, wh
}

func (b *Board) Place(pos Position) error {
	if !b.isPlaceable(pos) {
		return errors.New("invalid input")
	}

	get := b.goAround(pos)

	b.board[pos.I][pos.J] = b.player.toPiece()
	for _, newPos := range *get {
		b.board[newPos.I][newPos.J] = b.player.toPiece()
	}
	b.player = b.player.nextPlayer()
	b.turn++
	b.judge()

	return nil
}

func (b *Board) isPlaceable(pos Position) bool {
	i, j := pos.I, pos.J
	if pos.OutOfRange() {
		return false
	}
	if b.board[i][j] != Empty {
		return false
	}

	get := b.goAround(pos)
	if len(*get) == 0 {
		return false
	}

	return true
}

func (b *Board) EnumAllChoices() *[]Position {
	res := make([]Position, 0, 64)

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			pos := Position{i, j}
			if b.isPlaceable(pos) {
				res = append(res, pos)
			}
		}
	}

	return &res
}

func (b *Board) skip() {
	b.player = b.player.nextPlayer()
}

func (b *Board) judge() {
	choice := b.EnumAllChoices()

	if len(*choice) == 0 {
		b.skip()
		choice = b.EnumAllChoices()

		if len(*choice) == 0 {
			bl, wh := b.count()

			b.displayOnlyBoard()
			if bl > wh {
				fmt.Printf("Black has won by %d - %d", bl, wh)
			} else if bl < wh {
				fmt.Printf("White has won by %d - %d", bl, wh)
			} else {
				fmt.Printf("Draw  %d - %d", bl, wh)
			}

			os.Exit(0)
		} else {
			fmt.Printf("%s passed; It's %s turn again\n", b.player.nextPlayer().string(), b.player.string())
		}
	}
}

func (b *Board) goAround(pos Position) *[]Position {
	res := make([]Position, 0, 32)

	res = append(res, *b.search(pos, goUp)...)
	res = append(res, *b.search(pos, goUpRight)...)
	res = append(res, *b.search(pos, goRight)...)
	res = append(res, *b.search(pos, goDownRight)...)
	res = append(res, *b.search(pos, goDown)...)
	res = append(res, *b.search(pos, goDownLeft)...)
	res = append(res, *b.search(pos, goLeft)...)
	res = append(res, *b.search(pos, goUpLeft)...)

	return &res
}

func goUp(pos Position) Position        { return Position{pos.I - 1, pos.J} }
func goUpRight(pos Position) Position   { return Position{pos.I - 1, pos.J + 1} }
func goRight(pos Position) Position     { return Position{pos.I, pos.J + 1} }
func goDownRight(pos Position) Position { return Position{pos.I + 1, pos.J + 1} }
func goDown(pos Position) Position      { return Position{pos.I + 1, pos.J} }
func goDownLeft(pos Position) Position  { return Position{pos.I + 1, pos.J - 1} }
func goLeft(pos Position) Position      { return Position{pos.I, pos.J - 1} }
func goUpLeft(pos Position) Position    { return Position{pos.I - 1, pos.J - 1} }

func (b *Board) search(pos Position, f func(Position) Position) *[]Position {
	res := make([]Position, 0, 8)

	for nowPos := f(pos); !nowPos.OutOfRange(); nowPos = f(nowPos) {
		switch b.board[nowPos.I][nowPos.J] {
		case Empty:
			return &[]Position{}
		case b.player.toPiece():
			return &res
		case b.player.nextPlayer().toPiece():
			res = append(res, nowPos)
		}
	}

	return &[]Position{}
}
