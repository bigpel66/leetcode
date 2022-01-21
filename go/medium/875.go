// 875. Koko Eating Bananas
//
// https://leetcode.com/problems/koko-eating-bananas/

// math package has been imported to call Ceil.
// sort package has been imported to call Search for Binary Search.
import (
	"math"
	"sort"
)

// minEatingSpeed function finds the minimum portion to eat bananas in h hours.
// Parametric search has been used and this is done by Search function of sort package.
func minEatingSpeed(piles []int, h int) int {
	max := 0
	for _, p := range piles {
		if max < p {
			max = p
		}
	}
	k := sort.Search(max, func(x int) bool {
		if x == 0 {
			return len(piles) == 0
		}
		sum := 0
		for _, p := range piles {
			sum += int(math.Ceil(float64(p) / float64(x)))
		}
		return sum <= h
	})
	return k
}
