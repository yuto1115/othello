package tools

type Player int

func (p *Player) NextPlayer() Player {
	if *p == Black {
		return White
	} else {
		return Black
	}
}

func (p *Player) mark() string {
	if *p == Black {
		return "●"
	} else {
		return "○"
	}
}

func (p Player) string() string {
	if p == Black {
		return "Black"
	} else {
		return "White"
	}
}

func (p Player) ToPiece() Piece {
	if p == Black {
		return Black
	} else {
		return White
	}
}
