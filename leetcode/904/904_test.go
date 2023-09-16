package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTotalFruit(t *testing.T) {
	t.Run(`with fruits = [1,2,1]`, func(t *testing.T) {
		assert.Equal(t, 3, totalFruit([]int{1, 2, 1}))
	})

	t.Run(`with fruits = [0,1,2,2]`, func(t *testing.T) {
		assert.Equal(t, 3, totalFruit([]int{0, 1, 2, 2}))
	})

	t.Run(`with fruits = [1,2,3,2,2]`, func(t *testing.T) {
		assert.Equal(t, 4, totalFruit([]int{1, 2, 3, 2, 2}))
	})
}
