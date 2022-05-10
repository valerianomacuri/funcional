package funcional

import (
	"testing"
)

func TestFilter(t *testing.T) {
	// filtrar los numeros diferente de uno
	f := Filter([]int{1, 2, 3, 4}, func(num int) bool {
		if num == 1 {
			return false
		} else {
			return true
		}
	})
	if len(f) != 3 {
		t.Errorf("Se esperaba que la longitud f sea igual de 3, pero es %d", len(f))
	}
}
