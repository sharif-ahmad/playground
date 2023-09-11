package main

func searchMatrix(mat [][]int, t int) bool {
  s, e := 0, len(mat)-1

  for s <= e {
    m := (s + e) / 2

    if t < mat[m][0] {
      e = m - 1
    }

    if t >= mat[m][0] {
      if t <= mat[m][len(mat[m])-1] {
        return binarySearch(mat[m], t)
      }
      s = m + 1
    }
  }

  return false
}

func binarySearch(arr []int, t int) bool {
  s, e := 0, len(arr)-1

  for s <= e {
    m := (s + e) / 2

    if t == arr[m] {
      return true
    }

    if t < arr[m] {
      e = m - 1
    }

    if t > arr[m] {
      s = m + 1
    }
  }

  return false
}
