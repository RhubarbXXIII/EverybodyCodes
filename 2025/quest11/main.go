package main

import (
	"flag"
	"strconv"
	"strings"

	"everybodycodes/utils"
)

func updatePhase1(columns []int) bool {
	updated := false

	for i := range len(columns) - 1 {
		if columns[i] > columns[i+1] {
			columns[i]--
			columns[i+1]++

			updated = true
		}
	}

	return updated
}

func updatePhase2(columns []int) bool {
	updated := false

	for i := range len(columns) - 1 {
		if columns[i] < columns[i+1] {
			columns[i]++
			columns[i+1]--

			updated = true
		}
	}

	return updated
}

func solvePart1(input string) string {
	columns, _ := utils.ParseInts(strings.Split(input, "\n"))

	roundCount := 10
	phaseIndex := 1
	for range roundCount {
		if phaseIndex == 1 {
			if updated := updatePhase1(columns); !updated {
				phaseIndex++
			}
		}

		if phaseIndex == 2 {
			if updated := updatePhase2(columns); !updated {
				break
			}
		}
	}

	checksum := 0
	for i, column := range columns {
		checksum += (i + 1) * column
	}

	return strconv.Itoa(checksum)
}

func solvePart2(input string) string {
	columns, _ := utils.ParseInts(strings.Split(input, "\n"))

	roundCount := 0

	phaseIndex := 1
	for ; ; roundCount++ {
		if phaseIndex == 1 {
			if updated := updatePhase1(columns); !updated {
				phaseIndex++
			}
		}

		if phaseIndex == 2 {
			if updated := updatePhase2(columns); !updated {
				break
			}
		}
	}

	return strconv.Itoa(roundCount)
}

func solvePart3(input string) string {
	columns, _ := utils.ParseInts(strings.Split(input, "\n"))

	roundCount := 0

	birdCount := 0
	for _, birdColumnCount := range columns {
		birdCount += birdColumnCount
	}

	finalBirdColumnCount := birdCount / len(columns)

	for _, birdColumnCount := range columns {
		roundCount += max(birdColumnCount-finalBirdColumnCount, 0)
	}

	return strconv.Itoa(roundCount)
}

func main() {
	submit := flag.Bool("submit", false, "Submit answers instead of dry run")
	flag.Parse()

	utils.Run(solvePart1, solvePart2, solvePart3, *submit)
}
