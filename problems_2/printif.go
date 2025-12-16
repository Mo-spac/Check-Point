package problems_2

func Printif(str string) string {
	if len(str) == 0 || len(str) >= 3 {
		return "G\n"
	} else {
		return "Invalid Input\n"
	}
}
