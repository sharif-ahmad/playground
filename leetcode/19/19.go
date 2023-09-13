package main

type ListNode struct {
  Val  int
  Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
  var prev *ListNode
  curr, forward := head, head

  for i := 1; i < n; i++ {
    forward = forward.Next
  }

  for forward.Next != nil {
    prev = curr
    forward = forward.Next
    curr = curr.Next
  }

  if prev != nil {
    prev.Next = curr.Next
    return head
  }

  return curr.Next
}
