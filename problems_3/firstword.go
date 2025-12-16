package problems_3

// func FierstWord(s string) string {
// 	for i := 0; i < len(s); i++ {
// 		if s[i] == ' ' {
// 			return s[:i]
// 		}
// 	}
// 	return s
// }

// func FirstWord(s string) string {
// 	for i, r := range s {
// 		if r == ' ' {
// 			return s[:i]
// 		}
// 	}
// 	return s
// }

func FirstWord(s string) string {
	word := " "
	started := false
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && !started {
			started = true
			word += string(s[i])
		} else if started {
			if s[i] == ' ' {
				break
			}
			word += string(s[i])
		}
	}
	return word[1:] + "\n"
}

// ******************

func FirstWord1(s string) string {
	start := -1

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && start == -1 {
			start = i
		} else if s[i] == ' ' && start != -1 {
			return s[start:i] + "\n"
		}
	}

	if start != -1 {
		return s[start:] + "\n"
	}

	return "\n"
}

// func FirstWord(s string) string {
// 	runes := []rune(s)
// 	start := -1
// 	for i, r := range runes {
// 		if r != ' ' && start == -1 {
// 			start = i
// 		} else if r == ' ' && start != -1 {
// 			return string(runes[start:i]) + "\n"
// 		}
// 	}
// 	if start != -1 {
// 		return string(runes[start:]) + "\n"
// 	}
// 	return "\n"
// }
