package stringProblems

func isSubsequence(s string, t string) bool {
	subsequencePtr := 0

	for k := 0; k < len(t) && subsequencePtr < len(s); k++ {
		if s[subsequencePtr] == t[k] {
			subsequencePtr++
		}
	}

	return subsequencePtr == len(s)
}
