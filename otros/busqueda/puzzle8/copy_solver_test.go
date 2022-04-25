package puzzle8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrivialCase(test *testing.T) {
	var state State = NewState([3][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	})

	var solver Puzzle8Solver = CopySolver{}
	var solution State = solver.Solve(state, 1)

	if !solution.IsSolved() {
		test.Error("State was expected to be solved")
	}
}

func TestOneMovementCase(test *testing.T) {
	var state State = NewState([3][3]int{
		{1, 0, 2},
		{3, 4, 5},
		{6, 7, 8},
	})

	var solver Puzzle8Solver = CopySolver{}
	var solution State = solver.Solve(state, 5)

	assert.True(test, solution.IsSolved())
	assert.Equal(test, solution.AppliedActions, []Action{LEFT})
}

func TestTwoMovementsCase(test *testing.T) {
	var state State = NewState([3][3]int{
		{1, 4, 2},
		{3, 0, 5},
		{6, 7, 8},
	})

	var solver Puzzle8Solver = CopySolver{}
	var solution State = solver.Solve(state, 5)

	assert.True(test, solution.IsSolved())
	assert.Equal(test, solution.AppliedActions, []Action{UP, LEFT})
}

func TestThreeMovementsCase(test *testing.T) {
	var state State = NewState([3][3]int{
		{1, 4, 2},
		{3, 7, 5},
		{6, 0, 8},
	})

	var solver Puzzle8Solver = CopySolver{}
	var solution State = solver.Solve(state, 5)

	assert.True(test, solution.IsSolved())
	assert.Equal(test, solution.AppliedActions, []Action{UP, UP, LEFT})
}
