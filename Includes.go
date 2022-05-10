package funcional

func Includes[T1 comparable](slice []T1, value T1) bool {
	includes := false
	for _, v := range slice {
		if v == value {
			includes = true
			break
		}
	}
	return includes
}
