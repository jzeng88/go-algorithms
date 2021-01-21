// https://www.codewars.com/kata/5715eaedb436cf5606000381

package sum_of_positives

func PositiveSum(nums []int) int {
	sum := 0

	for _, v := range nums {
		if v > 0 {
			sum += v
		}
	}

	return sum
}
