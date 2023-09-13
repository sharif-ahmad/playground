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

func TestRemoveNthFromEnd(t *testing.T) {
  t.Run(`when remove from the middle of the list`, func(t *testing.T) {
    t.Run(`with head = [1,2,3,4,5], n = 2`, func(t *testing.T) {
      list := linkListFromSlice([]int{1, 2, 3, 4, 5})
      result := removeNthFromEnd(list, 2)

      assert.Equal(t, []int{1, 2, 3, 5}, result.ToSlice())
    })

    t.Run(`with head = [1,2,3], n = 2`, func(t *testing.T) {
      list := linkListFromSlice([]int{1, 2, 3})
      result := removeNthFromEnd(list, 2)

      assert.Equal(t, []int{1, 3}, result.ToSlice())
    })
  })

  t.Run(`when remove from the begining of the list`, func(t *testing.T) {
    t.Run(`with head = [1], n = 1`, func(t *testing.T) {
      list := linkListFromSlice([]int{1})
      result := removeNthFromEnd(list, 1)

      assert.Equal(t, []int{}, result.ToSlice())
    })

    t.Run(`with head = [1,2], n = 2`, func(t *testing.T) {
      list := linkListFromSlice([]int{1, 2})
      result := removeNthFromEnd(list, 2)

      assert.Equal(t, []int{2}, result.ToSlice())
    })
  })
}
