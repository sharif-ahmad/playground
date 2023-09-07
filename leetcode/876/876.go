package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slowP, fastP := head, head

	for fastP != nil && fastP.Next != nil {
		slowP = slowP.Next
		fastP = fastP.Next.Next
	}

	return slowP
}
