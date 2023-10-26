package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	t.Run(`with nums = [-2,1,-3,4,-1,2,1,-5,4]`, func(t *testing.T) {
		assert.Equal(t, 6, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	})

	t.Run(`with nums = [1]`, func(t *testing.T) {
		assert.Equal(t, 1, maxSubArray([]int{1}))
	})

	t.Run(`with nums = [5,4,-1,7,8]`, func(t *testing.T) {
		assert.Equal(t, 23, maxSubArray([]int{5, 4, -1, 7, 8}))
	})

	t.Run(`with nums = [-2,-1]`, func(t *testing.T) {
		assert.Equal(t, -1, maxSubArray([]int{-2, -1}))
	})
}
