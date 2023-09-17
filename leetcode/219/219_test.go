package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	t.Run(`with nums = [1,2,3,1], k = 3`, func(t *testing.T) {
		assert.True(t, containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	})

	t.Run(`with nums = [1,0,1,1], k = 1`, func(t *testing.T) {
		assert.True(t, containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	})

	t.Run(`with nums = [1,2,3,1,2,3], k = 2`, func(t *testing.T) {
		assert.False(t, containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
	})

	t.Run(`with nums = [99, 99], k = 2`, func(t *testing.T) {
		assert.True(t, containsNearbyDuplicate([]int{99, 99}, 2))
	})
}
