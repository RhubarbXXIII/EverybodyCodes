package main

import (
	"flag"
	"maps"
	"slices"
	"strconv"
	"strings"

	"everybodycodes/utils"
)

func parseInstruction(instruction string) (direction rune, distance int) {
	direction = []rune(instruction)[0]
	distance, _ = strconv.Atoi(string([]rune(instruction)[1:]))
	return
}

func buildWalls(instructions []string) (walls map[utils.Position]bool, start, end utils.Position) {
	start = utils.NewPosition(0, 0)
	end = start

	direction := utils.UP

	walls = map[utils.Position]bool{}

	for _, instruction := range instructions {
		directionRune, distance := parseInstruction(instruction)

		switch directionRune {
		case 'L':
			direction = direction.RotateLeft()
		case 'R':
			direction = direction.RotateRight()
		}

		for range distance {
			end = end.Add(direction)

			walls[end] = true
		}
	}

	delete(walls, end)
	return
}

func solvePart1(input string) string {
	instructions := strings.Split(input, ",")

	walls, startPosition, endPosition := buildWalls(instructions)
	path := utils.FindPath(startPosition, endPosition, walls)

	return strconv.Itoa(len(path) - 1)
}

func solvePart2(input string) string {
	instructions := strings.Split(input, ",")

	walls, startPosition, endPosition := buildWalls(instructions)
	path := utils.FindPath(startPosition, endPosition, walls)

	return strconv.Itoa(len(path) - 1)
}

