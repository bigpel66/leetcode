// 1291. Sequential Digits
//
// https://leetcode.com/problems/sequential-digits/

// sequentialDigits function finds the numbers which digits are sequential.
// 1 -> 12 -> 123 -> 1234 ...
// 2 -> 23 -> 234 -> 2345 ...
// 3 -> 34 -> 345 -> 3456 ...
// To make the above way, DFS or BFS could be the solution.
// Solving by BFS could make the number to execute in ascending order.
// By pushing 1 to 9 into the queue, 1 -> 2 -> 3 -> ... -> 9 -> 12 -> 23 -> 34 ... -> 123 -> 234 ... will be done.
// Not to make mess cases, appending into the queue should be done when the digit is less than 10.
// The variable high could be 10^9, and it is bigger than 123456789. Thus, no length check of queue could cause runtime error.
func sequentialDigits(low int, high int) []int {
	answer := []int{}
	queue := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if low <= cur && cur <= high {
			answer = append(answer, cur)
		} else if cur > high {
			break
		}
		if dig := cur%10 + 1; dig < 10 {
			queue = append(queue, cur*10+dig)
		}
	}
	return answer
}
