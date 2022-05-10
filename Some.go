package funcional

func Some[T1 any](slice []T1, fn func(value T1) bool) bool {
	some := false
	for _, value := range slice {
		if fn(value) {
			some = true
		}
	}
	return some
}
