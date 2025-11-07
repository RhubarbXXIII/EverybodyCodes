package main

import (
	"flag"
	"strconv"
	"strings"

	"everybodycodes/utils"
)

func parseInput(input string) []int {
	inputLines := strings.Split(input, "\n")

	sizes := make([]int, len(inputLines))
	for i := range sizes {
		sizes[i], _ = strconv.Atoi(inputLines[i])
	}

	return sizes
}

func solvePart1(input string) string {
	gearSizes := parseInput(input)

	return strconv.Itoa(gearSizes[0] * 2025 / gearSizes[len(gearSizes)-1])
}

func solvePart2(input string) string {
	gearSizes := parseInput(input)

	turnCount := gearSizes[len(gearSizes)-1] * 10000000000000 / gearSizes[0]
	if gearSizes[len(gearSizes)-1]*10000000000000%gearSizes[0] > 0 {
		turnCount++
	}

	// return strconv.FormatInt(turnCount, 10)
	return strconv.Itoa(turnCount)
}

func solvePart3(input string) string {
	inputLines := strings.Split(input, "\n")

	firstGearSize, _ := strconv.Atoi(inputLines[0])
	lastGearSize, _ := strconv.Atoi(inputLines[len(inputLines)-1])

	otherGearSizes := make([][]int, len(inputLines)-2)
	for i := range inputLines[1 : len(inputLines)-1] {
		gearLine := strings.Split(inputLines[i+1], "|")

		gearSize := make([]int, 2)
		for i := range 2 {
			gearSize[i], _ = strconv.Atoi(gearLine[i])
		}

		otherGearSizes[i] = gearSize
	}

	turnCount := 100 * firstGearSize
	for _, gearSize := range otherGearSizes {
		turnCount *= gearSize[1] / gearSize[0]
	}

	return strconv.Itoa(turnCount / lastGearSize)
}

func main() {
	submit := flag.Bool("submit", false, "Submit answers instead of dry run")
	flag.Parse()

	utils.Run(solvePart1, solvePart2, solvePart3, *submit)
}
