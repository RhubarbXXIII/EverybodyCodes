package main

import (
	"flag"

	"everybodycodes/utils"
)

func solvePart1(input string) string {
	return ""
}

func solvePart2(input string) string {
	return ""
}

func solvePart3(input string) string {
	return ""
}

func main() {
	submit := flag.Bool("submit", false, "Submit answers instead of dry run")
	flag.Parse()

	utils.Run(solvePart1, solvePart2, solvePart3, *submit)
}
