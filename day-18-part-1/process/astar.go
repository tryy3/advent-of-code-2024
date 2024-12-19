package process

import (
	"container/heap"
	"math"
)

var directions = []Position{
	{x: 0, y: 1}, {x: 0, y: -1}, {x: 1, y: 0}, {x: -1, y: 0},
}

func AStar(start, goal Position, mapGrid [][]bool, blockSchedule map[int]map[Position]bool, currentTime int) []Position {
	rows, cols := len(mapGrid), len(mapGrid[0])
	openSet := &PriorityQueue{}
	heap.Init(openSet)

	// Initialize the start node.
	startNode := &Node{
		pos:      start,
		gCost:    0,
		hCost:    manhattanDistance(start, goal),
		timeStep: currentTime,
	}
	heap.Push(openSet, startNode)

	// Visited set with time-based tracking.
	visited := make(map[Position]int)

	for openSet.Len() > 0 {
		// Get the node with the lowest fCost.
		current := heap.Pop(openSet).(*Node)

		// Check if we've reached the goal.
		if current.pos == goal {
			return reconstructPath(current)
		}

		// Mark as visited.
		visited[current.pos] = current.timeStep

		// Explore neighbors.
		for _, dir := range directions {
			neighbor := Position{
				x: current.pos.x + dir.x,
				y: current.pos.y + dir.y,
			}

			// Check if the neighbor is within bounds.
			if neighbor.x < 0 || neighbor.y < 0 || neighbor.x >= rows || neighbor.y >= cols {
				continue
			}

			// Check if the neighbor is blocked at the next time step.
			nextTime := current.timeStep + 1
			if mapGrid[neighbor.y][neighbor.x] || (blockSchedule[nextTime] != nil && blockSchedule[nextTime][neighbor]) {
				continue
			}

			// Check if we've already visited this node at the same or earlier time.
			if visitedTime, ok := visited[neighbor]; ok && visitedTime <= nextTime {
				continue
			}

			// Calculate costs.
			gCost := current.gCost + 1
			hCost := manhattanDistance(neighbor, goal)

			// Push the neighbor to the open set.
			heap.Push(openSet, &Node{
				pos:      neighbor,
				gCost:    gCost,
				hCost:    hCost,
				timeStep: nextTime,
				parent:   current,
			})
		}
	}

	// Return empty path if no solution.
	return nil
}

// manhattanDistance calculates the Manhattan distance between two positions.
func manhattanDistance(a, b Position) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

// reconstructPath reconstructs the path from the goal node.
func reconstructPath(goalNode *Node) []Position {
	path := []Position{}
	current := goalNode
	for current != nil {
		path = append([]Position{current.pos}, path...)
		current = current.parent
	}
	return path
}
