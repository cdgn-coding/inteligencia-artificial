package puzzle8

type CopySolver struct{}

func (solver CopySolver) limitedDepthFirstSearch(state State, depthLimit int) (State, *SolverError) {
	var stateStack stack = []State{}
	stateStack.Push(state)

	var visitedStates map[string]bool = make(map[string]bool)
	var currentState State
	var currentDepth int
	for len(stateStack) > 0 {
		currentState = stateStack.Pop()
		currentDepth = len(currentState.AppliedActions)

		if visitedStates[currentState.hash()] {
			continue
		}
		visitedStates[currentState.hash()] = true

		if currentState.IsSolved() {
			return currentState, nil
		}

		if currentDepth == depthLimit {
			continue
		}

		for _, action := range currentState.sucessorActions() {
			newState := currentState.copy()
			newState.applyAction(action)
			stateStack.Push(newState)
		}
	}

	return state, &SolverError{"No solution found"}
}

func (solver CopySolver) Solve(state State) (State, *SolverError) {
	for i := 1; ; i += 5 {
		solution, err := solver.limitedDepthFirstSearch(state, i)
		if err == nil {
			return solution, nil
		}
	}
}
