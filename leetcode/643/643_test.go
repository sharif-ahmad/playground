package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	t.Run(`with nums = [1,12,-5,-6,50,3], k = 4`, func(t *testing.T) {
		assert.Equal(t, 12.75000, findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
	})

	t.Run(`with nums = [5], k = 1`, func(t *testing.T) {
		assert.Equal(t, 5.00000, findMaxAverage([]int{5}, 1))
	})

	t.Run(`with nums = [-1], k = 1`, func(t *testing.T) {
		assert.Equal(t, -1.00000, findMaxAverage([]int{-1}, 1))
	})
}
