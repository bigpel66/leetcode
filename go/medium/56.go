// 56. Merge Intervals
//
// https://leetcode.com/problems/merge-intervals/

// sort package imported to order the interval slices
import (
	"sort"
)

// merge function literally merge the intervals.
// [[1, 3], [2, 6], [8, 10], [15, 18]] will be [[1, 6], [8, 10], [15, 18]].
// To merge, compare the start and the end should be easy. Sorting with the start point is necessary.
// After that, just compare the following if statement and update.
// No need to use a new [][]int, but adopted the way to make the code clear.
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	res := [][]int{}
	for i, interval := range intervals {
		if i == 0 {
			res = append(res, interval)
			continue
		}
		temp := res[len(res)-1]
		if temp[1] >= interval[0] && temp[1] < interval[1] {
			res[len(res)-1] = []int{temp[0], interval[1]}
		} else if temp[1] < interval[0] {
			res = append(res, interval)
		}
	}
	return res
}
