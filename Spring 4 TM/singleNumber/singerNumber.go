package singlenumber

import "sort"

func singleNumber(nums []int) int {
	sort.Ints(nums)
	i := 0
	for i < len(nums)-1 {
		i += 2
		if nums[i-2] != nums[i-1] {
			return nums[i-2]
		}
	}
	return nums[len(nums)-1]
}
