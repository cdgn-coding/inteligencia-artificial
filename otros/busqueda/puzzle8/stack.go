package puzzle8

type stack []State

func (s *stack) Push(v State) {
	*s = append(*s, v)
}

func (s *stack) Pop() State {
	ret := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return ret
}
