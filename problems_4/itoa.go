package problems_4

func Itoa(n int) string {
	if n == 0 {
		return ""
	}

	isNegative := false

	if n < 0 {
		n = -n
		isNegative = true
	}

	result := ""
	for n > 0 {
		digit := n % 10
		result = string(rune('0'+digit)) + result
		n /= 10
	}
	if isNegative {
		result = "-" + result
	}
	return result
}
