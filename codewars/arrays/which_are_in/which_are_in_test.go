package which_are_in

import (
	"testing"
)

func TestWhichAreIn(t *testing.T) {
	testsArrOne := [][]string{
		[]string{"live", "arp", "strong"},
		[]string{"cod", "code", "wars", "ewar", "ar"},
	}

	testsArrTwo := [][]string{
		[]string{"lively", "alive", "harp", "sharp", "armstrong"},
		[]string{},
	}

	results := [][]string{
		[]string{"arp", "live", "strong"},
		[]string{},
	}

	for i := 0; i < len(testsArrOne); i++ {
		ret := InArray(testsArrOne[i], testsArrTwo[i])

		for idx, substr := range ret {
			if substr != results[i][idx] {
				t.Fatalf("case %d fails: %v\n", i, ret)
			}
		}
	}
}
