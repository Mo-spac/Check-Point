package problem_5

// import "strings"

func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg
	}

	// var result strings.Builder
	result := ""

	save := true

	i := 0

	for i < len(arg) {
		// end := min(i + num, len(arg))
		end := i + num
		if end > len(arg) {
			end = len(arg)
		}

		if save {
			// result.WriteString(arg[i:end])
			result += arg[i:end]
		}

		save = !save
		i += num

	}

	return result
}
