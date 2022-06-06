package funcional

// Filter a slice passed as argument, with a function as condition
func Filter[T any](slice []T, fn func(T) bool) (result []T) {
	for _, value := range slice {
		if fn(value) {
			result = append(result, value)
		}
	}
	return result
}
