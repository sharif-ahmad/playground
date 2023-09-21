package main

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestMaxProfit(t *testing.T) {
  t.Run(`with prices = [7,1,5,3,6,4]`, func(t *testing.T) {
    assert.Equal(t, 5, maxProfit([]int{7, 1, 5, 3, 6, 4}))
  })

  t.Run(`with prices = [[7,6,4,3,1]]`, func(t *testing.T) {
    assert.Equal(t, 0, maxProfit([]int{7, 6, 4, 3, 1}))
  })
}
