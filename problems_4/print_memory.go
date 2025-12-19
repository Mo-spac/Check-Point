package problems_4

import (
	"fmt"

	"github.com/01-edu/z01"
)

func PrintMemory(arr [10]byte) {

	hex := "0123456789abcdef"

	n := len(arr)

	for i := 0; i < n; i += 4 {
		end := i + 4
		if end > n {
			end = n
		}

		for j := i; j < end; j++ {
			z01.PrintRune(rune(hex[arr[j]/16]))
			z01.PrintRune(rune(hex[arr[j]%16]))

			if j != end-1 {
				z01.PrintRune(' ')
				// fmt.Printf(" ")
			}
		}

		z01.PrintRune('\n')
		// fmt.Print("\n")
	}

	for i := 0; i < n; i++ {
		if 32 <= arr[i] && arr[i] <= 126 {
			z01.PrintRune(rune(arr[i]))
			// fmt.Printf("%c", arr[i])
		} else {
			z01.PrintRune('.')
			// fmt.Print(".")
		}
	}

	z01.PrintRune('\n')
	// fmt.Print("\n")
}

// ********************************

func PrintMemory1(arr [10]byte) {

	n := len(arr)

	for i := 0; i < n; i += 4 {
		end := i + 4
		if end > n {
			end = n
		}
		for j := i; j < end; j++ {
			fmt.Printf("%02x", arr[j])
			if j != end-1 {
				fmt.Printf(" ")
			}
		}
		fmt.Print("\n")
	}

	for i := 0; i < n; i++ {
		if 32 <= arr[i] && arr[i] <= 126 {
			fmt.Printf("%c", arr[i])
		} else {
			fmt.Print(".")
		}
	}

	fmt.Print("\n")
}
