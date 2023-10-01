package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortColors(t *testing.T) {
	t.Run(`with nums = [2,0,2,1,1,0]`, func(t *testing.T) {
		nums := []int{2, 0, 2, 1, 1, 0}
		sortColors(nums)

		assert.Equal(t, []int{0, 0, 1, 1, 2, 2}, nums)
	})

	t.Run(`with nums = [2,0,1]`, func(t *testing.T) {
		nums := []int{2, 0, 1}
		sortColors(nums)

		assert.Equal(t, []int{0, 1, 2}, nums)
	})

	t.Run(`with nums = [0]`, func(t *testing.T) {
		nums := []int{0}
		sortColors(nums)

		assert.Equal(t, []int{0}, nums)
	})

	t.Run(`with nums = [2, 1]`, func(t *testing.T) {
		nums := []int{2, 1}
		sortColors(nums)

		assert.Equal(t, []int{1, 2}, nums)
	})

	t.Run(`with nums = [0, 1, 1, 0]`, func(t *testing.T) {
		nums := []int{0, 1, 1, 0}
		sortColors(nums)

		assert.Equal(t, []int{0, 0, 1, 1}, nums)
	})

	t.Run(`with nums = [1,2,0]`, func(t *testing.T) {
		nums := []int{1, 2, 0}
		sortColors(nums)

		assert.Equal(t, []int{0, 1, 2}, nums)
	})
}
