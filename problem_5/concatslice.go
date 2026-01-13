package problem_5

func ConcatSlice(slice1, slice2 []int) []int {

	var result []int

	result = append(result, slice1...)
	result = append(result, slice2...)

	return result
}

func ConcatSlice1(slice1, slice2 []int) []int {
	result := make([]int, 0, len(slice1)+len(slice2))

	result = append(result, slice1...)
	result = append(result, slice2...)

	return result
}

func ConcatSlice2(slice1, slice2 []int) []int {
	totalLen := len(slice1) + len(slice2)
	result := make([]int, totalLen)

	index := 0

	for _, v := range slice1 {
		result[index] = v
		index++
	}

	for _, v := range slice2 {
		result[index] = v
		index++
	}

	return result

}

func ConcatSlice3(slice1, slice2 []int) []int {
	result := []int{}

	for _, v := range slice1 {
		result = append(result, v)
	}

	for _, v := range slice2 {
		result = append(result, v)
	}

	return result
}

func ConcatSlice4(slice1, slice2 []int) []int {
	result := make([]int, 0, len(slice1)+len(slice2))

	for i := 0; i < len(slice1); i++ {
		result[i] = slice1[i]
	}

	for i := 0; i < len(slice2); i++ {
		result[len(slice1)+i] = slice2[i]
	}

	return result
}
