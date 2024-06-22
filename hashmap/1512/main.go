package main

// not optimized
func numIdenticalPairs(nums []int) int {
	count := 0
	for i, val := range nums {
		for j := 1; j < len(nums); j++ {
			if val == nums[j] && i < j {
				count += 1
			}
		}
	}
	return count
}
