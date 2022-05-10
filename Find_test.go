package funcional

import (
	"testing"
)

func TestFind(t *testing.T) {
	find, ok := Find([]int{1, 2, 3, 4}, func(value int) bool {
		return value == 5
	})

	// si se dice que se encontro un valor y el valor encontrado no es igual a 5 lanza un error
	if ok && find != 5 {
		t.Errorf("Se esperaba un 5, pero es %d", find)
	}
}
