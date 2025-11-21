package main

import (
	"flag"
	"slices"
	"strconv"
	"strings"

	"everybodycodes/utils"
)

func parseInput(input string) utils.Grid {
	inputLines := strings.Split(input, "\n")

	board := make([][]rune, len(inputLines))
	for i, line := range inputLines {
		board[i] = []rune(line)
	}

	return utils.NewGrid(board)
}

func advanceTiles(currentTiles *utils.Grid) (nextTiles *utils.Grid, activeTileCount int) {
	nextTiles = currentTiles.Clone()
	activeTileCount = 0

	for tilePosition, tile := range currentTiles.Cells() {
		activeNeighborTileCount := 0

		for i := range utils.DIRECTIONS {
			neighborTilePosition := tilePosition.Add(utils.DIRECTIONS[i]).Add(utils.DIRECTIONS[(i+1)%len(utils.DIRECTIONS)])
			if !currentTiles.CheckPosition(neighborTilePosition) {
				continue
			}

			if currentTiles.AtPosition(neighborTilePosition) == '#' {
				activeNeighborTileCount++
			}
		}

		if tile == '#' {
			if activeNeighborTileCount%2 == 1 {
				nextTiles.SetPosition(tilePosition, '#')

				activeTileCount++
			} else {
				nextTiles.SetPosition(tilePosition, '.')
			}
		} else {
			if activeNeighborTileCount%2 == 0 {
				nextTiles.SetPosition(tilePosition, '#')

				activeTileCount++
			} else {
				nextTiles.SetPosition(tilePosition, '.')
			}
		}
	}

	return nextTiles, activeTileCount
}

func solvePart1(input string) string {
	tiles := parseInput(input)

	roundCount := 10

	totalActiveTileCount := 0

	currentTiles := tiles
	for range roundCount {
		nextTiles, activeTileCount := advanceTiles(&currentTiles)

		totalActiveTileCount += activeTileCount

		currentTiles = *nextTiles
	}

	return strconv.Itoa(totalActiveTileCount)
}

func solvePart2(input string) string {
	tiles := parseInput(input)

	roundCount := 2025

	totalActiveTileCount := 0

	currentTiles := tiles
	for range roundCount {
		nextTiles, activeTileCount := advanceTiles(&currentTiles)

		totalActiveTileCount += activeTileCount

		currentTiles = *nextTiles
	}

	return strconv.Itoa(totalActiveTileCount)
}

func solvePart3(input string) string {
	pattern := parseInput(input)
	patternWidth := pattern.RowCount()

	roundCount := 1000000000

	boardWidth := 34
	boardCells := make([][]rune, boardWidth)
	for i := range boardWidth {
		boardCells[i] = []rune(strings.Repeat(".", boardWidth))
	}

	patternStartIndex := (boardWidth - patternWidth) / 2

	tiles := utils.NewGrid(boardCells)
	tilesCache := map[string]int{tiles.String(): 0}

	activeTileCounts := []int{0}

	firstCycleStartRound := -1
	secondCycleStartRound := -1
	cycleRoundCount := -1

	roundIndex := 1
	for currentTiles := &tiles; roundIndex <= roundCount; roundIndex++ {
		nextTiles, activeTileCount := advanceTiles(currentTiles)
		nextTilesString := nextTiles.String()

		if cycleStartRound, present := tilesCache[nextTilesString]; present {
			firstCycleStartRound = cycleStartRound
			secondCycleStartRound = roundIndex

			cycleRoundCount = secondCycleStartRound - firstCycleStartRound
			break
		}

		matchesPattern := true
		for i, row := range nextTiles.RowsFrom(patternStartIndex, patternStartIndex+patternWidth) {
			if !slices.Equal(row[patternStartIndex:patternStartIndex+patternWidth], pattern.RowAt(i)) {
				matchesPattern = false
				break
			}
		}

		if matchesPattern {
			activeTileCounts = append(activeTileCounts, activeTileCount)
		} else {
			activeTileCounts = append(activeTileCounts, 0)
		}

		tilesCache[nextTilesString] = roundIndex

		currentTiles = nextTiles
	}

	totalActiveTileCount := 0
	cycleActiveTileCount := 0

	for i := 0; i < secondCycleStartRound; i++ {
		if i < firstCycleStartRound {
			totalActiveTileCount += activeTileCounts[i]
		} else {
			cycleActiveTileCount += activeTileCounts[i]
		}
	}

	totalActiveTileCount += (roundCount - firstCycleStartRound) / cycleRoundCount * cycleActiveTileCount

	for i := firstCycleStartRound; i <= (roundCount-firstCycleStartRound)%cycleRoundCount; i++ {
		totalActiveTileCount += activeTileCounts[i]
	}

	return strconv.Itoa(totalActiveTileCount)
}

func main() {
	submit := flag.Bool("submit", false, "Submit answers instead of dry run")
	flag.Parse()

	utils.Run(solvePart1, solvePart2, solvePart3, *submit)
}
