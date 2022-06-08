package binarySearch

// SearchMatrix
// The idea behind this algorithm is relatively straight-forward.
// First, we're going to use binary search to find what row we
// should be looking in. Once we've found a row we're interested in,
// we perform a binary search on the row to see if the element we're
// looking for is in there.
func SearchMatrix(matrix [][]int, target int) bool {
	rowToSearch := findRow(matrix, target, 0, len(matrix) - 1)

	if rowToSearch == -1 {
		return false
	}

	row := matrix[rowToSearch]

	return findEl(row, 0, len(matrix[rowToSearch]) - 1, target)
}


// find the row that the element *might* be in. Once we find the
// row, we'll perform a binary search on it.
func findRow (matrix [][]int, target, start, end int) int {
	if start > end {
		return -1
	}

	rowIdx := (start + end) / 2
	rowStart := matrix[rowIdx][0]
	rowEnd := matrix[rowIdx][len(matrix[0]) - 1]

	// we found the row we should be searching
	if target >= rowStart && target <= rowEnd {
		return rowIdx
	} else if target < rowStart {
		return findRow(matrix, target, start, rowIdx - 1)
	} else {
		return findRow(matrix, target, rowIdx + 1, end)
	}
}

// Perform binary search to find element within a given row
func findEl(row []int, start, end, target int) bool  {
	if start > end {
		return false
	}

	middle := (start + end) / 2
	middleEl := row[middle]

	if target == middleEl {
		return true
	} else if target < middleEl {
		return findEl(row, start, middle - 1, target)
	} else {
		return findEl(row, middle + 1, end, target)
	}

	return false
}
