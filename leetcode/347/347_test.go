package main

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestTopKFrequent(t *testing.T) {
  t.Run(`with nums = [1,1,1,2,2,3], k = 2`, func(t *testing.T) {
    assert.Equal(t, []int{1, 2}, topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
  })

  t.Run(`with nums = [1], k = 1`, func(t *testing.T) {
    assert.Equal(t, []int{1}, topKFrequent([]int{1}, 1))
  })
}
