package ip_validation

import "testing"

func TestValidateIp(t *testing.T) {
	tests := []string{
		"12.255.56.1",
		"",
		"abc.def.ghi.jkl",
		"123.456.789.0",
		"12.34.56",
		"12.34.56 .1",
	}

	results := []bool{
		true, false, false, false, false, false,
	}

	for idx, test := range tests {
		if ret := ValidateIp(test); ret != results[idx] {
			t.Fatalf("Expected %v but received %v", results[idx], ret)
		}
	}
}
