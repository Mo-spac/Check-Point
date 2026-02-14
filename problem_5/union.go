package problem_5

import "os"

func Run77union() {
	// إذا لم يوجد معاملين → نطبع فقط سطر جديد
	if len(os.Args) != 3 {
		os.Stdout.Write([]byte("\n"))
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	// map لمنع التكرار
	seen := make(map[rune]bool)

	// نمر على السلسلة الأولى
	for _, c := range s1 {
		if !seen[c] {
			os.Stdout.WriteString(string(c))
			seen[c] = true
		}
	}

	// نمر على السلسلة الثانية
	for _, c := range s2 {
		if !seen[c] {
			os.Stdout.WriteString(string(c))
			seen[c] = true
		}
	}

	os.Stdout.WriteString("\n")

}
