package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	t.Run(`with negetive number`, func(t *testing.T) {
		t.Run(`with x = -121`, func(t *testing.T) {
			assert.False(t, isPalindrome(-121))
		})

		t.Run(`with x = -2147483648`, func(t *testing.T) {
			assert.False(t, isPalindrome(-2147483648))
		})
	})

	t.Run(`with positive number`, func(t *testing.T) {
		t.Run(`with x = 121`, func(t *testing.T) {
			assert.True(t, isPalindrome(121))
		})

		t.Run(`with x = 2147483647`, func(t *testing.T) {
			assert.False(t, isPalindrome(2147483647))
		})
	})
}
