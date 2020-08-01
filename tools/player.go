package tools

type Player int

func (p Player) nextPlayer() Player {
	if p == 0 {
		return 1
	} else {
		return 0
	}
}

func (p Player) string() string {
	if p == Black {
		return "●"
	} else {
		return "○"
	}
}
