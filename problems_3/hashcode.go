package problems_3

func HashCode(dec string) string {
	n := len(dec)
	result := []rune{}
	for _, r := range dec {
		h := (int(r) + n) % 127
		if h < 33 {
			h += 33
		}
		result = append(result, rune(h))
	}
	return string(result)
}

// *****************

func HashCode1(dec string) string {
	n := len(dec)
	result := ""

	for _, r := range dec {
		h := (int(r) + n) % 127
		if h < 33 {
			h += 33
		}
		result += string(rune(h))
	}
	return result
}

// func HashCode(dec string) int {
// 	n := len(dec)
// 	result := []rune{}

// 	for _, r := range dec {
// 		h := (int(r) + n) % 127
// 		if h < 32 {
// 			h += 32
// 		}
// 		result = append(result, rune(h))
// 	}
// 	return int(result[0])
// }

// func HashCode(dec string) int {
// 	hash := 0
// 	for i := 0; i < len(dec); i++ {
// 		hash = (hash*31 + int(dec[i])) % 1_000_000_007
// 	}
// 	return hash
// }
