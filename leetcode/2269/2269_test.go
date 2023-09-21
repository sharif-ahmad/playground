package main

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestDivisorSbustrings(t *testing.T) {
  t.Run(`with num = 240, k = 2`, func(t *testing.T) {
    assert.Equal(t, 2, divisorSubstrings(240, 2))
  })

  t.Run(`with num = 430043, k = 2`, func(t *testing.T) {
    assert.Equal(t, 2, divisorSubstrings(430043, 2))
  })
}
