package duplicate_encoder

import "testing"

func TestDuplicateEncode(t *testing.T) {
	tests := []string{
		"din",
		"recede",
		"(( @",
	}

	results := []string{
		"(((",
		"()()()",
		"))((",
	}

	for i := 0; i < len(tests); i++ {
		if ret := DuplicateEncode(tests[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
