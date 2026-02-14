package problem_5

import "os"

func Run88wdmatch() {

	// إذا عدد المعاملات ليس 2 → لا شيء
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	i := 0 // مؤشر للسلسلة الأولى

	// نمر على السلسلة الثانية
	for _, c := range s2 {

		// إذا انتهينا من s1 نخرج
		if i == len(s1) {
			break
		}

		// إذا الحرف يساوي الحرف المطلوب من s1
		if c == rune(s1[i]) {
			i++
		}

	}

	// إذا وصلنا لنهاية s1 → نطبعها
	if i == len(s1) {
		os.Stdout.WriteString((s1))
		os.Stdout.WriteString(("\n"))
	}
}
