// 1345. Jump Game IV
//
// https://leetcode.com/problems/jump-game-iv/

// minJumps function finds the number of jumps to reach finish point.
// The player can jump to 1 step backward and forward, the point which number is same.
// Thus, there are 3 cases for each step. This will be solved by the BFS Queueing.
// For each queueing is for a 1 step, and checks whether reaching to the final point is possible or not.
// If it is, return the steps. Or enque the possible 3 cases from the point.
// The queued 3 cases will be processed on the next step naturally (the cases on the same step will be proessed first due to the queue).
// Not to visit the duplicated point, visitd slice has been used.
// Also, for jumping to the same value, position hash map has been used.
// The position map always have to be resovled only once the value entered. Thus, resolved hash map has been used.
func minJumps(arr []int) int {
	steps := 0
	queue := append([]int{}, 0)
	visited := make([]bool, len(arr))
	resolved := make(map[int]bool)
	position := make(map[int][]int)
	visited[0] = true
	for i, v := range arr {
		if _, ok := position[v]; !ok {
			position[v] = make([]int, 0)
		}
		position[v] = append(position[v], i)
	}
	for len(queue) > 0 {
		sizeOfSameSteps := len(queue)
		for sizeOfSameSteps > 0 {
			index := queue[0]
			queue = queue[1:]
			if index == len(arr)-1 {
				return steps
			}
			if !resolved[arr[index]] {
				resolved[arr[index]] = true
				for _, i := range position[arr[index]] {
					if index != i {
						queue = append(queue, i)
						visited[i] = true
					}
				}
			}
			if index+1 < len(arr) && !visited[index+1] {
				queue = append(queue, index+1)
				visited[index+1] = true
			}
			if index-1 >= 0 && !visited[index-1] {
				queue = append(queue, index-1)
				visited[index-1] = true
			}
			sizeOfSameSteps--
		}
		steps++
	}
	return -1
}
