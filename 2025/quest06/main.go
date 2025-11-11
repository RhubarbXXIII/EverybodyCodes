package main

import (
	"flag"
	"maps"
	"slices"
	"strconv"
	"unicode"

	"everybodycodes/utils"
)

func solvePart1(input string) string {
	mentorCount := 0
	pairingCount := 0

	for _, code := range input {
		switch code {
		case 'A':
			mentorCount++
		case 'a':
			pairingCount += mentorCount
		}
	}

	return strconv.Itoa(pairingCount)
}

func solvePart2(input string) string {
	mentorCounts := make(map[rune]int)
	pairingCounts := make(map[rune]int)

	for _, code := range input {
		if unicode.IsUpper(code) {
			mentorCounts[code] += 1
		} else if unicode.IsLower(code) {
			pairingCounts[code] += mentorCounts[unicode.ToUpper(code)]
		}
	}

	totalPairingCount := 0
	for _, pairingCount := range slices.Collect(maps.Values(pairingCounts)) {
		totalPairingCount += pairingCount
	}

	return strconv.Itoa(totalPairingCount)
}

func solvePart3(input string) string {
	distanceLimit := 1000
	patternRepetitionCount := 1000

	pattern := []rune(input)
	patternLength := len(pattern)
	patternLengthActual := patternRepetitionCount * patternLength

	cache := map[rune]map[int]int{'a': {}, 'b': {}, 'c': {}}

	totalPairingCount := 0

	for i := range patternLengthActual {
		currentCode := pattern[i%patternLength]
		if unicode.IsUpper(currentCode) {
			continue
		}

		if i-distanceLimit >= 0 && i+distanceLimit < patternLengthActual {
			if cachedPairingCount, present := cache[currentCode][i%patternLength]; present {
				totalPairingCount += cachedPairingCount
				continue
			}
		}

		pairingCount := 0
		for j := max(i-distanceLimit, 0); j <= min(i+distanceLimit, patternLengthActual-1); j++ {
			if pattern[j%patternLength] == unicode.ToUpper(currentCode) {
				pairingCount++
			}
		}

		totalPairingCount += pairingCount

		if i-distanceLimit >= 0 && i+distanceLimit < patternLengthActual {
			cache[currentCode][i%patternLength] = pairingCount
		}
	}

	return strconv.Itoa(totalPairingCount)
}

func main() {
	submit := flag.Bool("submit", false, "Submit answers instead of dry run")
	flag.Parse()

	utils.Run(solvePart1, solvePart2, solvePart3, *submit)
}
