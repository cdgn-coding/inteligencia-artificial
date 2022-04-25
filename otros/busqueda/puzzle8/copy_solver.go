package puzzle8

type CopySolver struct{}

func (solver CopySolver) Solve(state State, depthLimit int) State {
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
			return currentState
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

	panic("No solution found")
}
