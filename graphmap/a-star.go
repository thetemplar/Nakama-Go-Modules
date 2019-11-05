package graphmap

import (
	"container/heap"
)

// Returns the shortest path from the Triangle with key startKey to the Triangle with key endKey as a string slice, and if such a path exists at all, using a function to calculate an estimated distance from a Triangle to the endTriangle. The heuristic function is passed the keys of a Triangle and the end Triangle. This function uses the A* search algorithm.
// If startKey or endKey (or both) are invalid, path will be empty and exists will be false.
func (g *Graph) ShortestPathWithHeuristic(startKey, endKey int32, heuristic func(key, endKey int32) int) (path []int32, exists bool) {
	// start and end Triangle
	start := g.get(startKey)
	end := g.get(endKey)

	// check startKey and endKey for validity
	if start == nil || end == nil {
		return
	}

	// priorityQueue for Triangles that have not yet been visited (open Triangles)
	openQueue := &priorityQueue{}

	// list containing Triangles that have not yet been visited (open Triangles)
	openList := map[*Triangle]*Item{}

	// list containing Triangles that have been visited already (closed Triangles)
	closedList := map[*Triangle]*Item{}

	// add start Triangle to list of open Triangles
	item := &Item{start, nil, 0, 0, 0}
	openList[start] = item

	heap.Push(openQueue, item)

	for openQueue.Len() > 0 {
		current := heap.Pop(openQueue).(*Item).n

		// current Triangle was now visited; add to closed list
		closedList[current] = openList[current]
		delete(openList, current)

		// end Triangle found?
		if current == end {
			// path exists
			exists = true

			// build path
			for current != nil {
				path = append(path, current.Id)
				current = closedList[current].prev
			}

			return
		}

		// saved here for easy usage in following loop
		distance := closedList[current].distanceFromStart

		for neighbor, weight := range current.GetNeighbors() {
			if _, ok := closedList[neighbor]; ok {
				continue
			}

			distanceToNeighbor := distance + weight

			// skip neighbors that already have a better path leading to them
			if md, ok := openList[neighbor]; ok {
				if md.distanceFromStart < distanceToNeighbor {
					continue
				} else {
					heap.Remove(openQueue, md.index)
				}
			}

			item := &Item{
				neighbor,
				current,
				distanceToNeighbor,
				distanceToNeighbor + heuristic(neighbor.Id, endKey), // estimate (= priority)
				0,
			}

			// add neighbor Triangle to list of open Triangles
			openList[neighbor] = item

			// push into priority queue
			heap.Push(openQueue, item)
		}
	}

	return
}
