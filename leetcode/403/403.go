package main

func canCross(stones []int) bool {
  return NewFogJump(stones).canCross(stones[0], 0)
}

type FrogJump struct {
  stoneMap  map[int]struct{}
  lastStone int
  memo      map[int]map[int]bool
}

func NewFogJump(stones []int) *FrogJump {
  n := len(stones)

  stoneMap := make(map[int]struct{}) // using as a set
  for _, s := range stones {
    stoneMap[s] = struct{}{}
  }

  memo := make(map[int]map[int]bool)
  for _, k := range stones {
    memo[k] = make(map[int]bool)
  }

  return &FrogJump{
    stoneMap:  stoneMap,
    lastStone: stones[n-1],
    memo:      memo,
  }
}

func (f *FrogJump) getMemo(i, j int) (bool, bool) {
  row, found := f.memo[i]
  if !found {
    return false, false
  }

  col, found := row[j]

  return col, found
}

func (f *FrogJump) setMemo(i, j int, v bool) {
  f.memo[i][j] = v
}

// i - current stone
// k - previous jump length
func (f *FrogJump) canCross(i, k int) bool {
  if i == f.lastStone {
    return true
  }

  if _, found := f.stoneMap[i]; !found {
    return false
  }

  if res, found := f.getMemo(i, k); found {
    return res
  }

  res := false

  for nj := k - 1; nj <= k+1; nj++ {
    if nj <= 0 {
      continue
    }

    res = res || f.canCross(i+nj, nj)
  }

  f.setMemo(i, k, res)

  return res
}
