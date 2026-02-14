package problem_5

import "os"

func Run77inter() {
	// إذا لم يكن عدد المعاملات 2 (اسم البرنامج + معاملين) نخرج
	if len(os.Args) != 3 {
		return
	}

	// نخزن السلسلتين
	s1 := os.Args[1]
	s2 := os.Args[2]

	// map لمنع تكرار الحروف التي طبعناها
	seen := make(map[rune]bool)
	// map لمعرفة هل الحرف موجود في السلسلة الثانية
	inS2 := make(map[rune]bool)

	// نخزن كل حروف السلسلة الثانية داخل map
	for _, c := range s2 {
		inS2[c] = true
	}

	// نمر على السلسلة الأولى بالترتيب
	for _, c := range s1 {
		// إذا الحرف موجود في السلسلة الثانية
		// ولم نطبعه من قبل
		if inS2[c] && !seen[c] {
			// نطبع الحرف
			os.Stdout.WriteString(string(c))
			// نعلم أنه تم طباعته حتى لا نكرره
			seen[c] = true
		}
	}

	os.Stdout.WriteString("\n")

}
