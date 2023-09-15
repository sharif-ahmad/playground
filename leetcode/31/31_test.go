package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	t.Run(`with nums = [1,2,3]`, func(t *testing.T) {
		nums := []int{1, 2, 3}
		nextPermutation(nums)

		assert.Equal(t, []int{1, 3, 2}, nums)
	})

	t.Run(`with nums = [3,2,1]`, func(t *testing.T) {
		nums := []int{3, 2, 1}
		nextPermutation(nums)

		assert.Equal(t, []int{1, 2, 3}, nums)
	})

	t.Run(`with nums = [1, 1, 5]`, func(t *testing.T) {
		nums := []int{1, 1, 5}
		nextPermutation(nums)

		assert.Equal(t, []int{1, 5, 1}, nums)
	})

	t.Run(`with nums = [1, 3, 2]`, func(t *testing.T) {
		nums := []int{1, 3, 2}
		nextPermutation(nums)

		assert.Equal(t, []int{2, 1, 3}, nums)
	})

	t.Run(`with nums = [1]`, func(t *testing.T) {
		nums := []int{1}
		nextPermutation(nums)

		assert.Equal(t, []int{1}, nums)
	})
}

func TestReverseUpperBound(t *testing.T) {
	t.Run(`with nums = [8, 5, 5, 4, 3, 2, 1], t = 4`, func(t *testing.T) {
		assert.Equal(t, 2, reverseUpperBound([]int{8, 5, 5, 4, 3, 2, 1}, 4))
	})

	t.Run(`with nums = [8, 5, 5, 4, 3, 2, 1], t = 5`, func(t *testing.T) {
		assert.Equal(t, 0, reverseUpperBound([]int{8, 5, 5, 4, 3, 2, 1}, 5))
	})

	t.Run(`with nums = [8, 5, 5, 4, 3, 2, 1], t = 8`, func(t *testing.T) {
		assert.Equal(t, -1, reverseUpperBound([]int{8, 5, 5, 4, 3, 2, 1}, 8))
	})

	t.Run(`with nums = [8, 5, 5, 4, 3, 2, 1], t = 0`, func(t *testing.T) {
		assert.Equal(t, 6, reverseUpperBound([]int{8, 5, 5, 4, 3, 2, 1}, 0))
	})

	t.Run(`with nums = [1], t = 0`, func(t *testing.T) {
		assert.Equal(t, 0, reverseUpperBound([]int{1}, 0))
	})

	t.Run(`with nums = [1], t = 1`, func(t *testing.T) {
		assert.Equal(t, -1, reverseUpperBound([]int{1}, 1))
	})
}
