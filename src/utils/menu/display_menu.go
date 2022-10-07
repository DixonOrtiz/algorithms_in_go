package menu

import "fmt"

func DisplayMenu() {
	fmt.Print("Algorithms in Go \n\n")
	algorithms := []string{"bubble"}

	for key, algorithm := range algorithms {
		fmt.Printf("%d) %s\n", key+1, algorithm)
	}

	fmt.Print("\nSelect an algorithm by its number: ")
	var selected int
	fmt.Scan(&selected)

	executeFunctions(selected)
}
