// https://leetcode.com/problems/two-sum/
package two_sum

// O(n) time & space
// Use a hash to store the difference as key & idx as value
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, i}
		}
		m[v] = i
	}
	return nil
}
