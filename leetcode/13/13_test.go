package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	t.Run(`with s = "III"`, func(t *testing.T) {
		assert.Equal(t, 3, romanToInt("III"))
	})

	t.Run(`with s = "LVIII"`, func(t *testing.T) {
		assert.Equal(t, 58, romanToInt("LVIII"))
	})

	t.Run(`with s = "MCMXCIV"`, func(t *testing.T) {
		assert.Equal(t, 1994, romanToInt("MCMXCIV"))
	})
}
