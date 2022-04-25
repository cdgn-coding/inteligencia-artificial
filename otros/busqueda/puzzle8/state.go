package puzzle8

import "fmt"

type State struct {
	Board          [3][3]int
	Blank          Position
	AppliedActions []Action
}

// constructor
func NewState(board [3][3]int) State {
	var blank Position
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 0 {
				blank = Position{i, j}
				break
			}
		}
	}

	return State{
		Board:          board,
		Blank:          blank,
		AppliedActions: []Action{},
	}
}

func (state State) peek(position Position) int {
	return state.Board[position.X][position.Y]
}

func (state *State) assign(position Position, value int) {
	state.Board[position.X][position.Y] = value
}

func (state State) copy() State {
	return state
}

func (state State) hash() string {
	var hash string

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			position := Position{i, j}
			hash += fmt.Sprintf("%d", state.peek(position))
		}
	}

	return hash
}

func (state State) IsSolved() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if state.Board[i][j] != i*3+j {
				return false
			}
		}
	}

	return true
}

func (state State) sucessorActions() []Action {
	var actions []Action

	if state.Blank.X < 2 {
		actions = append(actions, DOWN)
	}

	if state.Blank.X > 0 {
		actions = append(actions, UP)
	}

	if state.Blank.Y > 0 {
		actions = append(actions, LEFT)
	}

	if state.Blank.Y < 2 {
		actions = append(actions, RIGHT)
	}

	return actions
}

func (state *State) applyAction(action Action) {
	newBlank := state.Blank.Move(action)
	oldBlank := state.Blank
	movingValue := state.peek(newBlank)
	state.assign(newBlank, 0)
	state.assign(oldBlank, movingValue)
	state.AppliedActions = append(state.AppliedActions, action)
	state.Blank = newBlank
}
