package utils

import (
	"strconv"
)

func ParseInts(ss []string) ([]int, error) {
	ints := make([]int, len(ss))
	var ok error

	for i, s := range ss {
		ints[i], ok = strconv.Atoi(s)
		if ok != nil {
			return ints[:i], ok
		}
	}

	return ints, nil
}
