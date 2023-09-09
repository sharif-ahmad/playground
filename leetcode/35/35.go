package main

func searchInsert(nums []int, target int) int {
  s, e := 0, len(nums)-1

  for s <= e {
    m := (s + e) / 2

    if target == nums[m] {
      return m
    }

    if target < nums[m] {
      e = m - 1
    }

    if target > nums[m] {
      s = m + 1
    }
  }

  return s
}
