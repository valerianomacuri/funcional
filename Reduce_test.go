package funcional

import (
	"testing"
)

func TestReduce(t *testing.T) {

	slice := []int{1, 2, 3, 4, 5}

	result := Reduce(slice, 0, func(num1, num2 int) int {
		return num1 + num2
	})

	if result != 15 {
		t.Errorf("Se esperaba que la suma de sus elementos %+v sea 15", slice)
	}
}
