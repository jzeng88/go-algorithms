package highest_rank_num

import "testing"

func TestHighestRank(t *testing.T) {
	test := []int{12, 10, 8, 12, 7, 6, 4, 10, 12, 10}
	result := 12

	ret := HighestRank(test)

	if ret != result {
		t.Fatalf("case fails: %v\n", ret)
	}
}