func solvePart3(input string) string {
	// input = `L6,L3,L6,R3,L6,L3,L3,R6,L6,R6,L6,L6,R3,L3,L3,R3,R3,L6,L6,L3`
	instructions := strings.Split(input, ",")

	var startPosition = utils.NewPosition(0, 0)
	var currentPosition = startPosition
	var endPosition utils.Position

	currentDirection := utils.UP
	nextDirection := currentDirection

	nodes := map[utils.Position]bool{currentPosition: true, currentPosition.Add(utils.UP): true, currentPosition.Add(utils.DOWN): true}

	type obstacle struct {
		start int
		end   int
	}

	rowObstacles := map[int][]obstacle{}
	columnObstacles := map[int][]obstacle{}

	addObstacle := func(currentPosition, nextPosition utils.Position, direction utils.Direction) {
		if len(rowObstacles) == 0 && len(columnObstacles) == 0 {
			switch direction {
			case utils.LEFT:
				rowObstacles[0] = append(rowObstacles[0], obstacle{start: nextPosition.Column(), end: -1})
			case utils.RIGHT:
				rowObstacles[0] = append(rowObstacles[0], obstacle{start: 1, end: nextPosition.Column()})
			}

			return
		}

		if currentDirection.IsHorizontal() {
			rowObstacles[currentPosition.Row()] = append(
				rowObstacles[currentPosition.Row()],
				obstacle{start: min(currentPosition.Column(), nextPosition.Column()), end: max(currentPosition.Column(), nextPosition.Column())},
			)
		} else if currentDirection.IsVertical() {
			columnObstacles[currentPosition.Column()] = append(
				columnObstacles[currentPosition.Column()],
				obstacle{start: min(currentPosition.Row(), nextPosition.Row()), end: max(currentPosition.Row(), nextPosition.Row())},
			)
		}
	}

	for i, currentInstruction := range instructions {
		currentDirectionRune, currentDistance := parseInstruction(currentInstruction)
		switch currentDirectionRune {
		case 'L':
			currentDirection = currentDirection.RotateLeft()
		case 'R':
			currentDirection = currentDirection.RotateRight()
		}

		nextPosition := currentPosition.Add(currentDirection.Multiply(currentDistance))

		if i == len(instructions)-1 {
			endPosition = nextPosition

			switch currentDirection {
			case utils.UP:
				columnObstacles[endPosition.Column()] = append(columnObstacles[endPosition.Column()], obstacle{start: endPosition.Row() + 1, end: currentPosition.Row()})
			case utils.RIGHT:
				rowObstacles[endPosition.Row()] = append(rowObstacles[endPosition.Row()], obstacle{start: currentPosition.Column(), end: endPosition.Column() - 1})
			case utils.DOWN:
				columnObstacles[endPosition.Column()] = append(columnObstacles[endPosition.Column()], obstacle{start: currentPosition.Row(), end: endPosition.Row() - 1})
			case utils.LEFT:
				rowObstacles[endPosition.Row()] = append(rowObstacles[endPosition.Row()], obstacle{start: endPosition.Column() + 1, end: currentPosition.Column()})
			}

			nodes[endPosition.Add(currentDirection.RotateLeft())] = true
			nodes[endPosition.Add(currentDirection.RotateRight())] = true
			nodes[endPosition] = true

			break
		}

		addObstacle(currentPosition, nextPosition, currentDirection)

		nextDirectionRune, _ := parseInstruction(instructions[i+1])
		switch nextDirectionRune {
		case 'L':
			nextDirection = currentDirection.RotateLeft()
		case 'R':
			nextDirection = currentDirection.RotateRight()
		}

		nodes[nextPosition.Add(currentDirection).Subtract(nextDirection)] = true
		nodes[nextPosition.Subtract(currentDirection).Add(nextDirection)] = true

		currentPosition = nextPosition
	}

	nodesSlice := slices.Collect(maps.Keys(nodes))
	for i := range nodesSlice {
		for j := range nodesSlice {
			nodes[utils.NewPosition(nodesSlice[i].Row(), nodesSlice[j].Column())] = true
			nodes[utils.NewPosition(nodesSlice[j].Row(), nodesSlice[i].Column())] = true
		}
	}

	nodesByRow := map[int][]utils.Position{}
	for node := range nodes {
		nodesByRow[node.Row()] = append(nodesByRow[node.Row()], node)
	}
	nodesByColumn := map[int][]utils.Position{}
	for node := range nodes {
		nodesByColumn[node.Column()] = append(nodesByColumn[node.Column()], node)
	}

	nodeRows := slices.Sorted(maps.Keys(nodesByRow))
	nodeColumns := slices.Sorted(maps.Keys(nodesByColumn))
	obstacleRows := slices.Sorted(maps.Keys(rowObstacles))
	obstacleColumns := slices.Sorted(maps.Keys(columnObstacles))

	for _, obstacles := range rowObstacles {
		slices.SortFunc(obstacles, func(left, right obstacle) int { return left.start - right.start })
	}
	for _, obstacles := range columnObstacles {
		slices.SortFunc(obstacles, func(left, right obstacle) int { return left.start - right.end })
	}

	graph := map[utils.Position]map[utils.Position]int{}

	connectNodes := func(first, second utils.Position, distance int) {
		if _, present := graph[first]; !present {
			graph[first] = map[utils.Position]int{}
		}
		if _, present := graph[second]; !present {
			graph[second] = map[utils.Position]int{}
		}

		graph[first][second] = distance
		graph[second][first] = distance
	}

	for _, row := range nodeRows {
		rowNodes := nodesByRow[row]
		slices.SortFunc(rowNodes, func(left, right utils.Position) int { return left.Column() - right.Column() })

	rowLoop:
		for i := range len(rowNodes) - 1 {
			currentNodeColumn := rowNodes[i].Column()
			nextNodeColumn := rowNodes[i+1].Column()

			for _, column := range obstacleColumns {
				if column < currentNodeColumn {
					continue
				}

				if column > nextNodeColumn {
					break
				}

				currentColumnObstacles := columnObstacles[column]
				for _, currentColumnObstacle := range currentColumnObstacles {
					if row >= currentColumnObstacle.start && row <= currentColumnObstacle.end {
						continue rowLoop
					}
				}
			}

			if currentRowObstacles, present := rowObstacles[row]; present {
				for _, currentRowObstacle := range currentRowObstacles {
					if !(nextNodeColumn < currentRowObstacle.start || currentNodeColumn > currentRowObstacle.end) {
						continue rowLoop
					}
				}
			}

			connectNodes(rowNodes[i], rowNodes[i+1], rowNodes[i].DistanceTo(rowNodes[i+1]))
		}
	}

	for _, column := range nodeColumns {
		columnNodes := nodesByColumn[column]
		slices.SortFunc(columnNodes, func(left, right utils.Position) int { return left.Row() - right.Row() })

	columnLoop:
		for i := range len(columnNodes) - 1 {
			currentNodeRow := columnNodes[i].Row()
			nextNodeRow := columnNodes[i+1].Row()

			for _, row := range obstacleRows {
				if row < currentNodeRow {
					continue
				}

				if row > nextNodeRow {
					break
				}

				currentRowObstacles := rowObstacles[row]
				for _, currentRowObstacle := range currentRowObstacles {
					if column >= currentRowObstacle.start && column <= currentRowObstacle.end {
						continue columnLoop
					}
				}
			}

			if currentColumnObstacles, present := columnObstacles[column]; present {
				for _, currentColumnObstacle := range currentColumnObstacles {
					if !(nextNodeRow < currentColumnObstacle.start || currentNodeRow > currentColumnObstacle.end) {
						continue columnLoop
					}
				}
			}

			connectNodes(columnNodes[i], columnNodes[i+1], columnNodes[i].DistanceTo(columnNodes[i+1]))
		}
	}

	path := utils.FindPathInGraph(startPosition, endPosition, graph)

	pathLength := 0
	for i := range len(path) - 1 {
		pathLength += path[i].DistanceTo(path[i+1])
	}

	return strconv.Itoa(pathLength)
}

func main() {
	submit := flag.Bool("submit", false, "Submit answers instead of dry run")
	flag.Parse()

	utils.Run(solvePart1, solvePart2, solvePart3, *submit)
}
