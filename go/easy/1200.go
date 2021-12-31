// 1200. Minimum Absolute Difference
//
// https://leetcode.com/problems/minimum-absolute-difference/

// To get the minimum difference, sorting is necessary.
// Package sort supports the O(nlogn) sorting algorithm.
import (
	"sort"
)

// minimumAbsDifference function finds some pairs of two integer elements which makes the  minimum difference of the given array.
// The function name implies the absolute value, but only the minimum difference is the hint.
// If the array is maintained with the ascending orders, the minimum difference can be found by a loop which calculates the value of the i + 1 th element - the value of the i th element.
// When a new difference is smaller than the previous one, make a new slices which holds the the vallue of the i th element, the value of the i + 1 element.
// If a difference is as same as the previous one, just append the pair into the original slices.
// Repeat to the end of the given array.
func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	result := [][]int{{arr[0], arr[1]}}
	value := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		if value > diff {
			value = diff
			result = [][]int{{arr[i], arr[i+1]}}
		} else if value == diff {
			result = append(result, []int{arr[i], arr[i+1]})
		}
	}
	return result
}
