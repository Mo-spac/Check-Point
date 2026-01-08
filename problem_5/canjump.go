package problem_5

func CanJump(steps []uint) bool {

	n := len(steps)

	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}

	pos := 0

	for pos < n {
		if steps[pos] == 0 {
			return false
		}

		pos = int(uint(pos) + steps[pos])

		if pos == n-1 {
			return true
		} else if pos > n-1 {
			return false
		}
	}
	return false
}
