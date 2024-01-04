package main

func removeElement(nums []int, val int) ([]int, int) {
	var result, counter int

	for i, n := range nums {
		if n != val {
			result++
		} else {
			nums[i] = 0
		}
	}
	for i, n := range nums {
		if n != 0 {
			nums[i] = 0
			nums[counter] = n
			counter++
		}
	}

	return nums, result
}
