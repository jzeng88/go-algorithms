package unique_num

import "testing"

func TestUniqNum(t *testing.T) {
	tests := [][]float32{
		[]float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0},
		[]float32{0, 0, 0.55, 0, 0},
	}

	results := []float32{2.0, 0.55}

	for i := 0; i < len(tests); i++ {
		if ret := FindUniq(tests[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
