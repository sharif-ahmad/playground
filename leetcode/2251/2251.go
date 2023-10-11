package main

import "sort"

func fullBloomFlowers(flowers [][]int, people []int) []int {
  starts := make([]int, 0)
  ends := make([]int, 0)

  for _, flower := range flowers {
    starts = append(starts, flower[0])
    ends = append(ends, flower[1])
  }

  sort.Ints(starts)
  sort.Ints(ends)

  ans := make([]int, 0)
  for _, p := range people {
    nStarted := sort.Search(len(starts), func(i int) bool { return starts[i] > p })
    nEnded := sort.Search(len(ends), func(i int) bool { return ends[i] >= p })

    ans = append(ans, nStarted-nEnded)
  }

  return ans
}
