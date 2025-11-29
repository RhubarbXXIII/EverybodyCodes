package main

import (
	"flag"
	"strconv"
	"strings"

	"everybodycodes/utils"
)

func parseInput(input string) utils.Grid {
	inputLines := strings.Split(input, "\n")

	cells := make([][]rune, len(inputLines))
	for i, inputLine := range inputLines {
		cells[i] = []rune(inputLine)
	}

	return utils.NewGrid(cells)
}

func solvePart1(input string) string {
	cells := parseInput(input)

	trampolinePairCount := 0

	for rowIndex, row := range cells.Rows() {
		for columnIndex := range row {
			if row[columnIndex] == '.' {
				continue
			}

			if rowIndex < cells.RowCount()-1 && rowIndex%2 != columnIndex%2 {
				if cells.At(rowIndex, columnIndex) == 'T' && cells.At(rowIndex+1, columnIndex) == 'T' {
					trampolinePairCount++
				}
			}

			if columnIndex < cells.ColumnCount()-1 {
				if row[columnIndex] == 'T' && row[columnIndex+1] == 'T' {
					trampolinePairCount++
				}
			}
		}
	}

	return strconv.Itoa(trampolinePairCount)
}

func solvePart2(input string) string {
	cells := parseInput(input)

	var startPosition, endPosition utils.Position
	for position, cell := range cells.Cells() {
		switch cell {
		case 'S':
			startPosition = position
		case 'E':
			endPosition = position
		}
	}
	startNode := utils.Node{
		Position: startPosition,
		Value:    nil,
		G:        0,
		H:        startPosition.DistanceTo(endPosition),
		Previous: nil,
	}

	queue := utils.NewHeap[utils.Node]()
	queue.Push(&startNode, startNode.F())

	visited := map[utils.Position]int{}

	var currentNode *utils.Node
	for !queue.Empty() {
		currentNode = queue.Pop()

		if distance, present := visited[currentNode.Position]; present && distance <= currentNode.G {
			continue
		}

		visited[currentNode.Position] = currentNode.G

		if currentNode.Position == endPosition {
			break
		}

		var verticalDirection utils.Direction
		if currentNode.Position.Row()%2 == currentNode.Position.Column()%2 {
			verticalDirection = utils.UP
		} else {
			verticalDirection = utils.DOWN
		}

		for _, direction := range []utils.Direction{utils.LEFT, utils.RIGHT, verticalDirection} {
			nextPosition := currentNode.Position.Add(direction)
			if !cells.CheckPosition(nextPosition) {
				continue
			}

			nextCell := cells.AtPosition(nextPosition)
			if nextCell == '.' || nextCell == '#' {
				continue
			}

			nextNode := utils.Node{
				Position: nextPosition,
				Value:    nil,
				G:        currentNode.G + 1,
				H:        nextPosition.DistanceTo(endPosition),
				Previous: currentNode,
			}

			queue.Push(&nextNode, nextNode.F())
		}
	}

	return strconv.Itoa(currentNode.G)
}

func solvePart3(input string) string {
	grids := make([]utils.Grid, 3)
	grids[0] = parseInput(input)

	var startPosition, endPosition utils.Position
	for position, cell := range grids[0].Cells() {
		switch cell {
		case 'S':
			startPosition = position
		case 'E':
			endPosition = position
		}
	}

	for i := range 2 {
		grids[i+1] = *grids[i].Clone()

		for newRowIndex := range grids[i].Rows() {
			oldColumnIndex := grids[i].RowCount() - 1 + newRowIndex
			newColumnIndex := newRowIndex

			for oldRowIndex := grids[i].RowCount() - 1 - newRowIndex; oldRowIndex >= 0; oldRowIndex-- {
				grids[i+1].Set(newRowIndex, newColumnIndex, grids[i].At(oldRowIndex, oldColumnIndex))

				if oldRowIndex > 0 {
					grids[i+1].Set(newRowIndex, newColumnIndex+1, grids[i].At(oldRowIndex-1, oldColumnIndex))
				}

				oldColumnIndex -= 1
				newColumnIndex += 2
			}
		}
	}

	startNode := utils.Node{
		Position: startPosition,
		Value:    nil,
		G:        0,
		H:        startPosition.DistanceTo(endPosition),
		Previous: nil,
	}

	queue := utils.NewHeap[utils.Node]()
	queue.Push(&startNode, startNode.G)

	var visited []map[utils.Position]int
	for range 3 {
		visited = append(visited, map[utils.Position]int{})
	}

	var currentNode *utils.Node
	for !queue.Empty() {
		currentNode = queue.Pop()

		gridIndex := currentNode.G % 3

		if distance, present := visited[gridIndex][currentNode.Position]; present && distance <= currentNode.G {
			continue
		}

		visited[gridIndex][currentNode.Position] = currentNode.G

		if grids[gridIndex].AtPosition(currentNode.Position) == 'E' {
			break
		}

		var verticalDirection utils.Direction
		if currentNode.Position.Row()%2 == currentNode.Position.Column()%2 {
			verticalDirection = utils.UP
		} else {
			verticalDirection = utils.DOWN
		}

		for _, direction := range []utils.Direction{utils.LEFT, utils.RIGHT, verticalDirection, utils.NewDirection(0, 0)} {
			nextGridIndex := (gridIndex + 1) % 3

			nextPosition := currentNode.Position.Add(direction)
			if !grids[nextGridIndex].CheckPosition(nextPosition) {
				continue
			}

			nextCell := grids[nextGridIndex].AtPosition(nextPosition)
			if nextCell == '.' || nextCell == '#' {
				continue
			}

			nextNode := utils.Node{
				Position: nextPosition,
				Value:    nil,
				G:        currentNode.G + 1,
				H:        nextPosition.DistanceTo(endPosition),
				Previous: currentNode,
			}

			queue.Push(&nextNode, nextNode.G)
		}
	}

	return strconv.Itoa(currentNode.G)
}

func main() {
	submit := flag.Bool("submit", false, "Submit answers instead of dry run")
	flag.Parse()

	utils.Run(solvePart1, solvePart2, solvePart3, *submit)
}
