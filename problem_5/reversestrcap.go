package problem_5

import "os"

func isLetter(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
}

func Run1() {
	args := os.Args[1:]

	for _, s := range args {

		for i := 0; i < len(s); i++ {
			c := s[i]
			if 'A' <= c && c <= 'Z' {
				c += 32
			}

			// if 'a' <= c && c <= 'z' && (i+1 == len(s) || s[i+1] == ' ') {
			if 'a' <= c && c <= 'z' && (i+1 == len(s) || !isLetter(s[i+1])) {
				c -= 32
			}
			os.Stdout.WriteString(string(c))
		}
		os.Stdout.WriteString("\n")
	}
}
