package sum_of_positives

import "testing"

func TestPositiveSum(t *testing.T) {
	tests := [][]int{
		[]int{-1, 5, 8, -2},
		[]int{1, 8, 0},
	}

	results := []int{
		13, 9, 10,
	}

	for i := 0; i < len(tests); i++ {
		if ret := PositiveSum(tests[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
