// https://leetcode.com/problems/plus-one/
package plus_one

// digits = [1,2,9]
// carry = 1			carry = 0		 carry = 0
//       |						|					|
// 1, 2, 9+1   =>  1, 3, 0 	=> 	1, 3, 0 	=> 	1, 3, 0

// digits = [0, 2]
// carry = 0			carry = 0
// 	  |						 |
// 0, 1+1			 =>  0, 2

// O(n) time, O(1) space
func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{}
	}

	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+carry > 9 {
			digits[i] = 0
			carry = 1
		} else {
			digits[i] += carry
			carry = 0
		}
	}

	// handle the [0] case
	if digits[0] == 0 && carry == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
