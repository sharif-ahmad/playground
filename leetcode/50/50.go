package main

func myPow(x float64, n int) float64 {
  if n < 0 {
    return 1.0 / myPow(x, -1*n)
  }

  if n == 0 {
    return 1
  }

  y := myPow(x, n/2)
  if isEven(n) {
    return y * y
  }

  return y * y * x
}

func isEven(n int) bool {
  return n%2 == 0
}
