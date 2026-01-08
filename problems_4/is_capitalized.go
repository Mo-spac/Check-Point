package problems_4

func IsCapitalized(s string) bool {

	if s == "" {
		return false
	}

	if 'a' <= s[0] && s[0] <= 'z' {
		return false
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] == ' ' {
			if 'a' <= s[i+1] && s[i+1] <= 'z' {
				return false
			}
		}

	}
	return true
}

// ********************

func IsCapitalized1(s string) bool {
	if s == "" {
		return false
	}

	for i := 0; i < len(s); i++ {
		// أول حرف في النص أو بعد مسافة = بداية كلمة
		// if (i == 0 || s[i-1] == ' ') && s[i] != ' ' {
		if i == 0 || s[i-1] == ' ' {

			// لو حرف صغير
			if 'a' <= s[i] && s[i] <= 'z' {
				return false
			}
		}
	}

	return true

}
