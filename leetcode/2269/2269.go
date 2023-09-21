package main

import (
  "fmt"
  "strconv"
)

func divisorSubstrings(num int, k int) int {

  strNum := intToStr(num)

  count := 0
  for i := 0; i+k <= len(strNum); i++ {
    subNum := strToInt(strNum[i : i+k])
    if subNum != 0 && num%subNum == 0 {
      count++
    }
  }

  return count
}

func intToStr(n int) string {
  return fmt.Sprintf("%d", n)
}

func strToInt(str string) int {
  res, _ := strconv.ParseInt(str, 10, 32)

  return int(res)
}
