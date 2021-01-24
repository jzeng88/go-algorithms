// https://www.codewars.com/kata/5946a0a64a2c5b596500019a/train/go
package split_and_add

func SplitAndAdd(numbers []int, n int) []int {
	if n <= 0 || len(numbers) == 1 {
		return numbers
	}

	midIdx := len(numbers) / 2
	left, right := numbers[:midIdx], numbers[midIdx:]
	result := []int{}

	if len(left) == len(right) {
		for idx, num := range left {
			result = append(result, num+right[idx])
		}
	} else if len(left) < len(right) {
		result = right[:len(numbers)%2]

		for idx, num := range right[len(numbers)%2:] {
			result = append(result, num+left[idx])
		}
	}

	return SplitAndAdd(result, n-1)
}
