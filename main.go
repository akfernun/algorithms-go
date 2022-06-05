package main

import (
	binarySearch "algorithms-go/binarySearch"
	"fmt"
)

func main() {
	nums := []int{1,2,3,4,5}
	result := binarySearch.FindPeakElement(nums)

	fmt.Println(result)
}
