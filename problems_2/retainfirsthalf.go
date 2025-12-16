package problems_2

func RetainFirstHalf(str string) string {
	runes := []rune(str)
	if len(runes) == 0 {
		return ""
	} else if len(runes) == 1 {
		return str
	} else {
		half := len(runes) / 2
		return string(runes[:half])
	}
}

// func RetainFirstHalf(str string) string {
// 	if str == "" {
// 		return ""
// 	} else if len(str) == 1 {
// 		return str
// 	} else {
// 		return str[0 : len(str)/2]
// 	}
// }
