package reduce

type Reducer[T any] func(accumulator, current T) T

// Reduce is a generic implmentation of a reduce function, where the passed
// function `reducer` is called using the starting state `initial` and the first
// value of `elements`, and then subsequently with each result and subsequent element
// of the passed sliced `elements`.
func Reduce[T any](elements []T, reducer Reducer[T], initial T) T {
	if len(elements) == 0 {
		return initial
	}
	state := reducer(initial, elements[0])
	for i := 1; i < len(elements); i++ {
		state = reducer(state, elements[i])
	}
	return state
}
