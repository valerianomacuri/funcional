package funcional

func Filter[T any](s []T, f func(T) bool) []T {
	var result []T
	for _, value := range s {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}
