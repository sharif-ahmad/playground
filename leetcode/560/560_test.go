package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubarraySum(t *testing.T) {
	t.Run(`with nums = [1,1,1], k = 2`, func(t *testing.T) {
		assert.Equal(t, 2, subarraySum([]int{1, 1, 1}, 2))
	})

	t.Run(`with nums = [1,2,3], k = 3`, func(t *testing.T) {
		assert.Equal(t, 2, subarraySum([]int{1, 2, 3}, 3))
	})

	t.Run(`with nums = [-1,-1,1], k = 0`, func(t *testing.T) {
		assert.Equal(t, 1, subarraySum([]int{-1, -1, 1}, 0))
	})
}
