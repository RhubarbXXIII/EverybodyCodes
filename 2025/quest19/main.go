package main

import (
	"flag"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"

	"everybodycodes/utils"
)

// func parseInput(input string) [][3]int {
// 	inputLines := strings.Split(input, "\n")

// 	segments := make([][3]int, len(inputLines))
// 	for i := range inputLines {
// 		segmentValues, _ := utils.ParseInts(strings.Split(inputLines[i], ","))
// 		segments[i] = [3]int{segmentValues[0], segmentValues[1], segmentValues[2]}
// 	}

// 	return segments
// }

func parseInput(input string) map[int][][2]int {
	inputLines := strings.Split(input, "\n")

	segments := map[int][][2]int{}
	for i := range inputLines {
		segmentValues, _ := utils.ParseInts(strings.Split(inputLines[i], ","))

		segments[segmentValues[0]] = append(segments[segmentValues[0]], [2]int{segmentValues[1], segmentValues[2]})
	}

	return segments
}

func calculateMinimumFlapCount(segments map[int][][2]int) int {
	paths := map[int]map[int]int{0: {0: 0}}

	currentX := 0

	targetXs := slices.Sorted(maps.Keys(segments))

	for _, targetX := range targetXs {
		for _, segment := range segments[targetX] {
			targetYMin := segment[0]
			targetYMax := segment[0] + segment[1]

			if _, present := paths[targetX]; !present {
				paths[targetX] = map[int]int{}
			}

			distanceX := targetX - currentX
			for currentY, currentFlapCount := range paths[currentX] {
				for targetY := targetYMin; targetY < targetYMax; targetY++ {
					newFlapCount := currentFlapCount

					distanceY := targetY - currentY
					if distanceX%2 != utils.Abs(distanceY)%2 {
						continue
					}

					if utils.Abs(distanceY) > distanceX {
						continue
					}

					if distanceY > 0 {
						newFlapCount += distanceY
					}

					if distanceX > utils.Abs(distanceY) {
						newFlapCount += (distanceX - utils.Abs(distanceY)) / 2
					}

					if previousFlapCount, present := paths[targetX][targetY]; !present || newFlapCount < previousFlapCount {
						paths[targetX][targetY] = newFlapCount
					}
				}
			}
		}

		currentX = targetX
	}

	minimumFlapCount := math.MaxInt
	for _, flapCount := range paths[currentX] {
		if flapCount < minimumFlapCount {
			minimumFlapCount = flapCount
		}
	}

	return minimumFlapCount
}

func solvePart1(input string) string {
	segments := parseInput(input)

	return strconv.Itoa(calculateMinimumFlapCount(segments))
}

func solvePart2(input string) string {
	segments := parseInput(input)

	return strconv.Itoa(calculateMinimumFlapCount(segments))
}

func solvePart3(input string) string {
	segments := parseInput(input)

	return strconv.Itoa(calculateMinimumFlapCount(segments))
}

func main() {
	submit := flag.Bool("submit", false, "Submit answers instead of dry run")
	flag.Parse()

	utils.Run(solvePart1, solvePart2, solvePart3, *submit)
}
