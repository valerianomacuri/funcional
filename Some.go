package funcional

// Some returns true if a value was found with the condition passed as a function
func Some[T1 any](slice []T1, fn func(value T1) bool) bool {
	some := false
	for _, value := range slice {
		if fn(value) {
			some = true
		}
	}
	return some
}
