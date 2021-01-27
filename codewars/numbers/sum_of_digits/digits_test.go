package sum_of_digits

import "testing"

func TestDigitalRoot(t *testing.T) {
	tests := []int{
		16,
		195,
		992,
		167346,
		0,
	}

	results := []int{
		7, 6, 2, 9, 0,
	}

	for idx, test := range tests {
		if ret := DigitalRoot(test); ret != results[idx] {
			t.Fatalf("Expected %d but received %d", results[idx], ret)
		}
	}
}
