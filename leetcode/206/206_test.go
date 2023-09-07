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

func (l *ListNode) ToSlice() []int {
	res := make([]int, 0)

	for curr := l; curr != nil; curr = curr.Next {
		res = append(res, curr.Val)
	}

	return res
}

func linkListFromSlice(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	curr := head
	for _, n := range nums[1:] {
		curr = curr.addNextVal(n)
	}

	return head
}

func TestReverseList(t *testing.T) {
	t.Run(`with empty list`, func(t *testing.T) {
		head := linkListFromSlice([]int{})
		rHead := reverseList(head)

		assert.Nil(t, head)
		assert.Nil(t, rHead)
		assert.Equal(t, []int{}, rHead.ToSlice())
	})

	t.Run(`with single element`, func(t *testing.T) {
		head := linkListFromSlice([]int{1})
		rHead := reverseList(head)

		assert.Equal(t, head, rHead)
		assert.Equal(t, []int{1}, rHead.ToSlice())
	})

	t.Run(`with 2 elements`, func(t *testing.T) {
		head := linkListFromSlice([]int{1, 2})
		lastNode := head.nthNode(1)
		rHead := reverseList(head)

		assert.Equal(t, lastNode, rHead)
		assert.Equal(t, []int{2, 1}, rHead.ToSlice())
	})

	t.Run(`with multiple elements`, func(t *testing.T) {
		head := linkListFromSlice([]int{1, 2, 3, 4, 5})
		lastNode := head.nthNode(4)
		rHead := reverseList(head)

		assert.Equal(t, lastNode, rHead)
		assert.Equal(t, []int{5, 4, 3, 2, 1}, rHead.ToSlice())
	})
}
