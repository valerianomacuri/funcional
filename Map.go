package funcional

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	slice := make([]T2, len(s))

	for index, value := range s {
		slice[index] = f(value)

	}
	return slice
}
