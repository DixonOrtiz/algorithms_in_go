package main

import (
	"algorithms_in_go/src/algorithms/repeated_cities"
	"fmt"
)

func main() {
	// ¡Welcome to my project!
	// Please take a look to the differente algorithms included in this repository.
	// You can check the functionality of each algorithm in the files with name ended in '_test.go'
	fmt.Println(repeated_cities.GetRepeatedCities(6, repeated_cities.Cities))
}
