package main

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestPlusOne(t *testing.T) {
  t.Run(`with 9`, func(t *testing.T) {
    assert.Equal(t, []int{1, 0}, plusOne([]int{9}))
  })

  t.Run(`with 10`, func(t *testing.T) {
    assert.Equal(t, []int{1, 1}, plusOne([]int{1, 0}))
  })
}

func TestBigInt_Add(t *testing.T) {
  t.Run(`with equal length numbers`, func(t *testing.T) {
    t.Run(`with 999 + 111`, func(t *testing.T) {
      res := BigInt{digits: []int{9, 9, 9}}.Add(BigInt{digits: []int{1, 1, 1}})

      assert.Equal(t, []int{1, 1, 1, 0}, res.Digits())
    })
  })

  t.Run(`with unequal length numbers`, func(t *testing.T) {
    t.Run(`with 999 + 1`, func(t *testing.T) {
      res := BigInt{digits: []int{9, 9, 9}}.Add(BigInt{digits: []int{1}})

      assert.Equal(t, []int{1, 0, 0, 0}, res.Digits())
    })
  })
}
