package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxArea(t *testing.T) {
	t.Run(`with height = [1,8,6,2,5,4,8,3,7]`, func(t *testing.T) {
		assert.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	})

	t.Run(`with height = [1,1]`, func(t *testing.T) {
		assert.Equal(t, 1, maxArea([]int{1, 1}))
	})
}
