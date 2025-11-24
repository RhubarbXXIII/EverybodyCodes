package utils

import (
	"slices"
)

type node struct {
	position Position
	value    any

	g int
	h int

	previous *node
}

func (n node) f() int {
	return n.g + n.h
}

func FindPath(start, end Position, obstacles map[Position]bool) []Position {
	if start == end {
		return []Position{start}
	}

	startNode := node{
		position: start,
		value:    nil,
		g:        0,
		h:        start.DistanceTo(end),
		previous: nil,
	}

	queue := NewHeap[node]()
	queue.Push(&startNode, startNode.f())

	visited := map[Position]int{}

	var currentNode *node
	for !queue.Empty() {
		currentNode = queue.Pop()

		if distance, present := visited[currentNode.position]; present && distance <= currentNode.g {
			continue
		}

		visited[currentNode.position] = currentNode.g

		if currentNode.position == end {
			break
		}

		for _, direction := range DIRECTIONS {
			nextPosition := currentNode.position.Add(direction)
			if obstacles[nextPosition] {
				continue
			}

			nextNode := node{
				position: nextPosition,
				value:    nil,
				g:        currentNode.g + 1,
				h:        nextPosition.DistanceTo(end),
				previous: currentNode,
			}

			queue.Push(&nextNode, nextNode.f())
		}
	}

	path := []Position{currentNode.position}
	for currentNode.previous != nil {
		currentNode = currentNode.previous

		path = append(path, currentNode.position)
	}

	slices.Reverse(path)
	return path
}

func FindPathInGraph(start, end Position, graph map[Position]map[Position]int) []Position {
	if start == end {
		return []Position{start}
	}

	startNode := node{
		position: start,
		value:    nil,
		g:        0,
		h:        start.DistanceTo(end),
		previous: nil,
	}

	queue := NewHeap[node]()
	queue.Push(&startNode, startNode.f())

	visited := map[Position]int{}

	var currentNode *node
	for !queue.Empty() {
		currentNode = queue.Pop()

		if distance, present := visited[currentNode.position]; present && distance <= currentNode.g {
			continue
		}

		visited[currentNode.position] = currentNode.g

		if currentNode.position == end {
			break
		}

		for nextPosition, distance := range graph[currentNode.position] {
			nextNode := node{
				position: nextPosition,
				value:    nil,
				g:        currentNode.g + distance,
				h:        nextPosition.DistanceTo(end),
				previous: currentNode,
			}

			queue.Push(&nextNode, nextNode.f())
		}
	}

	path := []Position{currentNode.position}
	for currentNode.previous != nil {
		currentNode = currentNode.previous

		path = append(path, currentNode.position)
	}

	slices.Reverse(path)
	return path
}
