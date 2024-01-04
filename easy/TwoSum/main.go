package main

func twoSum(nums []int, target int) []int {
	var result []int = make([]int, 2)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result[0], result[1] = i, j
				break
			}
		}
	}

	return result
}
