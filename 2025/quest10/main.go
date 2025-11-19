package main

import (
	"flag"
	"strconv"
	"strings"

	"everybodycodes/utils"
)

var knightMoves = [8]utils.Direction{
	utils.NewDirection(-2, 1),
	utils.NewDirection(-1, 2),
	utils.NewDirection(1, 2),
	utils.NewDirection(2, 1),
	utils.NewDirection(2, -1),
	utils.NewDirection(1, -2),
	utils.NewDirection(-1, -2),
	utils.NewDirection(-2, -1),
}

func parseInput(input string) utils.Grid {
	inputLines := strings.Split(input, "\n")

	board := make([][]rune, len(inputLines))
	for i, line := range inputLines {
		board[i] = []rune(line)
	}

	return utils.NewGrid(board)
}

func solvePart1(input string) string {
	board := parseInput(input)

	dragonMoveCount := 4

	dragonPositions := make([]map[utils.Position]bool, dragonMoveCount+1)
	for position := range board.Positions() {
		if board.AtPosition(position) == 'D' {
			dragonPositions[0] = map[utils.Position]bool{position: true}
		}
	}

	for i := range dragonMoveCount {
		dragonPositions[i+1] = map[utils.Position]bool{}

		for dragonPosition := range dragonPositions[i] {
			for _, move := range knightMoves {
				newPosition := dragonPosition.Add(move)
				if !board.CheckPosition(newPosition) {
					continue
				}

				dragonPositions[i+1][newPosition] = true
			}
		}
	}

	dragonSheepPositions := map[utils.Position]bool{}

	for i := range dragonMoveCount + 1 {
		for position := range dragonPositions[i] {
			if board.AtPosition(position) == 'S' {
				dragonSheepPositions[position] = true
			}
		}
	}

	return strconv.Itoa(len(dragonSheepPositions))
}

func solvePart2(input string) string {
	board := parseInput(input)

	moveCount := 20

	dragonPositions := make([]map[utils.Position]bool, moveCount+1)
	sheepPositions := make([]map[utils.Position]bool, moveCount+1)
	hideoutPositions := map[utils.Position]bool{}

	sheepPositions[0] = map[utils.Position]bool{}

	for position := range board.Positions() {
		switch board.AtPosition(position) {
		case 'D':
			dragonPositions[0] = map[utils.Position]bool{position: true}
		case 'S':
			sheepPositions[0][position] = true
		case '#':
			hideoutPositions[position] = true
		}
	}

	sheepCapturedCount := 0

	for i := range moveCount {
		dragonPositions[i+1] = map[utils.Position]bool{}
		sheepPositions[i+1] = map[utils.Position]bool{}

		for dragonPosition := range dragonPositions[i] {
			for _, move := range knightMoves {
				newDragonPosition := dragonPosition.Add(move)
				if !board.CheckPosition(newDragonPosition) {
					continue
				}

				dragonPositions[i+1][newDragonPosition] = true

				if sheepPositions[i][newDragonPosition] && !hideoutPositions[newDragonPosition] {
					sheepCapturedCount++

					delete(sheepPositions[i], newDragonPosition)
				}
			}
		}

		for sheepPosition := range sheepPositions[i] {
			newSheepPosition := sheepPosition.Add(utils.DOWN)
			if !board.CheckPosition(newSheepPosition) {
				continue
			}

			if dragonPositions[i+1][newSheepPosition] && !hideoutPositions[newSheepPosition] {
				sheepCapturedCount++
				continue
			}

			sheepPositions[i+1][newSheepPosition] = true
		}
	}

	return strconv.Itoa(sheepCapturedCount)
}

func solvePart3(input string) string {
	board := parseInput(input)

	boardHeight := board.RowCount()
	boardWidth := board.ColumnCount()

	dragonStart := utils.Position{}
	sheepStartRows := make([]int, boardWidth)
	sheepsCount := 0
	hideouts := map[utils.Position]bool{}
	escapes := make([]int, boardWidth)

	for i := range boardWidth {
		sheepStartRows[i] = -1
	}

	for position := range board.Positions() {
		switch board.AtPosition(position) {
		case 'D':
			dragonStart = position
		case 'S':
			sheepStartRows[position.Column()] = position.Row()

			sheepsCount++
		case '#':
			hideouts[position] = true
		}
	}

	for escapeColumn := range boardWidth {
		escapeRow := boardHeight
		for ; escapeRow > 0 && board.At(escapeRow-1, escapeColumn) == '#'; escapeRow-- {
		}

		escapes[escapeColumn] = escapeRow
	}

	cache := map[string]int{}
	cacheKey := func(dragonRow, dragonColumn int, sheepRows ...int) string {
		builder := strings.Builder{}
		builder.WriteByte(byte(dragonRow))
		builder.WriteByte(byte(dragonColumn))
		for _, sheepRow := range sheepRows {
			builder.WriteByte(byte(sheepRow))
		}

		return builder.String()
	}

	var countWins func(int, int, int, ...int) int
	countWins = func(dragonRow, dragonColumn, sheepsCount int, sheepRows ...int) int {
		if winCount, present := cache[cacheKey(dragonRow, dragonColumn, sheepRows...)]; present {
			return winCount
		}

		winCount := 0

		for sheepColumn, sheepRow := range sheepRows {
			if sheepRow < 0 {
				continue
			}

			newSheepRow := sheepRow + 1
			if sheepRow >= escapes[sheepColumn] {
				continue
			}

			newSheepRows := make([]int, len(sheepRows))
			copy(newSheepRows, sheepRows)

			if sheepColumn == dragonColumn && newSheepRow == dragonRow && !hideouts[utils.NewPosition(newSheepRow, sheepColumn)] {
				if sheepsCount > 1 {
					continue
				}
			} else {
				newSheepRows[sheepColumn] = newSheepRow
			}

			for _, knightMove := range knightMoves {
				newDragon := utils.NewPosition(dragonRow, dragonColumn).Add(knightMove)
				if !board.CheckPosition(newDragon) {
					continue
				}

				newSheepRowsAfterDragon := make([]int, len(newSheepRows))
				copy(newSheepRowsAfterDragon, newSheepRows)

				newSheepsCountAfterDragon := sheepsCount
				if newSheepRowsAfterDragon[newDragon.Column()] == newDragon.Row() && !hideouts[newDragon] {
					newSheepsCountAfterDragon = sheepsCount - 1

					if newSheepsCountAfterDragon == 0 {
						winCount++
						continue
					}

					newSheepRowsAfterDragon[newDragon.Column()] = -1
				}

				winCount += countWins(newDragon.Row(), newDragon.Column(), newSheepsCountAfterDragon, newSheepRowsAfterDragon...)
			}
		}

		cache[cacheKey(dragonRow, dragonColumn, sheepRows...)] = winCount
		return winCount
	}

	return strconv.Itoa(countWins(dragonStart.Row(), dragonStart.Column(), sheepsCount, sheepStartRows...))
}

func main() {
	submit := flag.Bool("submit", false, "Submit answers instead of dry run")
	flag.Parse()

	utils.Run(solvePart1, solvePart2, solvePart3, *submit)
}
