package sort

import "sort"

// Time complexity: O(n * log n)
// Space complexity: O(n)

func findKthLargest(nums []int, k int) int {
  sort.Ints(nums)
  n := len(nums)

  return nums[n-k]
}
