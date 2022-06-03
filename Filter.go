package funcional

func Filter[T any](slice []T, fn func(T) bool) (result []T) {
	for _, value := range slice {
		if fn(value) {
			result = append(result, value)
		}
	}
	return result
}
