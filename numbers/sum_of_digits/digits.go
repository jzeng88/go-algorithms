// https://www.codewars.com/kata/541c8630095125aba6000c00/
package sum_of_digits

import (
	"strconv"
	"strings"
)

func DigitalRoot(n int) int {
	sum := n

	for len(strconv.FormatInt(int64(sum), 10)) != 1 {
		strNum := strconv.FormatInt(int64(sum), 10)
		strNumArr := strings.Split(strNum, "")
		sum = 0

		for _, numStr := range strNumArr {
			num, _ := strconv.Atoi(numStr)
			sum += num
		}
	}

	return sum
}
