// 1094. Car Pooling
//
// https://leetcode.com/problems/car-pooling/

// carPooling function finds whether the car pooling is possible or not.
// [2, 1, 5] means 2 people get on the station 1, and they are dropped by on the station 5.
// There are only 1000 stations at most, just record how many people are incremented or decremented at each station on the slice.
// Only O(n) is required both time and memory complexity.
func carPooling(trips [][]int, capacity int) bool {
	station := make([]int, 1001)
	for _, t := range trips {
		station[t[1]] += t[0]
		station[t[2]] -= t[0]
	}
	for onboard, i := 0, 0; i < 1001; i++ {
		onboard += station[i]
		if onboard > capacity {
			return false
		}
	}
	return true
}
