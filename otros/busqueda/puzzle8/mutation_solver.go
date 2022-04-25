package puzzle8

type MutationSolver struct{}

func (solver MutationSolver) limitedDepthFirstSearch(currentState *State, depthLimit, currentDepth int, visitedStates map[string]bool) *SolverError {
	if visitedStates[currentState.hash()] {
		return &SolverError{"Path is not a solution"}
	}

	visitedStates[currentState.hash()] = true
	if currentState.IsSolved() {
		return nil
	}

	if currentDepth == depthLimit {
		return &SolverError{"Path reached depth limit"}
	}

	for _, action := range currentState.sucessorActions() {
		currentState.applyAction(action)
		err := solver.limitedDepthFirstSearch(currentState, depthLimit, currentDepth+1, visitedStates)

		if err == nil {
			return nil
		}

		currentState.revertLastAction()
	}

	return &SolverError{"Path is not a solution within the depth limit"}
}

func (solver MutationSolver) Solve(state State) (State, *SolverError) {
	for i := 1; ; i += 5 {
		var currentState State = state.copy()
		var visitedStates map[string]bool = make(map[string]bool)
		err := solver.limitedDepthFirstSearch(&currentState, i, 0, visitedStates)
		if err == nil {
			return currentState, nil
		}
	}
}
