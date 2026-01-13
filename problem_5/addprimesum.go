package problem_5

import (
	"os"
	"strconv"
)

func atoi(s string) (int, bool) {
	n := 0

	if len(s) == 0 {
		return 0, false
	}

	for _, digit := range s {
		// if digit < '0' || digit > '9' {
		if !('0' <= digit && digit <= '9') {
			return 0, false
		}

		n = n*10 + (int(digit) - '0')
	}

	return n, true
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}

	res := ""

	for n > 0 {
		digit := n % 10
		res = string(rune(digit+'0')) + res
		n /= 10
	}
	return res
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func Run2() {

	if len(os.Args) != 2 {
		os.Stdout.WriteString("0\n")
		return
	}

	num, ok := atoi(os.Args[1])
	if !ok || num < 0 {
		os.Stdout.WriteString("0\n")
		return
	}

	sum := 0
	for i := 2; i <= num; i++ {
		if isPrime(i) {
			sum += i
		}

	}
	os.Stdout.WriteString(itoa(sum) + "\n")
}

// ************* solution with strconv package *******************

func Run222() {

	if len(os.Args) != 2 {
		os.Stdout.WriteString("0\n")
		return
	}

	num, err := strconv.Atoi(os.Args[1])

	if err != nil || num < 0 {
		os.Stdout.WriteString("0\n")
		return
	}

	sum := 0
	for i := 2; i <= num; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	os.Stdout.WriteString(strconv.Itoa(sum) + "\n")
}
