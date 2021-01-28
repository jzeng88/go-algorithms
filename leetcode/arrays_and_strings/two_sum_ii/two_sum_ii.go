// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
package two_sum_ii

// Two Pointer approach:
// O(N) time & O(1) space
func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1

	for i < j {
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		} else if numbers[i]+numbers[j] < target {
			i++
		} else {
			j--
		}
	}
	return []int{-1, -1}
}

// Hashtable Approach:
// O(N) time & space
func twoSum1(numbers []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(numbers); i++ {
		another := target - numbers[i]
		if _, ok := m[another]; ok {
			return []int{m[another] + 1, i + 1}
		}
		m[numbers[i]] = i
	}
	return nil
}
