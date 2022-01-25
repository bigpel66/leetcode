// 941. Valid Mountain Array
//
// https://leetcode.com/problems/valid-mountain-array/

// validMountainArray function checks the array is formed like mountain or not.
// Mountain consists of increasing range, and decreasing range.
// There's no plain on the mountain.
func validMountainArray(arr []int) bool {
	top := 0
	for top = 0; top < len(arr)-1; top++ {
		if arr[top] >= arr[top+1] {
			break
		}
	}
	if top == 0 || top == len(arr)-1 {
		return false
	}
	for ; top < len(arr)-1; top++ {
		if arr[top] <= arr[top+1] {
			return false
		}
	}
	return true
}
