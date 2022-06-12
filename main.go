package main

import (
	"algorithms-go/quickSelect"
	"fmt"
)

func main() {
	points := [][]int{{1, 3}, {-2, 2}}
	k := 1
	result := quickSelect.KClosest(points, k)

	fmt.Println(result)
}
