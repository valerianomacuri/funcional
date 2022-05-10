package funcional

func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	result := initializer
	for _, value := range s {
		result = f(result, value)
	}
	return result
}
