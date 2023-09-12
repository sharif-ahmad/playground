package main

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestCanCross(t *testing.T) {
  t.Run(`with [0,1,3,5,6,8,12,17]`, func(t *testing.T) {
    assert.True(t, canCross([]int{0, 1, 3, 5, 6, 8, 12, 17}))
  })

  t.Run(`with [0,1,2,3,4,8,9,11]`, func(t *testing.T) {
    assert.False(t, canCross([]int{0, 1, 2, 3, 4, 8, 9, 11}))
  })
}
