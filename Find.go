package funcional

func Find[T any](s []T, f func(T) bool) (find T, ok bool) {
	for _, value := range s {
		if f(value) {
			ok = true
			find = value
			break
		}
	}
	return find, ok
}
