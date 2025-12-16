package problems_3

func CamelToSnakeCase(s string) string {
	runes := []rune(s)
	n := len(runes)
	if s == "" {
		return ""
	}

	// if n == 1 {
	// 	return s
	// }

	for i, char := range runes {
		if !('A' <= char && char <= 'Z') && !('a' <= char && char <= 'z') {
			return s
		}

		if i > 0 && ('A' <= char && char <= 'Z') && ('A' <= runes[i-1] && runes[i-1] <= 'Z') {
			return s
		}
	}

	if 'A' <= runes[n-1] && runes[n-1] <= 'Z' {
		return s
	}

	var result []rune

	for i, ch := range runes {
		if 'A' <= ch && ch <= 'Z' {
			if i > 0 {
				result = append(result, '_')
			}
		}
		result = append(result, ch)
	}
	return string(result)
}

// ********************************

func CamelToSnakeCase1(s string) string {
	runes := []rune(s)

	if s == "" {
		return ""
	}
	if !IsCamel(s) {
		return s
	}
	var result []rune
	for i, char := range runes {
		if 'A' <= char && char <= 'Z' {
			if i > 0 {
				result = append(result, '_')
			}
		}
		result = append(result, char)
	}
	return string(result)
}

func IsCamel(s string) bool {

	runes := []rune(s)
	n := len(runes)

	if n == 1 {
		return true
	}

	last := runes[n-1]
	// last cap
	if 'A' <= last && last <= 'Z' {
		return false
	}

	// 2 cap
	for i, ch := range runes {

		// 1 2 3 & & * #
		if !('a' <= ch && ch <= 'z') && !('A' <= ch && ch <= 'Z') {
			return false
		}

		if i > 0 && ('A' <= ch && ch <= 'Z') && ('A' <= runes[i-1] && runes[i-1] <= 'Z') {
			return false
		}

	}
	return true
}

// ****************************

func CamelToSnakeCase2(s string) string {
	if s == "" {
		return ""
	}

	if !isCamelCase(s) {
		return s
	}

	result := ""
	for i, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			if i != 0 {
				result += "_"
			}
		}
		result += string(ch)
	}

	return result
}

func isCamelCase(s string) bool {

	for i, ch := range s {
		if !(ch >= 'a' && ch <= 'z') &&
			!(ch >= 'A' && ch <= 'Z') {
			return false
		}
		if i > 0 && s[i] >= 'A' && s[i] <= 'Z' &&
			s[i-1] >= 'A' && s[i-1] <= 'Z' {
			return false
		}
	}

	last := s[len(s)-1]
	if last >= 'A' && last <= 'Z' {
		return false
	}

	// for i := 0; i < len(s)-1; i++ {
	// 	if s[i] >= 'A' && s[i] <= 'Z' &&
	// 		s[i+1] >= 'A' && s[i+1] <= 'Z' {
	// 		return false
	// 	}
	// }

	return true
}

// ***************************

// دوال مساعدة للتعامل مع الأحرف باستخدام قيم ASCII فقط
func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isLetter(r rune) bool {
	return isLower(r) || isUpper(r)
}

// ------------------------------------------------------------------

// IsCamel تتحقق من شروط كتابة CamelCase:
// 1. لا ينتهي بحرف كبير.
// 2. لا يتبع حرفان كبيران بعضهما مباشرة (باستثناء بداية الكلمة أو إذا كانت الكلمة كلها كبيرة).
// 3. لا يحتوي على أرقام أو علامات ترقيم.
func IsCamel3(s string) bool {
	runes := []rune(s)
	n := len(runes)

	if n == 0 {
		return true // يُعتبر سلسلة فارغة صالحة ولكن يتم  إعادتها لاحقاً
	}

	// 1. لا تنتهي الكلمة بحرف كبير (CamelCasE).
	if isUpper(runes[n-1]) {
		return false
	}

	for i, ch := range runes {
		// 3. الأرقام أو علامات الترقيم غير مسموح بها.
		if !isLetter(ch) {
			return false
		}

		// 2. لا يتبع حرفان كبيران بعضهما مباشرة (CamelCAse).
		// هذا التحقق يجب أن يسمح بـ "HTTPMethod" لكن يحظر "CamelCAse".
		// الطريقة الأبسط للالتزام بقاعدة التمرين: إذا وجدنا U-U (مثل 'CA' في 'CamelCAse')، نعود بـ false
		// ما لم يكن الحرف الذي قبلهم (إذا وجد) حرفاً كبيراً أيضاً (للسماح بالثلاثة أو أكثر مثل 'HTTM' في 'HTTPMethod').
		// لكن بالنظر للقاعدة الصارمة "CamelCAse" غير مسموح، نكتفي بالتحقق من حرفين متتاليين.
		if i > 0 && isUpper(ch) && isUpper(runes[i-1]) {
			// يجب السماح بالكلمات الكبيرة بالكامل، مثل "UPPER"
			// للتمييز بين 'CAMELtoSnackCASE' و 'UPPER'.
			// بما أننا نعود بـ `false` هنا، فإن 'CAMELtoSnackCASE' ستعود بـ `false` (غير صالحة)، وهذا صحيح حسب المثال المطلوب.

			// لتمكين معالجة "HTTPMethod" بشكل صحيح (إذا كان هذا هو القصد):
			// يجب أن يكون الحرف السابق صغيراً ليعتبر تتابع U-U غير صالح.
			// مثال: 'l' -> 'C' -> 'A' (هنا تكون العودة بـ false)
			// لكن إذا كان 'T' -> 'T' في 'HTTP' ، فالحرف السابق هو 'T' وهو كبير، لذا سنعود بـ false.
			// لذلك، سنلتزم بالتحقق البسيط الذي يجعل 'CAMELtoSnackCASE' غير صالحة، بناءً على المخرجات المطلوبة.
			return false
		}
	}
	return true
}

