package binarySearch

func FindPeakElement(nums []int) int {
	return find(nums, 0, len(nums) - 1)
}

func find (nums []int, start, end int) int {
	// if there's only one element, the element *is* the peak
	if start == end {
		return start
	}

	// if there are only two elements in the range, find peak and return
	if end - start == 1 {
		if nums[start] > nums[end] {
			return start
		} else {
			return end
		}
	}

	// check the middle element and if it's the peak, return its index,
	// otherwise recurse
	middle := (start + end) / 2
	if nums[middle] > nums[middle + 1] && nums[middle] > nums[middle - 1] {
		return middle
	} else if nums[middle + 1] > nums[middle] {
		return find(nums, middle + 1, end)
	} else {
		return find(nums, start, middle - 1)
	}
}
