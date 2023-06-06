package slices

func Filter[T any](input []T, predicate func(T) bool) []T {
	var output []T
	for _, v := range input {
		if predicate(v) {
			output = append(output, v)
		}
	}
	return output
}

func Map[S any, T any](input []S, mapper func(S) T) []T {
	output := make([]T, len(input))
	for i, v := range input {
		output[i] = mapper(v)
	}
	return output
}

func Reduce[S any, T any](input []S, reducer func(T, S) T, accumulator T) T {
	for _, v := range input {
		accumulator = reducer(accumulator, v)
	}
	return accumulator
}
