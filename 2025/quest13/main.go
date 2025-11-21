package main

import (
	"container/list"
	"flag"
	"strconv"
	"strings"

	"everybodycodes/utils"
)

type dialRange struct {
	start  int
	length int

	reversed bool
}

func parseInput(input string) [][2]int {
	inputLines := strings.Split(input, "\n")

	dialRangeNumbers := make([][2]int, len(inputLines))
	for i, inputLine := range inputLines {
		dialRangeNumbers[i][0], _ = strconv.Atoi(strings.Split(inputLine, "-")[0])
		dialRangeNumbers[i][1], _ = strconv.Atoi(strings.Split(inputLine, "-")[1])
	}

	return dialRangeNumbers
}

func solvePart1(input string) string {
	dialNumbersList, _ := utils.ParseInts(strings.Split(input, "\n"))

	dialNumbers := make([]int, len(dialNumbersList)+1)
	dialNumbers[0] = 1

	frontIndex := 1
	backIndex := len(dialNumbers) - 1
	listIndex := 0

	for frontIndex <= backIndex {
		if listIndex%2 == 0 {
			dialNumbers[frontIndex] = dialNumbersList[listIndex]

			frontIndex++
		} else {
			dialNumbers[backIndex] = dialNumbersList[listIndex]

			backIndex--
		}

		listIndex++
	}

	return strconv.Itoa(dialNumbers[2025%len(dialNumbers)])
}

func solvePart2(input string) string {
	dialRangesNumbers := parseInput(input)

	dialRanges := list.List{}
	dialRangesBack := list.List{}

	dialRanges.PushBack(dialRange{start: 1, length: 1, reversed: false})
	dialNumberCount := 1

	for i, dialRangeNumbers := range dialRangesNumbers {
		dialRange := dialRange{length: dialRangeNumbers[1] - dialRangeNumbers[0] + 1}

		if i%2 == 0 {
			dialRange.start = dialRangeNumbers[0]
			dialRange.reversed = false

			dialRanges.PushBack(dialRange)
		} else {
			dialRange.start = dialRangeNumbers[1]
			dialRange.reversed = true

			dialRangesBack.PushFront(dialRange)
		}

		dialNumberCount += dialRange.length
	}

	for e := dialRangesBack.Front(); e != nil; e = e.Next() {
		dialRanges.PushBack(e.Value)
	}

	goalIndex := 20252025 % dialNumberCount
	goal := -1

	index := 0
	for e := dialRanges.Front(); e != nil; e = e.Next() {
		dialRange := e.Value.(dialRange)

		if index+dialRange.length > goalIndex {
			if dialRange.reversed {
				goal = dialRange.start - (goalIndex - index)
			} else {
				goal = dialRange.start + (goalIndex - index)
			}

			break
		}

		index += dialRange.length
	}

	return strconv.Itoa(goal)
}

func solvePart3(input string) string {
	dialRangesNumbers := parseInput(input)

	dialRanges := list.List{}
	dialRangesBack := list.List{}

	dialRanges.PushBack(dialRange{start: 1, length: 1, reversed: false})
	dialNumberCount := 1

	for i, dialRangeNumbers := range dialRangesNumbers {
		dialRange := dialRange{length: dialRangeNumbers[1] - dialRangeNumbers[0] + 1}

		if i%2 == 0 {
			dialRange.start = dialRangeNumbers[0]
			dialRange.reversed = false

			dialRanges.PushBack(dialRange)
		} else {
			dialRange.start = dialRangeNumbers[1]
			dialRange.reversed = true

			dialRangesBack.PushFront(dialRange)
		}

		dialNumberCount += dialRange.length
	}

	for e := dialRangesBack.Front(); e != nil; e = e.Next() {
		dialRanges.PushBack(e.Value)
	}

	goalIndex := 202520252025 % dialNumberCount
	goal := -1

	index := 0
	for e := dialRanges.Front(); e != nil; e = e.Next() {
		dialRange := e.Value.(dialRange)

		if index+dialRange.length > goalIndex {
			if dialRange.reversed {
				goal = dialRange.start - (goalIndex - index)
			} else {
				goal = dialRange.start + (goalIndex - index)
			}

			break
		}

		index += dialRange.length
	}

	return strconv.Itoa(goal)
}

func main() {
	submit := flag.Bool("submit", false, "Submit answers instead of dry run")
	flag.Parse()

	utils.Run(solvePart1, solvePart2, solvePart3, *submit)
}
