package main

import (
	"algorithms-go/binarySearch"
	"fmt"
)

func main() {
	matrix := [][]int { {1} }

	result := binarySearch.SearchMatrix(matrix, 1)

	fmt.Println(result)
}
