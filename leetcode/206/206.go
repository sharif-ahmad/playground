package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func Reverse(prev, curr *ListNode) *ListNode {
	if curr == nil {
		return prev
	}

	reversHead := Reverse(curr, curr.Next)

	curr.Next = prev

	return reversHead
}

func reverseList(head *ListNode) *ListNode {
	return Reverse(nil, head)
}
