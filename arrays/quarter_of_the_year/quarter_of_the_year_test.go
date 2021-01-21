package quarter_of_the_year

import "testing"

func TestQuarterOf(t *testing.T) {
	tests := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	results := []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4}

	for i := 0; i < len(tests); i++ {
		if ret := QuarterOf(tests[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
