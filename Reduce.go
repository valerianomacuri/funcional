package funcional

func Reduce[T1, T2 any](slice []T1, initializer T2, fn func(T2, T1) T2) T2 {
	result := initializer
	for _, value := range slice {
		result = fn(result, value)
	}
	return result
}
