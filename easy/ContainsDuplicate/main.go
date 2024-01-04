package main

func containsDuplicate(nums []int) bool {
	pairs := make(map[int]struct{})

	for _, num := range nums {
		_, ok := pairs[num]
		if !ok {
			pairs[num] = struct{}{}
		} else {
			return true
		}
	}

	return false
}
