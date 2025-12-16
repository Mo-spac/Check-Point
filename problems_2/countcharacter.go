package problems_2

func CountChar(str string, c rune) int {
	count := 0
	for _, char := range str {
		if c == char {
			count++
		}
	}
	return count
}

// func CountChar(str string, c rune) int {
// 	count := 0
// 	for _, char := range str {
// 		if str == "" {
// 			return 0
// 		} else if char >= 1 && char <= 9 {
// 			return 0
// 		} else if c == char {
// 			count++
// 		}

// 	}
// 	return count
// }
