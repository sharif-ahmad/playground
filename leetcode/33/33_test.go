package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run(`with nums = [4,5,6,7,0,1,2], target = 0`, func(t *testing.T) {
		assert.Equal(t, 4, search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	})

	t.Run(`with nums = [4,5,6,7,0,1,2], target = 5`, func(t *testing.T) {
		assert.Equal(t, 1, search([]int{4, 5, 6, 7, 0, 1, 2}, 5))
	})

	t.Run(`with nums = [4,5,6,7,0,1,2], target = 4`, func(t *testing.T) {
		assert.Equal(t, 0, search([]int{4, 5, 6, 7, 0, 1, 2}, 4))
	})

	t.Run(`with nums = [4,5,6,7,0,1,2], target = 2`, func(t *testing.T) {
		assert.Equal(t, 6, search([]int{4, 5, 6, 7, 0, 1, 2}, 2))
	})

	t.Run(`with nums = [4,5,6,7,0,1,2], target = 3`, func(t *testing.T) {
		assert.Equal(t, -1, search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	})

	t.Run(`with nums = [1], target = 0`, func(t *testing.T) {
		assert.Equal(t, -1, search([]int{1}, 0))
	})

	t.Run(`with nums = [4,5,6,7,8,1,2,3], target = 8`, func(t *testing.T) {
		assert.Equal(t, 4, search([]int{4, 5, 6, 7, 8, 1, 2, 3}, 8))
	})
}
