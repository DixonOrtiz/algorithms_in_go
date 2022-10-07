package bubble

//
// 	This algorithm is usually used to teach logic and algorithms.
// 	It's order is O(n^2), that makes it an inefficient algorithm
// 	compared to others of the same kind.
//
func SorthWithBubble(numbers []int) []int {
	printTitle(numbers)

	// The numbers of steps in this algorithm
	// is equal to the length of the array.
	length := len(numbers)
	for step := 0; step < length-1; step++ {

		// Finds the largest value and moves it to the last position.
		// Next iterations don't consider ordered elements.
		for i := 0; i < length-step-1; i++ {
			if numbers[i] > numbers[i+1] {
				helper := numbers[i]
				numbers[i] = numbers[i+1]
				numbers[i+1] = helper
			}
		}
	}

	return numbers
}
