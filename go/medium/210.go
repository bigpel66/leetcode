// 210. Course Schedule II
//
// https://leetcode.com/problems/course-schedule-ii/

// fundOrder function finds the order of the course taking.
// Some of courses have some prerequisites.
// If the chaining count is 0, then the course should be append to the result slice.
// Others could be ordered and appended to the slice by using BFS + Topological Sort.
// To use Topological Sort, count slice and map have been used.
// Remember that prerequisites slice is not bidirectional graph. [a_i, b_i] means b_i is the prerequisites of a_i, so b_i -> a_i will be built. Also, the count of specific node which has some prerequisites will be incremented.
// If there's a cycle such as [[1, 0], [1, 2], [0, 1]], then output should be []. Thus, counting the number of course is necessary.
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make(map[int][]int, 0)
	count := make([]int, numCourses)
	for _, course := range prerequisites {
		graph[course[1]] = append(graph[course[1]], course[0])
		count[course[0]]++
	}
	entry := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if count[i] == 0 {
			entry = append(entry, i)
		}
	}
	result := make([]int, 0)
	for len(entry) > 0 {
		former := entry[0]
		result = append(result, former)
		entry = entry[1:]
		numCourses--
		for _, course := range graph[former] {
			count[course]--
			if count[course] == 0 {
				entry = append(entry, course)
			}
		}
	}
	if numCourses != 0 {
		return []int{}
	}
	return result
}
