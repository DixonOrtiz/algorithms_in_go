package search

import "fmt"

func SorthWithBubble(numbers []int) []int {
	fmt.Print("\nActual numbers: ", numbers)
	fmt.Print("\nSorting with bubble...\n")

	for step := 0; step < len(numbers)-1; step++ {
		for i := 0; i < len(numbers)-step-1; i++ {
			if numbers[i] > numbers[i+1] {
				helper := numbers[i]
				numbers[i] = numbers[i+1]
				numbers[i+1] = helper
			}
		}
	}

	fmt.Println("Sorted numbers: ", numbers)
	return numbers
}
