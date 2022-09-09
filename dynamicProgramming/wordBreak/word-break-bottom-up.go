package wordBreak
// https://leetcode.com/problems/word-break/

import (
	"strings"
)

// Adds a simple bottom-up solution to the word-break problem
func WordBreakBottomUp(s string, wordDict []string) bool {
	cache := make([]bool, len(s) + 1, len(s) + 1)
	cache[len(cache) - 1] = true

	for k := len(s) - 1; k >= 0 ; k-- {
		str := s[k:]

		for _, dictionaryWord := range wordDict {
			hasPrefix := strings.HasPrefix(str, dictionaryWord)

			if hasPrefix && cache[k + len(dictionaryWord)] {
				cache[k] = true
				break
			}
		}
	}

	return cache[0]
}
