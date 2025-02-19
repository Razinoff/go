package runningsum

func runningSum(nums []int) []int {
	var res []int
	num := 0
	for _, value := range nums {
		num += value
		res = append(res, num)
	}
	return res
}
