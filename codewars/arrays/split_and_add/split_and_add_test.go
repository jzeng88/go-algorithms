package split_and_add

import "testing"

func TestSplitAndAdd(t *testing.T) {
	testsFirstArg := [][]int{
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2, 3, 4, 5},
		[]int{15},
		[]int{1, 2, 3, 4},
		[]int{1, 2, 3, 4, 5, 6},
		[]int{32, 45, 43, 23, 54, 23, 54, 34},
		[]int{32, 45, 43, 23, 54, 23, 54, 34},
		[]int{3, 234, 25, 345, 45, 34, 234, 235, 345},
		[]int{3, 234, 25, 345, 45, 34, 234, 235, 345, 34, 534, 45, 645, 645, 645, 4656, 45, 3},
		[]int{23, 345, 345, 345, 34536, 567, 568, 6, 34536, 54, 7546, 456},
	}

	testsSecondArg := []int{
		2,
		3,
		3,
		1,
		20,
		2,
		0,
		3,
		4,
		20,
	}

	results := [][]int{
		[]int{5, 10},
		[]int{15},
		[]int{15},
		[]int{4, 6},
		[]int{21},
		[]int{183, 125},
		[]int{32, 45, 43, 23, 54, 23, 54, 34},
		[]int{305, 1195},
		[]int{1040, 7712},
		[]int{79327},
	}

	for idx, test := range testsFirstArg {
		ret := SplitAndAdd(test, testsSecondArg[idx])

		for i, v := range ret {
			if v != results[idx][i] {
				t.Fatalf("Expected %d but received %d", results[i], v)
			}
		}
	}
}
