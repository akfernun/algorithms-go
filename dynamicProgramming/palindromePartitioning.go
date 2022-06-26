package dynamicProgramming

// https://leetcode.com/problems/palindrome-partitioning/
func Partition(s string) [][]string {
	var results [][]string
	cache := buildPalindromicTable(s)

	doPartition(s, 0, nil, &results, cache)

	return results
}

func doPartition(s string, start int, partitions []string, results *[][]string, cache [][]bool) {
	if start == len(s) {
		result := make([]string, len(partitions))
		copy(result, partitions)

		*results = append(*results, result)
	}

	for end := start; end < len(s); end++ {
		// check to see if substring at s[start:end] is a palindrome
		if cache[start][end] {
			partitionsHere := append(partitions, s[start:end+1])

			doPartition(s, end+1, partitionsHere, results, cache)
		}
	}
}

// builds a 2D table that contains palindromic info for all substrings
// doing this in a separate step/loop breaks the algo into steps and is
// thus a little easier to understand as opposed to putting everything into
// one huge function/loop/etc.
func buildPalindromicTable(s string) [][]bool {
	table := make([][]bool, len(s))

	for idx := range table {
		table[idx] = make([]bool, len(s))
	}

	for startColumn := range table[0] {
		for row, column := 0, startColumn; column < len(table[0]); row, column = row+1, column+1 {
			if s[row] == s[column] {
				table[row][column] = column-row <= 1 || table[row+1][column-1]
			}
		}
	}

	return table
}
