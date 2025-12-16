package problems_3

func LastWord(s string) string {

	end := len(s) - 1

	// تجاهل المسافات في النهاية
	for end >= 0 && s[end] == ' ' {
		end--
	}

	// لو مفيش أي كلمة
	if end == -1 {
		return "\n"
	}

	// إيجاد بداية آخر كلمة
	start := end

	for start >= 0 && s[start] != ' ' {
		start--
	}

	return s[start+1:end+1] + "\n"

}

// func LastWord(s string) string {
// 	end := len(s) - 1

// 	for end >= 0 && s[end] == ' ' {
// 		end--
// 	}

// 	if end < 0 {
// 		return ""
// 	}

// 	start := end
// 	for start >= 0 && s[start] != ' ' {
// 		start--
// 	}
// 	return s[start+1 : end+1]
// }

func LastWord1(s string) string {

	end := -1

	// تجاهل المسافات في النهاية
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			end = i
			break
		}
	}

	// لو مفيش أي كلمة
	if end == -1 {
		return "\n"
	}

	// إيجاد بداية آخر كلمة
	start := end
	for start >= 0 && s[start] != ' ' {
		start--
	}

	return s[start+1:end+1] + "\n"

}
