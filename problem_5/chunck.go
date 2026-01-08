package problem_5

import "fmt"

func Chunk(slice []int, size int) {
	n := len(slice)
	if size == 0 {
		fmt.Println()
		return
	}

	if n == 0 {
		fmt.Println([][]int{})
		return
	}

	var result [][]int
	for i := 0; i < n; i += size {
		end := i + size
		if end > n {
			end = n
		}
		result = append(result, slice[i:end])
	}
	fmt.Println(result)
}
