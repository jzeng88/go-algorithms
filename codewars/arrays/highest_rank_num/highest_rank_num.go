// https://www.codewars.com/kata/5420fc9bb5b2c7fd57000004/go
package highest_rank_num

func HighestRank(nums []int) int {
	count := make(map[int]int)

	for _, num := range nums {
		count[num]++
	}

	var highestCount int
	var highestKey int

	for key, totalCount := range count {
		if totalCount > highestCount {
			highestCount = totalCount
			highestKey = key
		} else if totalCount == highestCount {
			if highestKey < key {
				highestKey = key
			}
		}
	}

	return highestKey
}
