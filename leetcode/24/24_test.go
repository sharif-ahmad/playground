package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func sliceToList(as []int) *ListNode {
	if len(as) == 0 {
		return nil
	}

	return &ListNode{Val: as[0], Next: sliceToList(as[1:])}
}

func ListToSlice(head *ListNode) []int {
	res := make([]int, 0)

	for ; head != nil; head = head.Next {
		res = append(res, head.Val)
	}

	return res
}

func TestSwapPairs(t *testing.T) {
	t.Run(`with head = [1,2,3,4]`, func(t *testing.T) {
		assert.Equal(t, []int{2, 1, 4, 3}, ListToSlice(swapPairs(sliceToList([]int{1, 2, 3, 4}))))
	})

	t.Run(`with head = []`, func(t *testing.T) {
		assert.Equal(t, []int{}, ListToSlice(swapPairs(sliceToList([]int{}))))
	})

	t.Run(`with head = [1]`, func(t *testing.T) {
		assert.Equal(t, []int{1}, ListToSlice(swapPairs(sliceToList([]int{1}))))
	})

	t.Run(`with head = [1,2,3]`, func(t *testing.T) {
		assert.Equal(t, []int{2, 1, 3}, ListToSlice(swapPairs(sliceToList([]int{1, 2, 3}))))
	})
}
