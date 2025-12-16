package problems_4

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	map1 := make(map[rune]bool)
	map2 := make(map[rune]bool)

	for _, ch := range str1 {
		map1[ch] = true
	}

	for _, ch := range str2 {
		map2[ch] = true
	}

	count := 0

	for ch := range map1 {
		if !map2[ch] {
			count++
		}
	}

	for ch := range map2 {
		if !map1[ch] {
			count++
		}
	}
	return count
}

//*************

func WeAreUnique1(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	var charInStr1 [256]bool
	var charInStr2 [256]bool

	for i := 0; i < len(str1); i++ {
		charInStr1[str1[i]] = true
	}

	for i := 0; i < len(str2); i++ {
		charInStr2[str2[i]] = true
	}

	count := 0

	for i := 0; i < 256; i++ {
		if charInStr1[i] && !charInStr2[i] {
			count++
		} else if charInStr2[i] && !charInStr1[i] {
			count++
		}
	}
	return count
}
