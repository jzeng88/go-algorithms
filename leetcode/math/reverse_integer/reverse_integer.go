// https://leetcode.com/problems/reverse-integer/
package reverse_integer

// 123 => 0 + 3 => 30 + 2 = 320 + 1 = 321
// -123 = 0 - 3 => -30 - 2 = -320 - 1 = -321
// 120 = 21
// 0 = 0

func reverse(x int) int {
	tmp := 0

	for x != 0 {
		tmp = tmp*10 + x%10
		x = x / 10
	}

	// Handles overflow/underflow for 32-bit integer
	// "1 << 31-1" === 1 times 2, 30 times, 2^31-1
	if tmp > 1<<31-1 || tmp < -(1<<31) {
		return 0
	}

	return tmp
}
