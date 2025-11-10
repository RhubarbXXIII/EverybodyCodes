package main

import (
	"flag"
	"maps"
	"slices"
	"sort"
	"strconv"
	"strings"

	"everybodycodes/utils"
)

type swordInfo struct {
	id      int
	quality int
	levels  []*int
}

func parseInputLine(input string) (id int, numbers []int) {
	inputSplit := strings.Split(input, ":")

	id, _ = strconv.Atoi(inputSplit[0])

	numbersStrings := strings.Split(inputSplit[1], ",")
	numbers = make([]int, len(numbersStrings))

	for i := range len(numbersStrings) {
		number, _ := strconv.Atoi(numbersStrings[i])
		numbers[i] = number
	}

	return
}

func calculateSwordInfo(id int, numbers []int) swordInfo {
	leftSegments := make([]*int, 1)
	spineSegments := make([]*int, 1)
	rightSegments := make([]*int, 1)

	spineSegments[0] = &numbers[0]

construction:
	for _, number := range numbers[1:] {
		for i, leftSegment := range leftSegments {
			if leftSegment == nil && number < *spineSegments[i] {
				leftSegments[i] = &number
				continue construction
			}
		}

		for i, rightSegment := range rightSegments {
			if rightSegment == nil && number > *spineSegments[i] {
				rightSegments[i] = &number
				continue construction
			}
		}

		leftSegments = append(leftSegments, nil)
		spineSegments = append(spineSegments, &number)
		rightSegments = append(rightSegments, nil)
	}

	qualityString := ""
	for _, segment := range spineSegments {
		segmentString := strconv.Itoa(*segment)

		qualityString += segmentString
	}

	quality, _ := strconv.Atoi(qualityString)

	levels := make([]*int, len(spineSegments))
	for i := range spineSegments {
		levelString := ""

		if leftSegments[i] != nil {
			levelString += strconv.Itoa(*leftSegments[i])
		}
		levelString += strconv.Itoa(*spineSegments[i])
		if rightSegments[i] != nil {
			levelString += strconv.Itoa(*rightSegments[i])
		}

		level, _ := strconv.Atoi(levelString)
		levels[i] = &level
	}

	return swordInfo{id: id, quality: quality, levels: levels}
}

func solvePart1(input string) string {
	id, numbers := parseInputLine(input)

	swordInfo := calculateSwordInfo(id, numbers)
	return strconv.Itoa(swordInfo.quality)
}

func solvePart2(input string) string {
	inputLines := strings.Split(input, "\n")

	qualities := make([]int, len(inputLines))

	for i, inputLine := range inputLines {
		id, numbers := parseInputLine(inputLine)

		swordInfo := calculateSwordInfo(id, numbers)
		qualities[i] = swordInfo.quality
	}

	sort.Ints(qualities)

	return strconv.Itoa(qualities[len(qualities)-1] - qualities[0])
}

func solvePart3(input string) string {
	inputLines := strings.Split(input, "\n")

	swordInfos := make(map[int]swordInfo)
	for _, inputLine := range inputLines {
		id, numbers := parseInputLine(inputLine)

		swordInfos[id] = calculateSwordInfo(id, numbers)
	}

	ids := slices.Collect(maps.Keys(swordInfos))

	compare := func(left, right int) int {
		leftSwordInfo := swordInfos[left]
		rightSwordInfo := swordInfos[right]

		if leftSwordInfo.quality != rightSwordInfo.quality {
			return rightSwordInfo.quality - leftSwordInfo.quality
		}

		for i := 0; i < len(leftSwordInfo.levels) && i < len(rightSwordInfo.levels); i++ {
			if *leftSwordInfo.levels[i] != *rightSwordInfo.levels[i] {
				return *rightSwordInfo.levels[i] - *leftSwordInfo.levels[i]
			}
		}

		return rightSwordInfo.id - leftSwordInfo.id
	}

	slices.SortFunc(ids, compare)

	checksum := 0
	for i, id := range ids {
		checksum += (i + 1) * id
	}

	return strconv.Itoa(checksum)
}

func main() {
	submit := flag.Bool("submit", false, "Submit answers instead of dry run")
	flag.Parse()

	utils.Run(solvePart1, solvePart2, solvePart3, *submit)
}
