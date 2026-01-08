package problems_4

import (
	"fmt"

	"github.com/01-edu/z01"

	"strconv"
)

func ZipString(s string) string {

	n := len(s)

	if n == 0 {
		return ""
	}
	result := ""

	count := 1

	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			count++

		} else {
			result += fmt.Sprintf("%d%c", count, s[i-1])
			count = 1
		}
	}
	result += fmt.Sprintf("%d%c", count, s[n-1])

	return result
}

// ********************************

func ZipString1(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	result := ""

	count := 1
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			result += IntTresultoString(count)
			result += string(s[i-1])
			count = 1
		}

	}
	result += IntTresultoString(count)
	result += string(s[n-1])
	return result
}

func IntTresultoString(n int) string {
	if n == 0 {
		return "0"
	}

	str := ""
	for n > 0 {
		digit := n % 10
		str = string(rune('0'+digit)) + str
		n = n / 10
	}
	return str
}

// ********************************
// باستخدام مكتبة z01
// import "github.com/01-edu/z01"

func ZipString2(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	count := 1

	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			PrintNumper(count)
			z01.PrintRune(rune(s[i-1]))
			count = 1
		}
	}
	PrintNumper(count)
	z01.PrintRune(rune(s[n-1]))
	return ""
}

func PrintNumper(n int) {
	if n >= 10 {
		PrintNumper(n / 10)
	}
	z01.PrintRune(rune('0' + n%10))
}

// *********** In Check Point **********

func ZipString3(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	result := ""
	count := 1

	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			result += strconv.Itoa(count)
			result += string(s[i-1])
			count = 1
		}
	}

	// آخر مجموعة
	result += strconv.Itoa(count)
	result += string(s[n-1])

	return result
}
