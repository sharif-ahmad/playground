package main

import "sort"

func maxFrequency(nums []int, k int) int {
  sort.Ints(nums)
  currSum := 0
  currMaxFreq := -99999999

  for i, j := 0, 0; i < len(nums); i++ {
    currSum += nums[i]
    count := elementBetween(j, i)
    if currSum+k < count*nums[i] {
      currSum -= nums[j]
      count--
      j++
    }

    if count > currMaxFreq {
      currMaxFreq = count
    }
  }

  return currMaxFreq
}

func elementBetween(i, j int) int {
  return j - i + 1
}
