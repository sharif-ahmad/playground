package main

import "sort"

func topKFrequent(nums []int, k int) []int {
  freq := calcFreq(nums)
  sort.Slice(freq, func(i, j int) bool {
    return freq[i].v > freq[j].v
  })

  res := make([]int, 0)

  for _, e := range freq[0:k] {
    res = append(res, e.k)
  }

  return res
}

type mapElem struct {
  k, v int
}

func calcFreq(arr []int) []mapElem {
  freqMap := make(map[int]int)

  for _, a := range arr {
    freqMap[a] = freqMap[a] + 1
  }

  freqSlice := make([]mapElem, 0)
  for k, v := range freqMap {
    freqSlice = append(freqSlice, mapElem{k: k, v: v})
  }

  return freqSlice
}
