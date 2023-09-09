package main

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestSearchInsert(t *testing.T) {
  t.Run(`with empty list`, func(t *testing.T) {
    assert.Equal(t, 0, searchInsert([]int{}, 5))
  })

  t.Run(`with single element list`, func(t *testing.T) {
    t.Run(`when target is larger than all elements`, func(t *testing.T) {
      assert.Equal(t, 1, searchInsert([]int{4}, 6))
    })

    t.Run(`when target is smaller than all elements`, func(t *testing.T) {
      assert.Equal(t, 0, searchInsert([]int{4}, 2))
    })

    t.Run(`when target is equal to an element int the array`, func(t *testing.T) {
      assert.Equal(t, 0, searchInsert([]int{4}, 2))
    })
  })

  t.Run(`with more elements`, func(t *testing.T) {
    t.Run(`when element is found in the array`, func(t *testing.T) {
      assert.Equal(t, 3, searchInsert([]int{1, 3, 5, 6, 7, 8, 9, 23, 56}, 6))
    })

    t.Run(`when element is not found in the array`, func(t *testing.T) {
      t.Run(`when target's position is in between the start and the end`, func(t *testing.T) {
        assert.Equal(t, 7, searchInsert([]int{1, 3, 5, 6, 7, 8, 9, 23, 56}, 21))
      })

      t.Run(`when target is larger than all elements`, func(t *testing.T) {
        assert.Equal(t, 9, searchInsert([]int{1, 3, 5, 6, 7, 8, 9, 23, 56}, 109))
      })

      t.Run(`when target is smaller than all elements`, func(t *testing.T) {
        assert.Equal(t, 0, searchInsert([]int{2, 3, 5, 6, 7, 8, 9, 23, 56}, 1))
      })
    })
  })
}
