package main

import (
	"flag"
	"strconv"
	"strings"

	"everybodycodes/utils"
)

func parseInput(input string) (names, commands []string) {
	inputLines := strings.Split(input, "\n")

	names = strings.Split(inputLines[0], ",")
	commands = strings.Split(inputLines[2], ",")
	return
}

func solvePart1(input string) string {
	names, commands := parseInput(input)

	index := 0
	name := names[0]
	for _, command := range commands {
		magnitude, _ := strconv.Atoi(command[1:])

		switch command[0] {
		case 'R':
			index = min(index+magnitude, len(names)-1)
		case 'L':
			index = max(index-magnitude, 0)
		}

		name = names[index]
	}

	return name
}

func solvePart2(input string) string {
	names, commands := parseInput(input)

	index := 0
	name := names[0]
	for _, command := range commands {
		magnitude, _ := strconv.Atoi(command[1:])

		switch command[0] {
		case 'R':
			index = (index + magnitude) % len(names)
		case 'L':
			index = (index - magnitude + len(names)) % len(names)
		}

		name = names[index]
	}

	return name
}

func solvePart3(input string) string {
	names, commands := parseInput(input)

	for _, command := range commands {
		index := 0
		magnitude, _ := strconv.Atoi(command[1:])

		switch command[0] {
		case 'R':
			index = (index + magnitude) % len(names)
		case 'L':
			index = (((index - magnitude) % len(names)) + len(names)) % len(names)
		}

		names[0], names[index] = names[index], names[0]
	}

	return names[0]
}

func main() {
	submit := flag.Bool("submit", false, "Submit answers instead of dry run")
	flag.Parse()

	utils.Run(solvePart1, solvePart2, solvePart3, *submit)
}
