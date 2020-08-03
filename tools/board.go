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

type History struct {
	pos   Position
	piece Piece
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

func (b *Board) GetPlayer() Player {
	return b.player
}

func (b *Board) GetBoard() [8][8]Piece {
	return b.board
}

func (b *Board) ChangePlayer() {
	b.player = b.player.NextPlayer()
}

func (b *Board) Reverse(histories *[]History) {
	for _, hist := range *histories {
		//fmt.Printf("%d %dを元に戻したよ\n",hist.pos.I+1,hist.pos.J+1)
		b.board[hist.pos.I][hist.pos.J] = hist.piece
		//b.Display()
	}
}

func (b *Board) Display() {
	b.displayOnlyBoard()

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

func (b *Board) Place(pos Position, histories *[]History) error {
	if !b.isPlaceable(pos) {
		return errors.New("You can't place your piece there; please try again")
	}

	get := b.goAround(pos)

	*histories = append(*histories, History{pos: pos, piece: b.board[pos.I][pos.J]})
	b.board[pos.I][pos.J] = b.player.ToPiece()
	//fmt.Printf("%d %dに新しく置くよ\n",pos.I+1,pos.J+1)
	//b.Display()
	for _, newPos := range *get {
		//fmt.Printf("%d %dをひっくり返したよ\n",newPos.I+1,newPos.J+1)
		*histories = append(*histories, History{pos: newPos, piece: b.board[newPos.I][newPos.J]})
		b.board[newPos.I][newPos.J] = b.player.ToPiece()
		//b.Display()
	}

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

func (b *Board) Skip() {
	b.player = b.player.NextPlayer()
	b.turn++
}

func (b *Board) ReverseSkip() {
	b.player = b.player.NextPlayer()
	b.turn--
}

func (b Board) Judge() string {
	choice := b.EnumAllChoices()

	if len(*choice) == 0 {
		b.Skip()
		choice = b.EnumAllChoices()

		if len(*choice) == 0 {
			bl, wh := b.count()

			if bl > wh {
				return "Black"
			} else if bl < wh {
				return "White"
			} else {
				return "Draw"
			}
		} else {
			return "Skip"
		}
	}
	return ""
}

func (b *Board) Proceed() {
	b.player = b.player.NextPlayer()
	b.turn++

	s := b.Judge()
	if s == "" {
		return
	}

	bl, wh := b.count()
	b.displayOnlyBoard()

	if s == "Black" {
		fmt.Printf("Black has won by %d - %d", bl, wh)
		os.Exit(0)
	} else if s == "White" {
		fmt.Printf("White has won by %d - %d", bl, wh)
		os.Exit(0)
	} else if s == "Draw" {
		fmt.Printf("Draw  %d - %d", bl, wh)
		os.Exit(0)
	} else {
		fmt.Printf("%s passed; It's %s turn again\n", b.player.NextPlayer().string(), b.player.string())
		b.Skip()
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
		case b.player.ToPiece():
			return &res
		case b.player.NextPlayer().ToPiece():
			res = append(res, nowPos)
		}
	}

	return &[]Position{}
}
