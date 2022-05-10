package funcional

import (
	"testing"
)

func TestSome(t *testing.T) {

	slice := []int{1, 3, 4, 5}

	result := Some(slice, func(value int) bool {
		return (value % 2) == 0
	})
	//if is not even (par) lanza un error
	if result != true {
		t.Errorf("Se esperaba que sea par, pero se obtuvo impar")
	}
}
