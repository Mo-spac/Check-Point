package problems_2

func Checknumber(arg string) bool {
	for _, ch := range arg {
		if '0' <= ch && ch <= '9' {
			return false
		}
	}
	return true

}
