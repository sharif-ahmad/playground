package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchRange(t *testing.T) {
	t.Run(`when nums = [5,7,7,8,8,10], target = 8`, func(t *testing.T) {
		assert.Equal(t, []int{3, 4}, searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	})

	t.Run(`when nums = [5,7,7,8,8,10], target = 6`, func(t *testing.T) {
		assert.Equal(t, []int{-1, -1}, searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	})

	t.Run(`when nums = [], target = 0`, func(t *testing.T) {
		assert.Equal(t, []int{-1, -1}, searchRange([]int{}, 0))
	})

	t.Run(`when nums = [1,1,3,4,8], target = 8`, func(t *testing.T) {
		assert.Equal(t, []int{4, 4}, searchRange([]int{1, 1, 3, 4, 8}, 8))
	})
}

func TestLowerBound(t *testing.T) {
	t.Run(`when all elements are less than the target`, func(t *testing.T) {
		assert.Equal(t, 3, lowerBound([]int{2, 3, 4}, 6))
	})

	t.Run(`when all elements are greater than the target`, func(t *testing.T) {
		assert.Equal(t, 0, lowerBound([]int{7, 9, 13}, 6))
	})

	t.Run(`when the target is present in the array`, func(t *testing.T) {
		assert.Equal(t, 1, lowerBound([]int{5, 6, 7, 9, 13}, 6))
	})

	t.Run(`when the target is present multiple times in the array`, func(t *testing.T) {
		assert.Equal(t, 1, lowerBound([]int{5, 6, 6, 6, 7, 9, 13}, 6))
	})
}

func TestUpperBound(t *testing.T) {
	t.Run(`when all elements are less than the target`, func(t *testing.T) {
		assert.Equal(t, 3, upperBound([]int{2, 3, 4}, 6))
	})

	t.Run(`when all elements are greater than the target`, func(t *testing.T) {
		assert.Equal(t, 0, upperBound([]int{7, 9, 13}, 6))
	})

	t.Run(`when the target is present in the array`, func(t *testing.T) {
		assert.Equal(t, 2, upperBound([]int{5, 6, 7, 9, 13}, 6))
	})

	t.Run(`when the target is present multiple times in the array`, func(t *testing.T) {
		assert.Equal(t, 4, upperBound([]int{5, 6, 6, 6, 7, 9, 13}, 6))
	})
}
