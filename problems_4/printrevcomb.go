package problems_4

import (
	"fmt"

	"github.com/01-edu/z01"
)

func Printrevcomb() {
	for i := 9; i >= 2; i-- {
		for j := i - 1; j >= 1; j-- {
			for k := j - 1; k >= 0; k-- {
				fmt.Printf("%d%d%d", i, j, k)
				if !(i == 2 && j == 1 && k == 0) {
					fmt.Print(", ")
				}
			}
		}
	}
	fmt.Println()
}

//********************************

func Printrevcomb1() {
	for i := 9; i >= 2; i-- {
		for j := i - 1; j >= 1; j-- {
			for k := j - 1; k >= 0; k-- {

				z01.PrintRune(rune(i + '0'))
				z01.PrintRune(rune(j + '0'))
				z01.PrintRune(rune(k + '0'))

				if !(i == 2 && j == 1 && k == 0) {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
