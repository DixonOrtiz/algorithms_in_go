package binary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchThreeWithBinary(t *testing.T) {
	index, number := SearchWithBinary([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3)
	assert.Equal(t, index, 2)
	assert.Equal(t, number, 3)
}

func TestSearchFiveWithBinary(t *testing.T) {
	index, number := SearchWithBinary([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5)
	assert.Equal(t, index, 4)
	assert.Equal(t, number, 5)
}

func TestSearchSevenWithBinary(t *testing.T) {
	index, number := SearchWithBinary([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7)
	assert.Equal(t, index, 6)
	assert.Equal(t, number, 7)
}

func TestSearchNonExistingNumberWithBinary(t *testing.T) {
	index, number := SearchWithBinary([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 14)
	assert.Equal(t, index, -1)
	assert.Equal(t, number, 0)
}
