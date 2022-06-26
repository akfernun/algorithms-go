package main

import (
	"algorithms-go/dynamicProgramming"
	"fmt"
)

func main() {
	results := dynamicProgramming.Partition("ababbbabbaba")

	fmt.Println(results)
}
