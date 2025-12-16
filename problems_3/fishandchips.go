package problems_3

func FishAndChips(n int) string {
	if n < 0 {
		return "Error: number in negative"
	}

	divby2 := n%2 == 0
	divby3 := n%3 == 0
	switch {
	case divby2 && divby3:
		return "Fish and Chips"
	case divby2:
		return "Fish"
	case divby3:
		return "Chips"
	default:
		return "error: non divisible"
	}
}

// func FishAndChips(n int) string {
// 	if n < 0 {
// 		return "error: number is negative"
// 	}

// 	switch {
// 	case n%6 == 0:
// 		return "Fish and Chips"
// 	case n%2 == 0:
// 		return "Fish"
// 	case n%3 == 0:
// 		return "Chips"
// 	default:
// 		return "error: non divisible"
// 	}
// }
