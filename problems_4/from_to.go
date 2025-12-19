package problems_4

import "strconv"

func FromTo(from int, to int) string {

	result := ""
	if (0 > from || from > 99) || (0 > to || to > 99) {
		return "Invalid\n"
	}

	if from <= to {
		for i := from; i <= to; i++ {
			if i != from {
				result += ", "
			}

			result += twoDigits(i)
		}
	} else {
		for i := from; i >= to; i-- {
			if i != from {
				result += ", "
			}

			result += twoDigits(i)
		}
	}

	return result + "\n"
}

func twoDigits(n int) string {
	if n < 10 {
		return "0" + string(rune('0'+n))
	}
	return string(rune('0'+n/10)) + string(rune('0'+n%10))
}

//*************************************

func FromTo1(from int, to int) string {
	result := ""
	if (from < 0 || from > 99) || (0 > to || to > 99) {
		return "Invalid\n"
	}

	if from <= to {
		for i := from; i <= to; i++ {
			s := strconv.Itoa(i)
			if i != from {
				result += ", "
			}
			if i < 10 {
				result += "0"
			}
			result += s
		}
	} else {
		for i := from; i >= to; i-- {
			s := strconv.Itoa(i)
			if i != from {
				result += ", "
			}
			if 0 <= i && i < 10 {
				result += "0"
			}
			result += s
		}
	}

	return result + "\n"
}
