package alphabet_symmetry

import "testing"

func TestSolve(t *testing.T) {
	tests := [][]string{
		[]string{"abode", "ABc", "xyzD"},
		[]string{"abide", "ABc", "xyz"},
		[]string{"IAMDEFANDJKL", "thedefgh", "xyzDEFghijabc"},
		[]string{"encode", "abc", "xyzD", "ABmD"},
	}

	results := [][]int{
		[]int{4, 3, 1},
		[]int{4, 3, 0},
		[]int{6, 5, 7},
		[]int{1, 3, 1, 3},
	}

	for i := 0; i < len(tests); i++ {
		ret := Solve(tests[i])

		for idx, count := range ret {
			if count != results[i][idx] {
				t.Fatalf("Expected %d but received %d", results[idx], count)
			}
		}
	}
}
