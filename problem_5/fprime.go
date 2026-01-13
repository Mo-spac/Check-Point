package problem_5

import (
	"os"
	"strconv"
)

func Run7Fprime() {
	if len(os.Args) != 2 {
		return
	}

	// تحويل string → int باستخدام strconv
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 1 {
		return
	}

	first := true

	// تحليل العوامل الأولية
	for d := 2; d <= n; d++ {
		for n%d == 0 {
			if !first {
				os.Stdout.WriteString("*")
			}
			os.Stdout.WriteString(strconv.Itoa(d))
			first = false
			n /= d
		}
		if n == 1 {
			break // توقف مبكر إذا تم تحليل كل العوامل
		}
	}

	os.Stdout.WriteString("\n")
}
