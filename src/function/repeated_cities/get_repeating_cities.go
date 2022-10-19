package repeated_cities

import (
	"errors"
)

//
// Obtain a dynamic number of most repeated
// cities from a cities slice
//
func GetRepeatedCities(number int, cities []string) ([]City, error) {
	var sortedCities []City
	register := make(map[string]int)

	if number < 1 {
		return nil, errors.New("a list of repeated cities cannot be zero or negative")
	}

	if number > len(cities) {
		return nil, errors.New("number of most repeated cities cannot me greatest than cities lenght")
	}

	for _, city := range cities {
		if _, exists := register[city]; exists {
			register[city] += 1
			continue
		}
		register[city] = 1
	}

	for number > 0 {
		city := getMostRepeatedCity(register)
		sortedCities = append(sortedCities, city)
		delete(register, city.Name)
		number--
	}

	return sortedCities, nil
}

//
// Obtain most repeated city from city register
//
func getMostRepeatedCity(register map[string]int) City {
	var repeated City

	for city, times := range register {
		if times > repeated.Times {
			repeated.Name = city
			repeated.Times = times
		}
	}

	return repeated
}
