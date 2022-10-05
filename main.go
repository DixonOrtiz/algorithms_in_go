package main

import (
	"algorithms_in_go/function/search"
	"fmt"
)

func main() {
	fmt.Print("Algorithms in Go \n\n")
	algorithms := []string{"bubble"}

	for key, algorithm := range algorithms {
		fmt.Printf("%d) %s\n", key+1, algorithm)
	}

	fmt.Print("\nSelect an algorithm by its name: ")
	var selected int
	fmt.Scan(&selected)

	executeFunctions(selected)
}

func executeFunctions(selected int) {
	switch selected {
	case 1:
		search.SorthWithBubble([]int{2, 1, 3, 0, 5, 4})

	default:
		fmt.Println("Invalid option entered")
	}
}