// ------------------------------------------------------------------

func CamelToSnakeCase3(s string) string {
	runes := []rune(s)

	// 1. إذا كانت السلسلة فارغة، أعد سلسلة فارغة.
	if s == "" {
		return ""
	}

	// 3. إذا لم تكن السلسلة camelCase (باستخدام القواعد المحددة)، أعد السلسلة كما هي.
	if !IsCamel3(s) {
		return s
	}

	var result []rune
	for i, char := range runes {
		// إذا كان الحرف كبيراً
		if isUpper(char) {
			// ندرج '_' إذا لم يكن الحرف هو الأول في السلسلة
			if i > 0 {
				result = append(result, '_')
			}
		}
		// نُضيف الحرف الأصلي. لا نحول الحرف إلى صغير ليتوافق مع المخرج المطلوب: "Hello_World"
		result = append(result, char)
	}

	// 2. إذا كانت السلسلة camelCase، أعد النسخة snake_case.
	return string(result)
}

// **********************************

// دوال مساعدة للتعامل مع الأحرف باستخدام قيم ASCII (البايتات) فقط
// يتم التعامل مع الحرف كقيمة uint8 (بايت)
func isLower4(b byte) bool {
	return b >= 'a' && b <= 'z'
}

func isUpper4(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func isLetter4(b byte) bool {
	return isLower4(b) || isUpper4(b)
}

// ------------------------------------------------------------------

// IsCamel تتحقق من شروط كتابة CamelCase باستخدام البايتات.
func IsCamel4(s string) bool {
	n := len(s)

	if n == 0 {
		return true
	}

	// 1. لا تنتهي الكلمة بحرف كبير (CamelCasE).
	// نستخدم s[n-1] للوصول إلى البايت الأخير.
	if isUpper4(s[n-1]) {
		return false
	}

	for i := 0; i < n; i++ {
		ch := s[i] // ch هو بايت

		// 3. الأرقام أو علامات الترقيم غير مسموح بها.
		if !isLetter4(ch) {
			return false
		}

		// 2. لا يتبع حرفان كبيران بعضهما مباشرة (CamelCAse).
		if i > 0 && isUpper4(ch) && isUpper4(s[i-1]) {
			return false
		}
	}
	return true
}

// ------------------------------------------------------------------

func CamelToSnakeCase4(s string) string {
	// 1. إذا كانت السلسلة فارغة، أعد سلسلة فارغة.
	if s == "" {
		return ""
	}

	// 3. إذا لم تكن السلسلة camelCase، أعد السلسلة كما هي.
	if !IsCamel4(s) {
		return s
	}

	// استخدام strings.Builder لتحسين الأداء بدلاً من concatenation
	// بما أن strings محظور، سنستخدم شريحة من البايتات []byte
	var result []byte

	for i := 0; i < len(s); i++ {
		charByte := s[i]

		// إذا كان الحرف كبيراً
		if isUpper4(charByte) {
			// ندرج '_' إذا لم يكن الحرف هو الأول في السلسلة
			if i > 0 {
				result = append(result, '_')
			}
		}
		// نُضيف الحرف الأصلي
		result = append(result, charByte)
	}

	// 2. إذا كانت السلسلة camelCase، أعد النسخة snake_case.
	return string(result)
}
