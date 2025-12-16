package problems_4

func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}

	for n := nb; n >= 2; n-- {
		isPrim := true
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				isPrim = false
				break
			}
		}
		if isPrim {
			return n
		}
	}
	return 0
}

// *******************

func FindPrevPrime1(nb int) int {
	if nb < 2 {
		return 0
	}

	for i := nb; i >= 2; i-- {
		if isPrim1(i) {
			return i
		}
	}
	return 0
}

func isPrim1(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//*************************

func FindPrevPrime2(nb int) int {
	if nb < 2 {
		return 0
	}

	for i := nb; i >= 2; i-- {
		if isPrim2(i) {
			return i
		}
	}
	return 0
}

func isPrim2(n int) bool {
	if n == 2 {
		return true
	}

	if n < 2 {
		return false
	}

	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
