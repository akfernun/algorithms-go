package dynamicProgramming

//https://leetcode.com/problems/triangle/
// this version modifies the original 2d array
// to obtain the result. 
//func MinimumTotal(triangle [][]int) int {
//	if len(triangle) == 1 {
//		return triangle[0][0]
//	}
//
//	for k := len(triangle) - 2; k >= 0; k-- {
//		for ind, h := range triangle[k] {
//			triangle[k][ind] = h + min(triangle[k + 1][ind], triangle[k + 1][ind + 1])
//		}
//	}
//
//	return triangle[0][0]
//}

// MinimumTotal Per a specification of the problem, this solution
// uses O(n) space where O(n) is the number of rows in the triangle.
// I'm assuming the problem added this  constraint in an
// effort to prevent the developer from modifying the
// original triangle input array. 🤷‍️
func MinimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	lastRow := triangle[len(triangle) - 1]

	cache := make([]int, len(lastRow), len(lastRow))
	copy(cache, lastRow)

	for k := len(triangle) - 2; k >= 0; k-- {
		for r, el := range triangle[k] {
			cache[r] = el + min(cache[r], cache[r + 1])
		}
	}


	return cache[0]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}