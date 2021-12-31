// 310. Minimum Height Trees
//
// https://leetcode.com/problems/minimum-height-trees/

// findMinHeightTrees function finds the possible root nodes which has the minimum height.
// Finding the longest path of the tree by DFS can be possible, but topological sort is more suitable.
// To find the minimum height, remove the possible leaf nodes by iterative loop.
// Removing leaf nodes can be done by looking up the edges.
// Also the iterating should be stopped when the existing nodes are less than or equal to 2.
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	graph := make(map[int][]int, n)
	count := make([]int, n)
	for _, e := range edges {
		if _, ok := graph[e[0]]; !ok {
			graph[e[0]] = make([]int, 0)
		}
		if _, ok := graph[e[1]]; !ok {
			graph[e[1]] = make([]int, 0)
		}
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
		count[e[0]]++
		count[e[1]]++
	}
	leaves := make([]int, 0)
	for i, c := range count {
		if c == 1 {
			leaves = append(leaves, i)
		}
	}
	for n > 2 {
		size := len(leaves)
		n -= size
		for i := 0; i < size; i++ {
			count[leaves[i]]--
			for _, c := range graph[leaves[i]] {
				count[c]--
				if count[c] == 1 {
					leaves = append(leaves, c)
				}
			}
		}
		leaves = leaves[size:]
	}
	return leaves
}
