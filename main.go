package main

import (
	"algorithms-go/dynamicProgramming/wordBreak"
	"fmt"
)

func main() {
	wordDict := []string { "leet", "code" }
	result := wordBreak.WordBreakBottomUp("leetcode", wordDict)

	fmt.Println(result)
}
