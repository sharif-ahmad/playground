package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	t.Run(`with target = 7, nums = [2,3,1,2,4,3]`, func(t *testing.T) {
		assert.Equal(t, 2, minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	})

	t.Run(`with target = 4, nums = [1,4,4]`, func(t *testing.T) {
		assert.Equal(t, 1, minSubArrayLen(4, []int{1, 4, 4}))
	})

	t.Run(`with target = 11, nums = [1,1,1,1,1,1,1,1]`, func(t *testing.T) {
		assert.Equal(t, 0, minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
	})

	t.Run(`with target = 4, nums = [1,3,100,4]`, func(t *testing.T) {
		assert.Equal(t, 1, minSubArrayLen(4, []int{1, 3, 100, 4}))
	})
}
