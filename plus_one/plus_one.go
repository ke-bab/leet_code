package plus_one

func plusOne(digits []int) []int {
	if len(digits) < 1 {
		return digits
	}
	incrRecursive(len(digits)-1, &digits)

	return digits
}

func incrRecursive(index int, digits *[]int) {
	(*digits)[index]++
	if (*digits)[index] > 9 {
		(*digits)[index] = 0
		if hasLeftDigit(index) {
			incrRecursive(index-1, digits)
		} else {
			createLeftAndSetToOne(digits)
		}
	}
}

func hasLeftDigit(index int) bool {
	return index-1 >= 0
}

func createLeftAndSetToOne(digits *[]int) {
	newSlice := make([]int, len(*digits)+1)
	newSlice[0] = 1

	*digits = newSlice
}
