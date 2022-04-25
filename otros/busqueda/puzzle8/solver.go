package puzzle8

type SolverError struct {
	message string
}

func (e *SolverError) Error() string {
	return e.message
}

type Puzzle8Solver interface {
	Solve(state State) (State, *SolverError)
}
