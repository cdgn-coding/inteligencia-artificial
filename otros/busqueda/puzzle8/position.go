package puzzle8

type Position struct {
	X int
	Y int
}

func (p Position) Move(action Action) Position {
	if action == UP {
		return Position{p.X - 1, p.Y}
	}

	if action == DOWN {
		return Position{p.X + 1, p.Y}
	}

	if action == LEFT {
		return Position{p.X, p.Y - 1}
	}

	if action == RIGHT {
		return Position{p.X, p.Y + 1}
	}

	panic("Invalid action")
}
