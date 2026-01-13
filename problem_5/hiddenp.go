package problem_5

import (
	"fmt"
	"os"
)

func Run6Hidden() {

	if len(os.Args) != 3 {
		return
	}

	// s1 := os.Args[1]
	// s2 := os.Args[2]

	s1 := ""
	s2 := ""

	if len(os.Args) > 1 {
		s1 = os.Args[1]
	}
	if len(os.Args) > 2 {
		s2 = os.Args[2]
	}

	if s1 == "" {
		os.Stdout.WriteString("1\n")
		return
	}

	i := 0
	for _, c := range s2 {
		if c == rune(s1[i]) {
			i++
			if i == len(s1) {
				os.Stdout.WriteString("1\n")
				return
			}
		}
	}

	os.Stdout.WriteString("0\n")
}

// *******************
func Run66Hidden() {

	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	// s1 := ""
	// s2 := ""

	// if len(os.Args) > 1 {
	// 	s1 = os.Args[1]
	// }
	// if len(os.Args) > 2 {
	// 	s2 = os.Args[2]
	// }

	if s1 == "" {
		fmt.Println("1")
		return
	}

	i := 0

	for _, c := range s2 {
		if c == rune(s1[i]) {
			i++
			if i == len(s1) {
				fmt.Println("1")
				return
			}
		}
	}
	fmt.Println("0")
}
