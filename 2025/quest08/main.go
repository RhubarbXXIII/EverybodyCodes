package main

import (
	"flag"
	"slices"
	"strconv"
	"strings"

	"everybodycodes/utils"
)

func solvePart1(input string) string {
	nails, _ := utils.ParseInts(strings.Split(input, ","))
	nailCount := 32

	diameterCount := 0
	for i, currentNail := range nails[:len(nails)-1] {
		nextNail := nails[i+1]

		if (nextNail+nailCount-currentNail)%nailCount == nailCount/2 {
			diameterCount++
		}
	}

	return strconv.Itoa(diameterCount)
}

func solvePart2(input string) string {
	nails, _ := utils.ParseInts(strings.Split(input, ","))
	nailCount := 256

	threads := make([][]int, nailCount+1)
	knotCount := 0
	for i, currentNail := range nails[:len(nails)-1] {
		nextNail := nails[i+1]

		lowerNail := min(currentNail, nextNail)
		higherNail := max(currentNail, nextNail)

		for startNail := lowerNail + 1; startNail < higherNail; startNail++ {
			for _, endNail := range threads[startNail] {
				lowerThreadNail := min(startNail, endNail)
				higherThreadNail := max(startNail, endNail)

				if (lowerThreadNail < lowerNail && higherThreadNail < higherNail) || (lowerThreadNail > lowerNail && higherThreadNail > higherNail) {
					knotCount++
				}
			}
		}

		threads[currentNail] = append(threads[currentNail], nextNail)
		threads[nextNail] = append(threads[nextNail], currentNail)
	}

	return strconv.Itoa(knotCount)
}

func solvePart3(input string) string {
	nails, _ := utils.ParseInts(strings.Split(input, ","))
	nailCount := 256

	threads := make([][]int, nailCount+1)
	for i, currentNail := range nails[:len(nails)-1] {
		nextNail := nails[i+1]

		threads[currentNail] = append(threads[currentNail], nextNail)
		threads[nextNail] = append(threads[nextNail], currentNail)
	}

	maximumCutCount := 0
	for cutStart := 1; cutStart < nailCount; cutStart++ {
		for cutEnd := cutStart + 1; cutEnd <= nailCount; cutEnd++ {
			cutCount := 0

			for threadStart := cutStart + 1; threadStart < cutEnd; threadStart++ {
				for _, threadEnd := range threads[threadStart] {
					threadLower := min(threadStart, threadEnd)
					threadHigher := max(threadStart, threadEnd)

					if (threadLower < cutStart && threadHigher < cutEnd) || (threadLower > cutStart && threadHigher > cutEnd) {
						cutCount++
					}
				}
			}

			if slices.Contains(threads[cutStart], cutEnd) {
				cutCount++
			}

			maximumCutCount = max(maximumCutCount, cutCount)
		}
	}

	return strconv.Itoa(maximumCutCount)
}

func main() {
	submit := flag.Bool("submit", false, "Submit answers instead of dry run")
	flag.Parse()

	utils.Run(solvePart1, solvePart2, solvePart3, *submit)
}
