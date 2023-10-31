package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	rest := swapPairs(head.Next.Next)

	second := head.Next
	head.Next = rest
	second.Next = head

	return second
}
