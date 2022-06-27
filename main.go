package main

import (
	"algorithms-go/dynamicProgramming"
	"fmt"
)

func main() {
	triangle := [][]int {{2},{3,4},{6,5,7},{4,1,8,3}}
	//triangle := [][]int {{-10}}

	result := dynamicProgramming.MinimumTotal(triangle)

	fmt.Println(result)
}
