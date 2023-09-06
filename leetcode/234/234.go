package main

import (
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	str := linkListToStr(head)

	return isPalindromeStr(str)
}

// O(n)
// n - number of nodes
func linkListToStr(head *ListNode) string {
	var s strings.Builder
	for curr := head; curr != nil; curr = curr.Next {
		s.WriteRune(intToChar(curr.Val))
	}

	return s.String()
}

func intToChar(i int) rune {
	return (rune)('0' - i)
}

func isPalindromeStr(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}
