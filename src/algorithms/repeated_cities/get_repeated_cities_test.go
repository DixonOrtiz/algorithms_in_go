package repeated_cities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNegativeRepeatingCities(t *testing.T) {
	cities, err := GetRepeatedCities(-3, nil)

	assert.Nil(t, cities)
	assert.EqualError(t, err, "a list of repeated cities cannot be zero or negative")
}

func TestGetZeroRepeatedCity(t *testing.T) {
	cities, err := GetRepeatedCities(0, nil)

	assert.Nil(t, cities)
	assert.EqualError(t, err, "a list of repeated cities cannot be zero or negative")
}

func TestNumberIsGreatestThanCitiesLenght(t *testing.T) {
	cities, err := GetRepeatedCities(2, []string{"santiago"})
	assert.Nil(t, cities)
	assert.EqualError(t, err, "number of most repeated cities cannot be greatest than cities lenght")
}

func TestGetMostRepeatedCity(t *testing.T) {
	expectedCity := City{Name: "santiago", Times: 2}

	cities, err := GetRepeatedCities(1, []string{"santiago", "valdivia", "santiago"})

	assert.Equal(t, 1, len(cities))
	assert.Equal(t, expectedCity, cities[0])
	assert.Nil(t, err)
}

func TestGetTwoMostRepeatedCities(t *testing.T) {
	expectedFirst := City{Name: "santiago", Times: 2}
	expectedSec := City{Name: "valdivia", Times: 1}

	cities, err := GetRepeatedCities(2, []string{"santiago", "valdivia", "santiago"})
	firstCity, secondCity := cities[0], cities[1]

	assert.Equal(t, 2, len(cities))
	assert.Equal(t, expectedFirst, firstCity)
	assert.Equal(t, expectedSec, secondCity)
	assert.Nil(t, err)
}

func TestGetThreeMostRepeatedCities(t *testing.T) {
	expectedFirst := City{Name: "santiago", Times: 2}
	expectedSec := City{Name: "valparaiso", Times: 2}
	expectedThird := City{Name: "arica", Times: 1}

	cities, err := GetRepeatedCities(3, []string{"santiago", "valparaiso", "santiago", "valparaiso", "arica"})
	firstCity, secondCity, thirdCity := cities[0], cities[1], cities[2]

	assert.Equal(t, 3, len(cities))
	assert.Equal(t, expectedFirst, firstCity)
	assert.Equal(t, expectedSec, secondCity)
	assert.Equal(t, expectedThird, thirdCity)
	assert.Nil(t, err)
}
