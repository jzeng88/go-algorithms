package highest_scoring_word

import "testing"

func TestHigh(t *testing.T) {
	test := "man i need a taxi up to ubud"
	result := "taxi"

	ret := High(test)

	if ret != result {
		t.Fatalf("Expected %s but received %s", result, ret)
	}
}
