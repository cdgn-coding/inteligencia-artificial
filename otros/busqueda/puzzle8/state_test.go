package puzzle8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSolvedOK(t *testing.T) {
	var state State = NewState([3][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	})

	if state.IsSolved() != true {
		t.Errorf("Expected state to be solved")
	}
}

func TestActionLeftOK(test *testing.T) {
	var state State = NewState([3][3]int{
		{1, 0, 2},
		{3, 4, 5},
		{6, 7, 8},
	})

	state.applyAction(LEFT)

	assert.True(test, state.IsSolved())
}

func TestActionDownOK(test *testing.T) {
	var state State = NewState([3][3]int{
		{1, 0, 2},
		{3, 4, 5},
		{6, 7, 8},
	})

	state.applyAction(DOWN)

	assert.False(test, state.IsSolved())
	assert.Equal(test, state.Board, [3][3]int{
		{1, 4, 2},
		{3, 0, 5},
		{6, 7, 8},
	})
}
