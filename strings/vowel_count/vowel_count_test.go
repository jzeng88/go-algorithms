package vowel_count

import "testing"

func TestGetCount(t *testing.T) {
	test := "abracadabra"
	result := 5

	if ret := GetCount(test); ret != result {
		t.Fatalf("Expected %d but received %d", result, ret)
	}
}
