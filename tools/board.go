package tools

import "fmt"

type Board struct {
	board  [8][8]Piece
	player Player
	turn   int
}

func NewBoard() *Board {
	var b Board
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			b.board[i][j] = 2
		}
	}
	b.board[3][3] = 1
	b.board[3][4] = 0
	b.board[4][3] = 0
	b.board[4][4] = 1
	b.player = Black
	b.turn = 1
	return &b
}

func (b Board) count() (int, int) {
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

func (b Board) Display() {
	fmt.Println("| |0|1|2|3|4|5|6|7|")
	for i := 0; i < 8; i++ {
		fmt.Printf("|%d|", i)
		for j := 0; j < 8; j++ {
			b.board[i][j].display()
			fmt.Printf("|")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("next player is %s\n", b.player.string())
	bl, wh := b.count()
	fmt.Printf("black: %d, white: %d\n", bl, wh)
}

func (b Board) skip() {
	b.player = b.player.nextPlayer()
}
