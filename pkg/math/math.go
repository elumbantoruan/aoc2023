package math

func GdList(l []int) int {
	g := l[1]
	for _, i := range l[1:] {
		g = Gcd(g, i)
	}
	return g
}

func Gcd(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

func Lcm(a, b int) int {
	return a * b / Gcd(a, b)
}

func LcmList(l []int) int {
	lcm := l[1]
	for _, i := range l[1:] {
		lcm = Lcm(lcm, i)
	}
	return lcm
}

func GreatestCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LeastCommonMultiple(a, b int, integers ...int) int {
	result := a * b / GreatestCommonDivisor(a, b)
	for i := 0; i < len(integers); i++ {
		result = LeastCommonMultiple(result, integers[i])
	}
	return result
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LeastCommonMultipleList(numbers []int) int {
	lcm := numbers[0]
	for _, number := range numbers[1:] {
		lcm = lcm * number / gcd(lcm, number)
	}
	return lcm
}
