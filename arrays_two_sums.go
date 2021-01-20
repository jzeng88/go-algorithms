// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

package main

import "fmt"

// O(n) time complexity
// O(n) space complexity
func twoSum(nums []int, target int) []int {
	// Create hash table
	m := make(map[int]int)

	// Iterate through nums slice
	for i := 0; i < len(nums); i++ {

		diff := target - nums[i]

		// If difference exist in hash table
		if _, ok := m[diff]; ok {
			// return the index of the difference and current index
			return []int{m[diff], i}
		}

		// store current slice value and its index
		m[nums[i]] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{1,2,3,4,5}, 9))
}