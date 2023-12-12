package main

func subsets(nums []int) [][]int {
  currentSubset = make([]int, 0)

  return recursiveSubset(nums, 0)
}

var currentSubset []int

func recursiveSubset(nums []int, i int) [][]int {
  if i >= len(nums) {
    cp := make([]int, len(currentSubset)) // May God bless golang and its slice
    copy(cp, currentSubset)
    return [][]int{cp}
  }

  // skip the current element
  r1 := recursiveSubset(nums, i+1)

  // take the current element
  currentSubset = append(currentSubset, nums[i])
  r2 := recursiveSubset(nums, i+1)
  currentSubset = currentSubset[:len(currentSubset)-1]

  return append(r1, r2...)
}
