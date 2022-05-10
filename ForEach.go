package funcional

func ForEach[T1 any](s []T1, f func(T1)) {
	slice := make([]T1, len(s))
	for _, value := range slice {
		f(value)

	}
}
