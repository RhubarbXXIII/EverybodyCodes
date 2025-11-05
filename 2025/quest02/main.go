package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"

	"everybodycodes/utils"
)

type complex struct {
	x int
	y int
}

func (c complex) String() string {
	return fmt.Sprintf("[%d,%d]", c.x, c.y)
}

func add(left, right complex) complex {
	return complex{left.x + right.x, left.y + right.y}
}

func multiply(left, right complex) complex {
	return complex{left.x*right.x - left.y*right.y, left.x*right.y + left.y*right.x}
}

func divide(left, right complex) complex {
	return complex{left.x / right.x, left.y / right.y}
}

func parseInput(input string) complex {
	matches := regexp.MustCompile(`(-?\d+)`).FindAllString(input, 2)
	x, _ := strconv.Atoi(matches[0])
	y, _ := strconv.Atoi(matches[1])
	return complex{x, y}
}

func solvePart1(input string) string {
	a := parseInput(input)
	r := complex{}

	for range 3 {
		r = multiply(r, r)
		r = divide(r, complex{10, 10})
		r = add(r, a)
	}

	return r.String()
}

func solvePart2(input string) string {
	a0 := parseInput(input)

	engraving_point_count := 0

	for dx := 0; dx <= 1000; dx += 10 {
	point:
		for dy := 0; dy <= 1000; dy += 10 {
			a := add(a0, complex{dx, dy})
			r := complex{}

			for range 100 {
				r = multiply(r, r)
				r = divide(r, complex{100000, 100000})
				r = add(r, a)

				if r.x < -1000000 || r.x > 1000000 || r.y < -1000000 || r.y > 1000000 {
					continue point
				}
			}

			engraving_point_count++
		}
	}

	return strconv.Itoa(engraving_point_count)
}

func solvePart3(input string) string {
	a0 := parseInput(input)

	engraving_point_count := 0

	for dx := 0; dx <= 1000; dx += 1 {
	point:
		for dy := 0; dy <= 1000; dy += 1 {
			a := add(a0, complex{dx, dy})
			r := complex{}

			for range 100 {
				r = multiply(r, r)
				r = divide(r, complex{100000, 100000})
				r = add(r, a)

				if r.x < -1000000 || r.x > 1000000 || r.y < -1000000 || r.y > 1000000 {
					continue point
				}
			}

			engraving_point_count++
		}
	}

	return strconv.Itoa(engraving_point_count)
}

func main() {
	submit := flag.Bool("submit", false, "Submit answers instead of dry run")
	flag.Parse()

	utils.Run(solvePart1, solvePart2, solvePart3, *submit)
}
