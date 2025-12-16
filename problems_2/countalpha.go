package problems_2

func Countalpha(s string) int {

	count := 0
	for _, r := range s {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
			count++
		}
	}
	return count
}
