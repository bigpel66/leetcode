// 605. Can Place Flowers
//
// https://leetcode.com/problems/can-place-flowers/

// canPlaceFlowers function checks whether planting n flowers is possible or not.
// Adjacent planting is not allowed.
// Just iterating the blocks on the flowerbed, and check the empty field whether planting is possible or not.
func canPlaceFlowers(flowerbed []int, n int) bool {
	flowerbed = append(append([]int{0}, flowerbed...), 0)
	for i := 1; i < len(flowerbed)-1; i++ {
		if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			n--
		}
	}
	return n <= 0
}
