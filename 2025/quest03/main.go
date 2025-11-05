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

func parseInput(input string) []int {
	numberStrings := strings.Split(input, ",")

	numbers := make([]int, len(numberStrings))
	for i, s := range numberStrings {
		numbers[i], _ = strconv.Atoi(s)
	}

	return numbers
}

func solvePart1(input string) string {
	crateSizes := parseInput(input)

	sort.Sort(sort.Reverse(sort.IntSlice(crateSizes)))

	setSize := 0
	for i := range crateSizes {
		if i > 0 && crateSizes[i] == crateSizes[i-1] {
			continue
		}

		setSize += crateSizes[i]
	}

	return strconv.Itoa(setSize)
}

func solvePart2(input string) string {
	crateSizes := parseInput(input)

	sort.Ints(crateSizes)

	setSize := 0
	setCount := 0
	for i := 0; i < len(crateSizes)-1; i++ {
		if i > 0 && crateSizes[i] == crateSizes[i-1] {
			continue
		}

		setSize += crateSizes[i]
		setCount++

		if setCount == 20 {
			break
		}
	}

	return strconv.Itoa(setSize)
}

func solvePart3(input string) string {
	crateSizes := parseInput(input)

	crateSizeCounts := make(map[int]int)
	for _, n := range crateSizes {
		crateSizeCounts[n]++
	}

	return strconv.Itoa(slices.Max(slices.Collect(maps.Values(crateSizeCounts))))
}

func main() {
	submit := flag.Bool("submit", false, "Submit answers instead of dry run")
	flag.Parse()

	utils.Run(solvePart1, solvePart2, solvePart3, *submit)
}
