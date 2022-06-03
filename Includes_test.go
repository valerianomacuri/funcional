package funcional

import (
	"testing"
)

func TestIncludes(t *testing.T) {

	slice := []int{1, 3, 4, 5}

	result := Includes(slice, 1)
	// slice includes 1, then result is true
	if result != true {
		t.Errorf("Se esperaba que %+v incluya 1", slice)
	}

	result = Includes(slice, 6)
	// slice don't includes 6, then result is false
	if result == true {
		t.Errorf("Se esperaba que %+v no incluya 6", slice)
	}
}
