package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	t.Run(`with num = 3`, func(t *testing.T) {
		assert.Equal(t, "III", intToRoman(3))
	})

	t.Run(`with num = 58`, func(t *testing.T) {
		assert.Equal(t, "LVIII", intToRoman(58))
	})

	t.Run(`with num = 1994`, func(t *testing.T) {
		assert.Equal(t, "MCMXCIV", intToRoman(1994))
	})
}
