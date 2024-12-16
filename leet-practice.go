package main

func twoSum(nums []int, target int) []int {
	var response []int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if i+1 < len(nums) {
				if nums[i]+nums[j] == target {
					response = append(response, i)
					response = append(response, j)
					break
				}
			}
		}
	}

	return response
}
