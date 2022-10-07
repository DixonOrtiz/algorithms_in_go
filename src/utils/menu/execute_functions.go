package menu

import (
	"algorithms_in_go/src/function/search/bubble"
	"fmt"
)

func executeFunctions(selected int) {
	switch selected {
	case 1:
		fmt.Println("Sorted numbers: ", bubble.SorthWithBubble([]int{2, 1, 3, 0, 5, 4}))

	default:
		fmt.Println("Invalid option entered")
	}
	fmt.Println()
}
