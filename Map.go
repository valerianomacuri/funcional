package funcional

// Map iterates over a slice
func Map[T1, T2 any](slice []T1, fn func(T1) T2) []T2 {
	result := make([]T2, len(slice))
	for index, value := range slice {
		result[index] = fn(value)

	}
	return result
}
