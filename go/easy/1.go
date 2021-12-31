// 1. Two Sum
//
// https://leetcode.com/problems/two-sum/

// twoSum function finds indice of the two elements to make target.
// Target number can be found only once.
// Mapping to make target is necessary.
func twoSum(nums []int, target int) []int {
	cache := map[int]int{}
	for i, v := range nums {
		if j, ok := cache[target-v]; ok {
			return []int{j, i}
		}
		cache[v] = i
	}
	return []int{}
}
