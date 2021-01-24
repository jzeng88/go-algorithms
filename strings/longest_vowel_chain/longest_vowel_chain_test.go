package longest_vowel_chain

import "testing"

func TestSolve(t *testing.T) {
	tests := []string{
		"codewarriors",
		"suoidea",
		"ultrarevolutionariees",
		"strengthlessnesses",
		"cuboideonavicuare",
		"chrononhotonthuooaos",
		"iiihoovaeaaaoougjyaw",
	}

	results := []int{
		2,
		3,
		3,
		1,
		2,
		5,
		8,
	}

	for i := 0; i < len(tests); i++ {
		if ret := Solve(tests[i]); ret != results[i] {
			t.Fatalf("Expected %d but received %d", results[i], ret)
		}
	}
}
