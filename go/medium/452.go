// 452. Minimum Number of Arrows to Burst Ballons
//
// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/

// sort package has been imported to sort the slice of int slice.
import (
	"sort"
)

// findMinArrowShots function finds the minimum number of arrows to burst all of the balloons.
// Sort all of the points by the X_start or X_end. (In this case, they are sorted by X_end.)
// Then, regardless of the other position, the truth that a specific point is overlapped or not can be determined.
// If X_end of the point A is less than X_start of the point B, they are not overlapped.
// If the points are not overlapped, then increase the count because they are not bustable by an arrow. Also, update the criteria to determine whether the other points are overlapped or not.
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	answer := 0
	pivot := -int(^uint(0)>>1) - 1
	for _, p := range points {
		if p[0] > pivot {
			answer++
			pivot = p[1]
		}
	}
	return answer
}
