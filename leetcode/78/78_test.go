package main

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestSubsets(t *testing.T) {
  t.Run(`with nums = [1,2,3]`, func(t *testing.T) {
    assert.ElementsMatch(t, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}, subsets([]int{1, 2, 3}))
  })

  t.Run(`with nums = [0]`, func(t *testing.T) {
    assert.ElementsMatch(t, [][]int{{}, {0}}, subsets([]int{0}))
  })
}
