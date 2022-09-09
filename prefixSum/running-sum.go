package prefixSum

func runningSum(nums []int) []int {
	runningSums := make([]int, len(nums), len(nums))
	runningSums[0] = nums[0]

	for k := 1; k < len(nums); k++ {
		runningSums[k] = nums[k] + runningSums[k - 1]
	}

	return runningSums
}
