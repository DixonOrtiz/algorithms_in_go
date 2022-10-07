package bubble

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortUnorderedWithBubble(t *testing.T) {
	ordered := SorthWithBubble([]int{2, 1, 3})
	assert.Equal(t, ordered, []int{1, 2, 3})
}

func TestSortOrderedWithBubble(t *testing.T) {
	ordered := SorthWithBubble([]int{0, 1, 2, 3})
	assert.Equal(t, ordered, []int{0, 1, 2, 3})
}
