package funcional

// Find a value of a slice passed as an argument and a function as condition, ok is true is a value was found
func Find[T any](slice []T, fn func(T) bool) (find T, ok bool) {
	for _, value := range slice {
		if fn(value) {
			ok = true
			find = value
			break
		}
	}
	return find, ok
}
