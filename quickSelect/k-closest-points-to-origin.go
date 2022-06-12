package quickSelect

import (
	"math"
)

type Point = struct {
	point    []int
	distance float64
}

//https://leetcode.com/problems/k-closest-points-to-origin/
func KClosest(points [][]int, k int) [][]int {
	pointsWithDistance := make([]Point, len(points), len(points))

	for idx, point := range points {
		xVal := math.Pow(float64(point[0])-0, 2)
		yVal := math.Pow(float64(point[1])-0, 2)

		pointWithDistance := Point{distance: math.Sqrt(xVal + yVal), point: point}

		pointsWithDistance[idx] = pointWithDistance
	}

	// now that we have all the distances mapped out,
	// we can run a quick select algorithm to select the
	// k-smallest values
	results := quickSelect(0, len(pointsWithDistance)-1, k, pointsWithDistance)

	var closest [][]int

	for _, result := range results {
		closest = append(closest, result.point)
	}

	return closest
}

func quickSelect(start, end, k int, distances []Point) []Point {
	middle := -1

	for middle+1 != k {
		middle = partition(start, end, distances)
		if middle+1 > k {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}

	return distances[:k]
}

func partition(start, end int, distances []Point) int {
	pivotEl := distances[end].distance
	part := start

	for i := start; i < end; i++ {
		distance := distances[i].distance

		if distance <= pivotEl {
			distances[i], distances[part] = distances[part], distances[i]
			part += 1
		}
	}

	distances[part], distances[end] = distances[end], distances[part]

	return part
}
