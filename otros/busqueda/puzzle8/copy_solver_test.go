package puzzle8

import (
	"fmt"
	"math/rand"
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
	var solution State

	solution, _ = solver.Solve(state)

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
	var solution State

	solution, _ = solver.Solve(state)

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
	var solution State

	solution, _ = solver.Solve(state)

	assert.True(test, solution.IsSolved())
}

func TestThreeMovementsCase(test *testing.T) {
	var state State = NewState([3][3]int{
		{1, 4, 2},
		{3, 7, 5},
		{6, 0, 8},
	})

	var solver Puzzle8Solver = CopySolver{}
	var solution State

	solution, _ = solver.Solve(state)

	assert.True(test, solution.IsSolved())
}

func TestPresetStates(test *testing.T) {
	var solver Puzzle8Solver = CopySolver{}
	var solution State
	for _, fixture := range fixtures {
		solution, _ = solver.Solve(fixture.state)
		assert.True(test, solution.IsSolved(), fmt.Sprintf("Expected %s to found a solution", fixture.state.hash()))
	}
}

func BenchmarkCopySolver(b *testing.B) {
	rand.Seed(42)
	var solver Puzzle8Solver = CopySolver{}
	for _, fixture := range fixtures {
		b.Run(fmt.Sprintf("Solving Puzzle with %d movements", fixture.movements), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				solver.Solve(fixture.state)
			}
		})
	}
}
