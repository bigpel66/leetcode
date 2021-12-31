// 973. K Closest Points to Origin
//
// https://leetcode.com/problems/k-closest-points-to-origin/

// container/heap package is used for implementing priorityQueue by Heap Interface.
import (
	"container/heap"
)

// priorityQueue is used for queueing several points.
// priorityQueue is based on Heap Interface.
type priorityQueue [][]int

// Len function for priorityQueue.
func (pq priorityQueue) Len() int {
	return len(pq)
}

// Descending order should be maintained on the priorityQueue.
func (pq priorityQueue) Less(i, j int) bool {
	return dist(pq[i]) > dist(pq[j])
}

// Swap function for priorityQueue.
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push function for Heap Interface.
func (pq *priorityQueue) Push(element interface{}) {
	*pq = append(*pq, element.([]int))
}

// Pop function for Heap Interface.
func (pq *priorityQueue) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return x
}

// dist function returns the relative distance between point x and point y.
func dist(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}

// kClosest function finds k closest points to the origin (0, 0).
// A simple solution is just sorting from the distance to the origin.
// Also, picking k closest points (k minimum distance points) could be solved with Heap Tree.
// Push a point unconditionally, pop a point if it is not suitable.
// Popping procedure extract the maximum distance due to priorityQueue is built with Max Heap
func kClosest(points [][]int, k int) [][]int {
	pq := &priorityQueue{}
	res := [][]int{}
	heap.Init(pq)
	for _, point := range points {
		heap.Push(pq, point)
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}
	for pq.Len() > 0 {
		res = append(res, heap.Pop(pq).([]int))
	}
	return res
}
