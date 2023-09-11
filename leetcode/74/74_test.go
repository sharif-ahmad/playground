package main

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestSearchMatrix(t *testing.T) {
  t.Run(`when target is found`, func(t *testing.T) {
    t.Run(`with mat = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], t = 3`, func(t *testing.T) {
      mat := [][]int{
        []int{1, 3, 5, 7},
        []int{10, 11, 16, 20},
        []int{23, 30, 34, 60},
      }

      assert.True(t, searchMatrix(mat, 3))
    })
  })

  t.Run(`when target not found`, func(t *testing.T) {
    t.Run(`with mat = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], t = 4`, func(t *testing.T) {
      mat := [][]int{
        []int{1, 3, 5, 7},
        []int{10, 11, 16, 20},
        []int{23, 30, 34, 60},
      }

      assert.False(t, searchMatrix(mat, 4))
    })
  })
}
