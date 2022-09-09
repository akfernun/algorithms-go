package wordBreak
// https://leetcode.com/problems/word-break/

import (
	"strings"
)

// This is a simple top-down DP solution in which I iterate
// through the dictionary and see if the dict word matches the
// beginning of the string. If it does, recurse, else keep going
func wordBreak(s string, wordDict []string) bool {
	return helper(s, map[string]bool{}, wordDict)
}

func helper (s string, cache map[string]bool, wordDict []string) bool {
	if s == "" {
		return true
	}

	if val, ok := cache[s]; ok {
		return val
	}

	for _, prefix := range wordDict {
		if strings.HasPrefix(s, prefix) {
			if result := helper(s[len(prefix):], cache, wordDict); result {
				cache[s] = true

				return true
			}
			cache[s] = false
		}
	}

	cache[s] = false
	return false
}
