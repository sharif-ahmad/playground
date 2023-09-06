package counting_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindKthLargest(t *testing.T) {

	assert.Equal(t, 5, findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	assert.Equal(t, 4, findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	assert.Equal(t, 2, findKthLargest([]int{1, 1, 1, 2, 3, 4, 4}, 4))
	assert.Equal(t, 6, findKthLargest([]int{7, 6, 5, 4, 3, 2, 1}, 2))
	assert.Equal(t, 7, findKthLargest([]int{7, 7}, 2))
	assert.Equal(t, 5, findKthLargest([]int{5}, 1))
	assert.Equal(t, 3, findKthLargest([]int{7, 6, 5, 4, 3, 2, 1}, 5))
}
