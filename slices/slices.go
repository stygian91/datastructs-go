package slices

// Finds the index of the last occurrence of the search value in the slice.
// Comparison is done with the `==` operator.
// Returns -1 if the element is not found.
func LastIndex[S ~[]E, E comparable](s S, el E) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == el {
			return i
		}
	}

	return -1
}

// Finds the index of the last element, that satisfies fn(s[i]).
// Returns -1 if no elements satisfy the callback.
func LastIndexFunc[S ~[]E, E any](s S, fn func(E) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if fn(s[i]) {
			return i
		}
	}

	return -1
}
