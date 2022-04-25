package puzzle8

import "math/rand"

type fixture struct {
	state     State
	movements int
}

func generateStates(size int, maxMovements int) []fixture {
	rand.Seed(42)
	var table []fixture = make([]fixture, size)
	var movements int
	var state State

	for i := 0; i < size; i++ {
		movements = rand.Intn(maxMovements)
		state = NewRandomState(movements)
		table[i] = fixture{
			state:     state,
			movements: movements,
		}
	}

	return table
}

var fixtures []fixture = generateStates(10, 100)
