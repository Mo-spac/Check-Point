package problems_3

func Gcd(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Gcd1(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd1(b, a%b)
}

func Gcd2(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	min := a
	if b < a {
		min = b
	}
	gcd := 0
	for i := 1; i <= min; i++ {
		if a%i == 0 && b%i == 0 {
			gcd = i
		}
	}
	return gcd
}
