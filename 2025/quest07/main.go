package main

import (
	"flag"
	"strconv"
	"strings"

	"everybodycodes/utils"
)

func parseInput(input string) (names []string, rules map[rune]map[rune]bool) {
	inputLines := strings.Split(input, "\n")

	names = strings.Split(inputLines[0], ",")
	rules = map[rune]map[rune]bool{}

	for _, inputLine := range inputLines[2:] {
		inputLineSplit := strings.Split(inputLine, " > ")

		start := inputLineSplit[0][0]
		endsStrings := strings.Split(inputLineSplit[1], ",")

		ends := map[rune]bool{}
		for _, endString := range endsStrings {
			ends[rune(endString[0])] = true
		}

		rules[rune(start)] = ends
	}

	return
}

func solvePart1(input string) string {
	names, rules := parseInput(input)

nameLoop:
	for _, name := range names {
		for i, letter := range name[:len(name)-1] {
			if _, present := rules[letter][rune(name[i+1])]; !present {
				continue nameLoop
			}
		}

		return name
	}

	panic("No name found")
}

func solvePart2(input string) string {
	names, rules := parseInput(input)

	sum := 0

nameLoop:
	for n, name := range names {
		for i, letter := range name[:len(name)-1] {
			if _, present := rules[letter][rune(name[i+1])]; !present {
				continue nameLoop
			}
		}

		sum += n + 1
	}

	return strconv.Itoa(sum)
}

func solvePart3(input string) string {
	prefixes, rules := parseInput(input)

	uniqueNames := map[string]bool{}

	var generateNames func(string) []string
	generateNames = func(name string) []string {
		if len(name) >= 11 {
			return []string{name}
		}

		names := []string{}

		for letter := range rules[rune(name[len(name)-1])] {
			newName := append([]rune(name), letter)
			if _, present := uniqueNames[string(newName)]; present {
				continue
			}

			if len(newName) >= 7 {
				uniqueNames[string(newName)] = true
			}

			for _, generatedName := range generateNames(string(newName)) {
				uniqueNames[generatedName] = true
			}
		}

		return names
	}

nameLoop:
	for _, prefix := range prefixes {
		for i, letter := range prefix[:len(prefix)-1] {
			if _, present := rules[letter][rune(prefix[i+1])]; !present {
				continue nameLoop
			}
		}

		generateNames(prefix)
	}

	return strconv.Itoa(len(uniqueNames))
}

func main() {
	submit := flag.Bool("submit", false, "Submit answers instead of dry run")
	flag.Parse()

	utils.Run(solvePart1, solvePart2, solvePart3, *submit)
}
