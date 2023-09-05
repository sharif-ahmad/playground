package quick_sort

// Time complexity: O(n * log n)
// Space complexity: O(n)

func findKthLargest(nums []int, k int) int {
  n := len(nums)

  return find_recursive(nums, 0, n-1, n-k)
}

// quick sort technique
// t: target index
func find_recursive(nums []int, l, r, t int) int {
  p := partition(nums, l, r)

  if p > t {
    return find_recursive(nums, l, p-1, t)
  } else if p < t {
    return find_recursive(nums, p+1, r, t)
  }

  return nums[p]
}

// Lomuto partition algorithm
// nums: the number array
// l: the left index
// r: right index
// returns the new (correct) index for the pivot element
func partition(nums []int, l, r int) int {
  p := nums[r] // the pivot
  i, j := l-1, l
  // main loop
  for ; j <= r-1; j++ {
    if nums[j] <= p {
      i++
      swap(&nums[i], &nums[j])
    }
  }

  i++
  swap(&nums[i], &nums[r])

  return i
}

func swap[T any](a, b *T) {
  *a, *b = *b, *a
}
