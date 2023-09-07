package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func (l *ListNode) addNextVal(v int) *ListNode {
	curr := l
	for ; curr.Next != nil; curr = curr.Next {
	}

	curr.Next = &ListNode{Val: v}

	return curr.Next
}

// n is 0-based index/position of the node
func (l *ListNode) nthNode(n int) *ListNode {
	if n == 0 {
		return l
	}

	return l.Next.nthNode(n - 1)
}

func linkListFromSlice(nums []int) *ListNode {
	head := &ListNode{Val: nums[0]}
	curr := head
	for _, n := range nums[1:] {
		curr = curr.addNextVal(n)
	}

	return head
}

func TestMiddleNode(t *testing.T) {
	t.Run(`with even number of nodes`, func(t *testing.T) {
		t.Run(`with [1, 2, 3, 4, 5, 6]`, func(t *testing.T) {
			head := linkListFromSlice([]int{1, 2, 3, 4, 5})

			assert.Equal(t, head.nthNode(2), middleNode(head))
		})

		t.Run(`with [1, 2]`, func(t *testing.T) {
			head := linkListFromSlice([]int{1, 2})

			assert.Equal(t, head.nthNode(1), middleNode(head))
		})
	})

	t.Run(`with odd number of nodes`, func(t *testing.T) {
		t.Run(`with [1, 2, 3, 4, 5]`, func(t *testing.T) {
			head := linkListFromSlice([]int{1, 2, 3, 4, 5})

			assert.Equal(t, head.nthNode(2), middleNode(head))
		})

		t.Run(`with [1]`, func(t *testing.T) {
			head := linkListFromSlice([]int{1})

			assert.Equal(t, head.nthNode(0), middleNode(head))
		})

		t.Run(`with [1, 2, 3]`, func(t *testing.T) {
			head := linkListFromSlice([]int{1, 2, 3})

			assert.Equal(t, head.nthNode(1), middleNode(head))
		})
	})
}
