package main

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestMaxFrequency(t *testing.T) {
  t.Run(`with nums = [1,2,4], k = 5`, func(t *testing.T) {
    assert.Equal(t, 3, maxFrequency([]int{1, 2, 4}, 5))
  })

  t.Run(`with nums = [1,4,8,13], k = 5`, func(t *testing.T) {
    assert.Equal(t, 2, maxFrequency([]int{1, 4, 8, 13}, 5))
  })

  t.Run(`with nums = [3,9,6], k = 2`, func(t *testing.T) {
    assert.Equal(t, 1, maxFrequency([]int{3, 9, 6}, 2))
  })
}
