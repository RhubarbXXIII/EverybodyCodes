package main

import (
	"flag"
	"strconv"
	"strings"

	"everybodycodes/utils"
)

func findSpell(wall []int) []int {
	currentWall := make([]int, len(wall))
	copy(currentWall, wall)

	spell := []int{}

decreaseLoop:
	for i := 1; i < len(wall); {
		nextWall := make([]int, len(wall))
		copy(nextWall, currentWall)

		for j := i - 1; j < len(wall); j += i {
			if nextWall[j] == 0 {
				i++
				continue decreaseLoop
			}

			nextWall[j]--
		}

		spell = append(spell, i)

		currentWall = nextWall

		for _, column := range nextWall {
			if column != 0 {
				continue decreaseLoop
			}
		}

		break
	}

	return spell
}

func countBlocksInSpell(spell []int, length int) int {
	blockCount := 0

	for _, layer := range spell {
		blockCount += length / layer
	}

	return blockCount
}

func solvePart1(input string) string {
	spell, _ := utils.ParseInts(strings.Split(input, ","))

	return strconv.Itoa(countBlocksInSpell(spell, 90))
}

func solvePart2(input string) string {
	wall, _ := utils.ParseInts(strings.Split(input, ","))

	spell := findSpell(wall)

	spellProduct := 1
	for _, spellStep := range spell {
		spellProduct *= spellStep
	}

	return strconv.Itoa(spellProduct)
}

func solvePart3(input string) string {
	wall, _ := utils.ParseInts(strings.Split(input, ","))

	totalBlockCount := 202520252025000

	spell := findSpell(wall)

	wallLength := -1
	wallLengthGuessLow := len(wall) - 1
	wallLengthGuessHigh := totalBlockCount + 1

	for wallLength < 0 {
		wallLengthGuess := wallLengthGuessLow + (wallLengthGuessHigh-wallLengthGuessLow)/2

		blockCountLow := countBlocksInSpell(spell, wallLengthGuess)
		blockCountHigh := countBlocksInSpell(spell, wallLengthGuess+1)

		if blockCountLow <= totalBlockCount && blockCountHigh > totalBlockCount {
			wallLength = wallLengthGuess
			break
		}

		if blockCountLow > totalBlockCount {
			wallLengthGuessHigh = wallLengthGuess
		} else {
			wallLengthGuessLow = wallLengthGuess
		}
	}

	return strconv.Itoa(wallLength)
}

func main() {
	submit := flag.Bool("submit", false, "Submit answers instead of dry run")
	flag.Parse()

	utils.Run(solvePart1, solvePart2, solvePart3, *submit)
}
