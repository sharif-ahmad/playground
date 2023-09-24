package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindRepeatedDnaSequences(t *testing.T) {
	t.Run(`with s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"`, func(t *testing.T) {
		assert.ElementsMatch(t, []string{"AAAAACCCCC", "CCCCCAAAAA"}, findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	})

	t.Run(`with s = "AAAAAAAAAAAAA"`, func(t *testing.T) {
		assert.ElementsMatch(t, []string{"AAAAAAAAAA"}, findRepeatedDnaSequences("AAAAAAAAAAAAA"))
	})

	t.Run(`with s = "AAAAAAAAAAGAAAAAAAAAA"`, func(t *testing.T) {
		assert.ElementsMatch(t, []string{"AAAAAAAAAA"}, findRepeatedDnaSequences("AAAAAAAAAAGAAAAAAAAAA"))
	})
}
