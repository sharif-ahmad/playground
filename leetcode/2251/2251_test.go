package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullBloomFlowers(t *testing.T) {
	t.Run(`with flowers = [[1,6],[3,7],[9,12],[4,13]], people = [2,3,7,11]`, func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 2, 2}, fullBloomFlowers([][]int{{1, 6}, {3, 7}, {9, 12}, {4, 13}}, []int{2, 3, 7, 11}))
	})

	t.Run(`with flowers = [[1,10],[3,3]], people = [3,3,2]`, func(t *testing.T) {
		assert.Equal(t, []int{2, 2, 1}, fullBloomFlowers([][]int{{1, 10}, {3, 3}}, []int{3, 3, 2}))
	})
}
