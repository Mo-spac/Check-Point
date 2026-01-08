package problem_5

import "os"

func Run55() {
	args := os.Args[1:]

	for _, s := range args {
		for i := 1; i < len(s); i++ {
			for j := 0; j < i; j++ {
				if s[j] == s[i] {
					continue
				}
			}
			if s[i] == s[i-1] {
				continue
			}
			os.Stdout.WriteString(string(s[i-1]))
		}

	}

}

func DoNotRepeat(s1 string, s2 string) string {
	if len(s1) == 0 && len(s2) == 0 {
		return ""
	}

	result := ""
	found := false

	for i := 1; i < len(s1); i++ {

		if s1[i] == s1[i-1] {
			found = true
			break
		}
		found = false

	}
	if !found {
		result += string(s1)
	}

	for i := 1; i < len(s2); i++ {

		if s2[i] == s2[i-1] {
			found = true
			break
		}
		found = false

	}
	if !found {
		result += string(s2)
	}

	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(result); j++ {

			if s1[i] == result[j] {
				found = true
				break
			}
			found = false
		}

	}
	if !found {
		result += string(s1)
	}

	for i := 0; i < len(s2); i++ {
		for j := 0; j < len(result); j++ {

			if s2[i] == result[j] {
				found = true
				break
			}
			found = false
		}
	}
	if !found {
		result += string(s2)
	}

	return result
}
