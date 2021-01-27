package split_strings

import "testing"

func TestSolution(t *testing.T) {
	tests := []string{
		"abc",
		"dawsd",
		"awsaws",
	}

	results := [][]string{
		[]string{"ab", "c_"},
		[]string{"da", "ws", "d_"},
		[]string{"aw", "sa", "ws"},
	}

	for i, test := range tests {
		ret := Solution(test)

		for j := 0; j < len(ret); j++ {
			if ret[j] != results[i][j] {
				t.Fatalf("Expected %s but received %s", results[i][j], ret[j])
			}
		}
	}
}
