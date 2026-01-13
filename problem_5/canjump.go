package problem_5

func CanJump(steps []uint) bool {

	n := len(steps)

	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}

	pos := 0

	for pos < n {
		if steps[pos] == 0 {
			return false
		}

		pos = int(uint(pos) + steps[pos])

		if pos == n-1 {
			return true
		} else if pos > n-1 {
			return false
		}
	}
	return false
}

// *******************

func CanJump1(input []uint) bool {
	// 1) لو الـ slice فاضية → false
	if len(input) == 0 {
		return false
	}

	// 2) لو فيها عنصر واحد → true
	if len(input) == 1 {
		return true
	}

	// 3) نبدأ من أول index
	index := 0
	lastIndex := len(input) - 1

	// 4) نكرر الحركة
	for {
		// لو وصلنا للآخر بالظبط
		if index == lastIndex {
			return true
		}

		// عدد الخطوات من المكان الحالي
		steps := input[index]

		// لو الخطوات = 0 ولسه موصلناش للآخر
		if steps == 0 {
			return false
		}

		// نتحرك للأمام
		index = index + int(steps)

		// لو خرجنا برّه حدود الـ slice
		if index > lastIndex {
			return false
		}
	}
}
