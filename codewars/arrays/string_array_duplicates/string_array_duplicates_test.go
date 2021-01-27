package string_array_duplicates

import "testing"

func TestDup(t *testing.T) {
	test := []string{"ccooddddddewwwaaaaarrrrsssss", "piccaninny", "hubbubbubboo"}
	result := []string{"codewars", "picaniny", "hubububo"}

	ret := Dup(test)

	for idx, str := range ret {
		if str != result[idx] {
			t.Fatalf("case %d fails: %v\n", idx, ret)
		}
	}
}
