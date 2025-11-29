package utils

import (
	"slices"
)

type Node struct {
	Position Position
	Value    any

	G int
	H int

	Previous *Node
}

func (n Node) F() int {
	return n.G + n.H
}

func FindPath(start, end Position, obstacles map[Position]bool) []Position {
	if start == end {
		return []Position{start}
	}

	startNode := Node{
		Position: start,
		Value:    nil,
		G:        0,
		H:        start.DistanceTo(end),
		Previous: nil,
	}

	queue := NewHeap[Node]()
	queue.Push(&startNode, startNode.F())

	visited := map[Position]int{}

	var currentNode *Node
	for !queue.Empty() {
		currentNode = queue.Pop()

		if distance, present := visited[currentNode.Position]; present && distance <= currentNode.G {
			continue
		}

		visited[currentNode.Position] = currentNode.G

		if currentNode.Position == end {
			break
		}

		for _, direction := range DIRECTIONS {
			nextPosition := currentNode.Position.Add(direction)
			if obstacles[nextPosition] {
				continue
			}

			nextNode := Node{
				Position: nextPosition,
				Value:    nil,
				G:        currentNode.G + 1,
				H:        nextPosition.DistanceTo(end),
				Previous: currentNode,
			}

			queue.Push(&nextNode, nextNode.F())
		}
	}

	path := []Position{currentNode.Position}
	for currentNode.Previous != nil {
		currentNode = currentNode.Previous

		path = append(path, currentNode.Position)
	}

	slices.Reverse(path)
	return path
}

func FindPathInGraph(start, end Position, graph map[Position]map[Position]int) []Position {
	if start == end {
		return []Position{start}
	}

	startNode := Node{
		Position: start,
		Value:    nil,
		G:        0,
		H:        start.DistanceTo(end),
		Previous: nil,
	}

	queue := NewHeap[Node]()
	queue.Push(&startNode, startNode.F())

	visited := map[Position]int{}

	var currentNode *Node
	for !queue.Empty() {
		currentNode = queue.Pop()

		if distance, present := visited[currentNode.Position]; present && distance <= currentNode.G {
			continue
		}

		visited[currentNode.Position] = currentNode.G

		if currentNode.Position == end {
			break
		}

		for nextPosition, distance := range graph[currentNode.Position] {
			nextNode := Node{
				Position: nextPosition,
				Value:    nil,
				G:        currentNode.G + distance,
				H:        nextPosition.DistanceTo(end),
				Previous: currentNode,
			}

			queue.Push(&nextNode, nextNode.F())
		}
	}

	path := []Position{currentNode.Position}
	for currentNode.Previous != nil {
		currentNode = currentNode.Previous

		path = append(path, currentNode.Position)
	}

	slices.Reverse(path)
	return path
}
