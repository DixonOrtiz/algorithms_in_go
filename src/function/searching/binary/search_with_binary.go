package binary

//
// Binary search is an efficient searching algorithm.
// It's order is 0(log(n)), but the searching array must
// be ordered before searching.
//
func SearchWitchBinary(numbers []int, number int) (int, int) {
	left, right := 0, len(numbers)-1

	for left != right {
		center := (left + right) / 2

		// Number is found
		if numbers[center] == number {
			return center, number
		}

		// Pivot to the right
		if numbers[center] < number {
			left++
		}

		// Pivot to the left
		if numbers[center] > number {
			right--
		}
	}

	return -1, 0
}
