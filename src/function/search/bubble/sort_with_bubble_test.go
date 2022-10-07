package bubble

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortWithBubble(t *testing.T) {
	ordered := SorthWithBubble([]int{2, 1, 3, 0, 5, 4})
	assert.Equal(t, ordered, []int{0, 1, 2, 3, 4, 5})
}
