package main

import (
	"flag"
	"math"
	"strconv"
	"strings"

	"everybodycodes/utils"
)

func parseInput(input string) (field utils.Grid, volcanoPosition, startPosition utils.Position) {
	inputLines := strings.Split(input, "\n")

	fieldCells := make([][]rune, len(inputLines))
	for i, inputLine := range inputLines {
		fieldCells[i] = []rune(inputLine)
	}

	field = utils.NewGrid(fieldCells)

	for position, cell := range field.Cells() {
		switch cell {
		case '@':
			volcanoPosition = position
		case 'S':
			startPosition = position
		}
	}

	return
}

func cellInRadius(cellPosition, volcanoPosition utils.Position, volcanoRadius int) bool {
	return math.Pow(float64(volcanoPosition.Column()-cellPosition.Column()), 2)+math.Pow(float64(volcanoPosition.Row()-cellPosition.Row()), 2) <= float64(volcanoRadius*volcanoRadius)
}

func solvePart1(input string) string {
	field, volcanoPosition, _ := parseInput(input)

	volcanoRadius := 10

	cellValueSum := 0
	for cellPosition, cell := range field.Cells() {
		if cell == '@' {
			continue
		}

		if cellInRadius(cellPosition, volcanoPosition, volcanoRadius) {
			cellValue, _ := strconv.Atoi(string(cell))
			cellValueSum += cellValue
		}
	}

	return strconv.Itoa(cellValueSum)
}

func solvePart2(input string) string {
	field, volcanoPosition, _ := parseInput(input)

	volcanoRadius := volcanoPosition.Row()

	destroyedCells := map[utils.Position]bool{}
	destroyedCellValueSums := make([]int, volcanoRadius)

	for currentRadius := 1; currentRadius <= volcanoRadius; currentRadius++ {
		for cellPosition, cell := range field.Cells() {
			if cell == '@' {
				continue
			}

			if destroyedCells[cellPosition] {
				continue
			}

			if !cellInRadius(cellPosition, volcanoPosition, currentRadius) {
				continue
			}

			cellValue, _ := strconv.Atoi(string(cell))

			destroyedCells[cellPosition] = true
			destroyedCellValueSums[currentRadius-1] += cellValue
		}
	}

	maximumDestroyedCellRadius := -1
	maximumDestroyedCellCount := 0
	for i, cellValueCount := range destroyedCellValueSums {
		if cellValueCount > maximumDestroyedCellCount {
			maximumDestroyedCellRadius = i + 1
			maximumDestroyedCellCount = cellValueCount
		}
	}

	return strconv.Itoa(maximumDestroyedCellRadius * maximumDestroyedCellCount)
}

func solvePart3(input string) string {
	field, volcanoPosition, startPosition := parseInput(input)

	quadrant := func(position utils.Position) int {
		if position.Row() < volcanoPosition.Row() && position.Column() <= volcanoPosition.Column() {
			return 1
		} else if position.Row() >= volcanoPosition.Row() && position.Column() < volcanoPosition.Column() {
			return 2
		} else if position.Row() > volcanoPosition.Row() && position.Column() >= volcanoPosition.Column() {
			return 3
		} else if position.Row() <= volcanoPosition.Row() && position.Column() > volcanoPosition.Column() {
			return 4
		} else {
			return 2
		}
	}

radiusLoop:
	for currentRadius := 1; currentRadius <= field.RowCount()/2; currentRadius++ {
		currentPosition := startPosition

		currentTime := 0
		maximumTime := 30 * currentRadius

		currentNode := &utils.Node{
			Position: currentPosition,
			Value:    nil,
			G:        0,
			H:        0,
			Previous: nil,
		}

		queue := utils.NewHeap[utils.Node]()
		queue.Push(currentNode, currentNode.F())

		visited := map[int]map[utils.Position]int{}
		for i := range 4 {
			visited[i+1] = map[utils.Position]int{}
		}

		for !queue.Empty() {
			currentNode = queue.Pop()
			currentQuadrant := quadrant(currentNode.Position)

			if distance, present := visited[currentQuadrant][currentNode.Position]; present && distance <= currentNode.G {
				continue
			}

			if currentQuadrant == 4 && len(visited[4]) == 0 {
				clear(visited[1])
			}

			visited[currentQuadrant][currentNode.Position] = currentNode.G

			if currentNode.Position == startPosition && len(visited[4]) > 0 {
				break
			}

			for _, direction := range utils.DIRECTIONS {
				nextPosition := currentNode.Position.Add(direction)
				nextQuadrant := quadrant(nextPosition)

				if !field.CheckPosition(nextPosition) {
					continue
				}

				if field.AtPosition(nextPosition) == '@' {
					continue
				}

				if cellInRadius(nextPosition, volcanoPosition, currentRadius-1) {
					continue
				}

				if nextQuadrant == 4 && len(visited[3]) == 0 {
					continue
				}

				if nextPosition == startPosition && len(visited[4]) == 0 {
					continue
				}

				time, _ := strconv.Atoi(string(field.AtPosition(nextPosition)))

				nextNode := utils.Node{
					Position: nextPosition,
					Value:    nil,
					G:        currentNode.G + time,
					H:        nextPosition.DistanceTo(startPosition),
					Previous: currentNode,
				}

				queue.Push(&nextNode, nextNode.F())
			}
		}

		currentTime += currentNode.G
		if currentTime >= maximumTime {
			continue radiusLoop
		}

		return strconv.Itoa((currentRadius - 1) * currentTime)
	}

	panic("No timely path found!")
}

func main() {
	submit := flag.Bool("submit", false, "Submit answers instead of dry run")
	flag.Parse()

	utils.Run(solvePart1, solvePart2, solvePart3, *submit)
}
