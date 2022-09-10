package stringProblems
//https://leetcode.com/problems/isomorphic-strings/

func IsIsomorphic(s string, t string) bool {
	return helper(s, t) && helper(t, s)
}

func helper(s, t string) bool {
	transformations := map[byte]byte {}

	for i, char := range s {
		byteChar := byte(char)

		if transformationVal, ok := transformations[byteChar]; ok && transformationVal != t[i] {
			 return false
		}

		transformations[byteChar] = t[i]
	}

	return true
}
