package puzzle8

type Puzzle8Solver interface {
	Solve(state State, depthLimit int) State
}
