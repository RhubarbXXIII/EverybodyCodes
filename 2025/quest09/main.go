package main

import (
	"flag"
	"strconv"
	"strings"

	"everybodycodes/utils"
)

func parseInput(input string) []string {
	inputLines := strings.Split(input, "\n")

	dnas := make([]string, len(inputLines))
	for i, line := range inputLines {
		dnas[i] = strings.Split(line, ":")[1]
	}

	return dnas
}

func isValidChild(childDna, parent1Dna, parent2Dna string) bool {
	parent1DnaRunes := []rune(parent1Dna)
	parent2DnaRunes := []rune(parent2Dna)
	for i, childDnaRune := range childDna {
		if childDnaRune != parent1DnaRunes[i] && childDnaRune != parent2DnaRunes[i] {
			return false
		}
	}

	return true
}

func calculateSimilarity(childDna, parent1Dna, parent2Dna string) int {
	similarity := 1

	for _, parentDna := range [2]string{parent1Dna, parent2Dna} {
		matchCount := 0

		parentDnaRunes := []rune(parentDna)
		for i, childDnaRune := range childDna {
			if childDnaRune == parentDnaRunes[i] {
				matchCount++
			}
		}

		similarity *= matchCount
	}

	return similarity
}

type duck struct {
	id       int
	familyId int

	dna string

	parent1 *duck
	parent2 *duck
}

func solvePart1(input string) string {
	dnas := parseInput(input)
	dnaCount := len(dnas)

	childId := -1
	parentIds := [2]int{}
	for i := range 3 {
		if isValidChild(dnas[i], dnas[(i+1)%dnaCount], dnas[(i+2)%dnaCount]) {
			childId = i
			parentIds = [2]int{(i + 1) % dnaCount, (i + 2) % dnaCount}
			break
		}
	}

	return strconv.Itoa(calculateSimilarity(dnas[childId], dnas[parentIds[0]], dnas[parentIds[1]]))
}

func solvePart2(input string) string {
	dnas := parseInput(input)

	similaritySum := 0
	for childId, childDna := range dnas {
		for parent1Id, parent1Dna := range dnas {
			if childId == parent1Id {
				continue
			}

			for parent2Id, parent2Dna := range dnas {
				if parent2Id <= parent1Id {
					continue
				}

				if childId == parent2Id {
					continue
				}

				if isValidChild(childDna, parent1Dna, parent2Dna) {
					similaritySum += calculateSimilarity(childDna, parent1Dna, parent2Dna)
				}
			}
		}
	}

	return strconv.Itoa(similaritySum)
}

func solvePart3(input string) string {
	dnas := parseInput(input)

	ducks := make([]*duck, len(dnas))
	for i, dna := range dnas {
		ducks[i] = &duck{id: i, familyId: -1, dna: dna}
	}

	for childId, childDna := range dnas {
		for parent1Id, parent1Dna := range dnas {
			if childId == parent1Id {
				continue
			}

			for parent2Id, parent2Dna := range dnas {
				if parent2Id <= parent1Id {
					continue
				}

				if childId == parent2Id {
					continue
				}

				if !isValidChild(childDna, parent1Dna, parent2Dna) {
					continue
				}

				parent1 := ducks[parent1Id]
				parent2 := ducks[parent2Id]
				child := ducks[childId]
				child.parent1 = parent1
				child.parent2 = parent2
			}
		}
	}

	duckFamiliesAssigned := map[*duck]bool{}
	duckFamilies := map[int]map[*duck]bool{}

	var assignFamily func(*duck, int)
	assignFamily = func(child *duck, familyId int) {
		duckFamiliesAssigned[child] = true

		child.familyId = familyId

		duckFamilies[familyId][child] = true

		for _, parent := range [2]*duck{child.parent1, child.parent2} {
			if parent == nil {
				continue
			}

			if parent.familyId == familyId {
				continue
			}

			if parent.familyId < 0 {
				assignFamily(parent, familyId)
				continue
			}

			oldFamilyId := parent.familyId

			for duck := range duckFamilies[oldFamilyId] {
				duck.familyId = familyId

				duckFamilies[familyId][duck] = true
			}

			delete(duckFamilies, oldFamilyId)
		}
	}

	for _, child := range ducks {
		if duckFamiliesAssigned[child] {
			continue
		}

		duckFamilies[child.id] = map[*duck]bool{}

		assignFamily(child, child.id)
	}

	largestFamilyId := -1
	for familyId, family := range duckFamilies {
		if largestFamilyId < 0 || len(family) > len(duckFamilies[largestFamilyId]) {
			largestFamilyId = familyId
		}
	}

	scaleSum := 0
	for duck := range duckFamilies[largestFamilyId] {
		scaleSum += duck.id + 1
	}

	return strconv.Itoa(scaleSum)
}

func main() {
	submit := flag.Bool("submit", false, "Submit answers instead of dry run")
	flag.Parse()

	utils.Run(solvePart1, solvePart2, solvePart3, *submit)
}
