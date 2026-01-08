package problem_5

func ConcatAlternate(slice1, slice2 []int) []int {

	var result []int

	first := slice1
	second := slice2

	if len(slice2) > len(slice1) {
		first = slice2
		second = slice1
	}

	for i := 0; i < len(second); i++ {
		result = append(result, first[i], second[i])
	}
	result = append(result, first[len(second):]...)
	return result
}

// *********************************

func ConcatAlternate1(slice1, slice2 []int) []int {
	var result []int

	maxLen := len(slice1)
	if len(slice2) > maxLen {
		maxLen = len(slice2)
	}

	for i := 0; i < maxLen; i++ {
		// لو slice1 هي الأطول أو الطول متساوي
		if len(slice1) >= len(slice2) {
			if i < len(slice1) {
				result = append(result, slice1[i])
			}
			if i < len(slice2) {
				result = append(result, slice2[i])
			}
		} else { // slice2 أطول
			if i < len(slice2) {
				result = append(result, slice2[i])
			}
			if i < len(slice1) {
				result = append(result, slice1[i])
			}
		}
	}
	return result
}

// *********************************

func ConcatAlternate2(slice1, slice2 []int) []int {
	var result []int
	i := 0

	for i < len(slice1) || i < len(slice2) {

		if len(slice1) >= len(slice2) {
			if i < len(slice1) {
				result = append(result, slice1[i])
			}
			if i < len(slice2) {
				result = append(result, slice2[i])
			}
		} else {
			if i < len(slice2) {
				result = append(result, slice2[i])
			}
			if i < len(slice1) {
				result = append(result, slice1[i])
			}
		}

		i++
	}
	return result
}
