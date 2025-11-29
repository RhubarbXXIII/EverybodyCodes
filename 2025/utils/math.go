package utils

func Abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func Lcm(ns ...int) int {
	if len(ns) == 0 {
		return 1
	} else if len(ns) == 1 {
		return ns[0]
	}

	lcm := 1
	for _, n := range ns {
		lcm = lcm * n / Gcd(lcm, n)
	}

	return lcm
}
