package funcional

import "testing"

func TestMap(t *testing.T) {

	result := Map([]int{1, 2, 3, 4, 5}, func(value int) int {
		return value + 1
	})

	if result[0] != 2 {
		t.Errorf("Se esperaba que result[0] sea igual a 2, pero es %d", result[0])
	}
}
