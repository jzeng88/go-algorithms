package longest_consecutive_rep

import "testing"

type Result struct {
	C rune
	L int
}

func TestLongestRepitition(t *testing.T) {

	tests := []string{
		"aaaabb",
		"bbbaaabaaaa",
		"cbdeuuu900",
		"abbbbb",
		"aabb",
		"ba",
		"",
	}

	results := []Result{
		Result{'a', 4},
		Result{'a', 4},
		Result{'u', 3},
		Result{'b', 5},
		Result{'a', 2},
		Result{'b', 1},
		Result{},
	}

	for idx, test := range tests {
		if ret := LongestRepetition(test); ret != results[idx] {
			t.Fatalf("Expected %d but received %d", results[idx], ret)
		}
	}
}
