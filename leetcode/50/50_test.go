package main

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestMyPow(t *testing.T) {
  t.Run(`with x = 2.00000, n = 10`, func(t *testing.T) {
    assert.InDelta(t, 1024.00000, myPow(2.00000, 10), 0.000001)
  })

  t.Run(`with x = 2.10000, n = 3`, func(t *testing.T) {
    assert.InDelta(t, 9.26100, myPow(2.10000, 3), 0.000001)
  })

  t.Run(`with x = 2.00000, n = -2`, func(t *testing.T) {
    assert.InDelta(t, 0.25000, myPow(2.00000, -2), 0.000001)
  })
}
