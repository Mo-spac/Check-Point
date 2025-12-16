package problems_3

func RepeatAlpha(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		r := s[i]
		var count int
		if 'A' <= r && r <= 'Z' {
			count = int(r - 'A' + 1)
		} else if 'a' <= r && r <= 'z' {
			count = int(r - 'a' + 1)
		} else {
			count = 1
		}
		for j := 0; j < count; j++ {
			result += string(r)
		}
	}
	return result
}

// func RepeatAlpha(s string) string {
// 	result := ""

// 	for _, r := range s {
// 		var count int
// 		if r >= 'A' && r <= 'Z' {
// 			count = int(r - 'A' + 1)
// 		} else if r >= 'a' && r <= 'z' {
// 			count = int(r - 'a' + 1)
// 		} else {
// 			count = 1
// 		}
// 		for j := 0; j < count; j++ {
// 			result += string(r)
// 		}
// 	}
// 	return result
// }

// func RepeatAlpha(s string) string {
// 	result := []rune{s}
// 	for _, r := range s {
// 		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
// 			var count int
// 			if r >= 'A' && r <= 'Z' {
// 				count = int(r - 'A' + 1)
// 			} else {
// 				count = int(r - 'a' + 1)
// 			}
// 			for i := 0; i < count; i++ {
// 				result = append(result, r)
// 			}
// 		}
// 	}
// 	return string(result)
// }
