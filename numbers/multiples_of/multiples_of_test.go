package multiples

import "testing"

func TestMultiples3And5(t *testing.T) {
	test := 10
	result := 23

	if ret := Multiple3And5(test); ret != result {
		t.Fatalf("Expected %d but received %d", result, ret)
	}
}
