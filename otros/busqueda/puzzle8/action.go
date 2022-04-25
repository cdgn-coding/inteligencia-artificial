package puzzle8

type Action string

const (
	UP    Action = "UP"
	DOWN  Action = "DOWN"
	LEFT  Action = "LEFT"
	RIGHT Action = "RIGHT"
)

func (action Action) Inverse() Action {
	if action == UP {
		return DOWN
	}

	if action == DOWN {
		return UP
	}

	if action == LEFT {
		return RIGHT
	}

	if action == RIGHT {
		return LEFT
	}

	panic("Invalid action")
}
