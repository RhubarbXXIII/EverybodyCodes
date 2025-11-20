package main

import (
	"container/list"
	"flag"
	"maps"
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

func solvePart1(input string) string {
	barrels := parseInput(input)

	barrelPositionsSeen := map[utils.Position]bool{}
	barrelPositionsRemaining := list.New()

	for barrelPositionsRemaining.PushBack(utils.NewPosition(0, 0)); barrelPositionsRemaining.Len() > 0; {
		barrelPosition := barrelPositionsRemaining.Front().Value.(utils.Position)
		barrelSize, _ := strconv.Atoi(string(barrels.AtPosition(barrelPosition)))

		barrelPositionsRemaining.Remove(barrelPositionsRemaining.Front())

		if barrelPositionsSeen[barrelPosition] {
			continue
		}

		barrelPositionsSeen[barrelPosition] = true

		for _, direction := range utils.DIRECTIONS {
			newBarrelPosition := barrelPosition.Add(direction)
			if !barrels.CheckPosition(newBarrelPosition) {
				continue
			}

			if barrelPositionsSeen[newBarrelPosition] {
				continue
			}

			if newBarrelSize, _ := strconv.Atoi(string(barrels.AtPosition(newBarrelPosition))); newBarrelSize > barrelSize {
				continue
			}

			barrelPositionsRemaining.PushBack(newBarrelPosition)
		}
	}

	return strconv.Itoa(len(barrelPositionsSeen))
}

func solvePart2(input string) string {
	barrels := parseInput(input)

	barrelPositionsSeen := map[utils.Position]bool{}
	barrelPositionsRemaining := list.New()

	barrelPositionsRemaining.PushBack(utils.NewPosition(0, 0))
	barrelPositionsRemaining.PushBack(utils.NewPosition(barrels.RowCount()-1, barrels.ColumnCount()-1))

	for barrelPositionsRemaining.Len() > 0 {
		barrelPosition := barrelPositionsRemaining.Front().Value.(utils.Position)
		barrelSize, _ := strconv.Atoi(string(barrels.AtPosition(barrelPosition)))

		barrelPositionsRemaining.Remove(barrelPositionsRemaining.Front())

		if barrelPositionsSeen[barrelPosition] {
			continue
		}

		barrelPositionsSeen[barrelPosition] = true

		for _, direction := range utils.DIRECTIONS {
			newBarrelPosition := barrelPosition.Add(direction)
			if !barrels.CheckPosition(newBarrelPosition) {
				continue
			}

			if barrelPositionsSeen[newBarrelPosition] {
				continue
			}

			if newBarrelSize, _ := strconv.Atoi(string(barrels.AtPosition(newBarrelPosition))); newBarrelSize > barrelSize {
				continue
			}

			barrelPositionsRemaining.PushBack(newBarrelPosition)
		}
	}

	return strconv.Itoa(len(barrelPositionsSeen))
}

func solvePart3(input string) string {
	// 	input = `41951111131882511179
	// 32112222211518122215
	// 31223333322115122219
	// 31234444432147511128
	// 91223333322176121892
	// 61112222211166431583
	// 14661111166111111746
	// 11111119142122222177
	// 41222118881233333219
	// 71222127839122222196
	// 56111126279711111517`
	barrels := parseInput(input)

	barrelsRemaining := map[utils.Position]int{}
	for barrelPosition, barrelSizeRune := range barrels.Cells() {
		barrelSize, _ := strconv.Atoi(string(barrelSizeRune))

		barrelsRemaining[barrelPosition] = barrelSize
	}

	barrelCount := 0

	for range 3 {
		cache := map[utils.Position][]utils.Position{}

		barrelPositionsBySize := slices.SortedFunc(maps.Keys(barrelsRemaining), func(left, right utils.Position) int {
			return barrelsRemaining[left] - barrelsRemaining[right]
		})

		for _, startBarrelPosition := range barrelPositionsBySize {
			barrelPositionsRemaining := list.New()
			barrelPositionsSeen := map[utils.Position]bool{}

			for barrelPositionsRemaining.PushBack(startBarrelPosition); barrelPositionsRemaining.Len() > 0; {
				barrelPosition := barrelPositionsRemaining.Front().Value.(utils.Position)
				barrelSize, _ := strconv.Atoi(string(barrels.AtPosition(barrelPosition)))

				barrelPositionsRemaining.Remove(barrelPositionsRemaining.Front())

				if barrelPositionsSeen[barrelPosition] {
					continue
				}

				barrelPositionsSeen[barrelPosition] = true

				for _, direction := range utils.DIRECTIONS {
					newBarrelPosition := barrelPosition.Add(direction)
					if _, present := barrelsRemaining[newBarrelPosition]; !present {
						continue
					}

					if barrelPositionsSeen[newBarrelPosition] {
						continue
					}

					if newBarrelSize, _ := strconv.Atoi(string(barrels.AtPosition(newBarrelPosition))); newBarrelSize > barrelSize {
						continue
					}

					if cachedBarrelPositions, present := cache[newBarrelPosition]; present {
						for _, cachedBarrelPosition := range cachedBarrelPositions {
							barrelPositionsSeen[cachedBarrelPosition] = true
						}

						continue
					}

					barrelPositionsRemaining.PushBack(newBarrelPosition)
				}
			}

			cache[startBarrelPosition] = slices.Collect(maps.Keys(barrelPositionsSeen))
		}

		maxChainedBarrelPosition := utils.NewPosition(0, 0)
		for barrelPosition, chainedBarrelPositions := range cache {
			if len(chainedBarrelPositions) > len(cache[maxChainedBarrelPosition]) {
				maxChainedBarrelPosition = barrelPosition
			}
		}

		barrelCount += len(cache[maxChainedBarrelPosition])

		for _, barrelPosition := range cache[maxChainedBarrelPosition] {
			delete(barrelsRemaining, barrelPosition)
		}
	}

	return strconv.Itoa(barrelCount)
}

func main() {
	submit := flag.Bool("submit", false, "Submit answers instead of dry run")
	flag.Parse()

	utils.Run(solvePart1, solvePart2, solvePart3, *submit)
}
