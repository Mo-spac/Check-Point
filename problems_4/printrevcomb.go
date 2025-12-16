package problems_4

import "fmt"

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
