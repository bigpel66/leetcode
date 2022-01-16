// 849. Maximize Distance to Closest Person
//
// https://leetcode.com/problems/maximize-distance-to-closest-person/

// max function finds the maximum value from the given inputs.
func max(nums ...int) int {
	val := -int(^uint(0)>>1) - 1
	for _, n := range nums {
		if val < n {
			val = n
		}
	}
	return val
}

// maxDistToClosest function finds the maximum distance to the closest person from the given seats.
// 0 0 0 1 -> value should be 3 which is the number of 0s on the left side.
// 1 0 0 0 0 -> value should be 4 which is the number of 0s on the right side.
// 0 1 0 0 0 0 0 0 1 0 0 -> value should be 3 which is the number of zeros in the middle of the each side 1s.
// Thus, check the 0s on the each side. (Temporarilly, one of those will be the maximum distance.)
// Also, check the 0s on the middle area. This procedure only needs the sequential search to count the 0s. The number of 0s divided by 2 should be compared to the previous maximum value and update it.
// On searching the middle area, the skipping 1s and checking 0s have been used.
func maxDistToClosest(seats []int) int {
	left, right := 0, 0
	for left < len(seats) && seats[left] == 0 {
		left++
	}
	for right < len(seats) && seats[len(seats)-1-right] == 0 {
		right++
	}
	answer := max(left, right)
	for begin, end := left, len(seats)-1-right; begin < end; begin++ {
		count := 0
		for ; begin < end && seats[begin] == 1; begin++ {
		}
		for ; begin < end && seats[begin] == 0; begin++ {
			count++
		}
		answer = max(answer, (count+1)/2)
	}
	return answer
}
