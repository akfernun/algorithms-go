package prefixSum


//https://leetcode.com/problems/find-pivot-index/

func PivotIndex(nums []int) int {
	totalSum := sum(nums)
	leftSum := 0

	for index, num := range nums {
		rightSum := totalSum - num - leftSum

		if rightSum == leftSum {
			return index
		}

		leftSum += num
	}

	return -1
}

func sum(nums []int) int {
	var totalSum int
	for _, num := range nums {
		totalSum = totalSum + num
	}

	return totalSum
}
