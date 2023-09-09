package main

func plusOne(digits []int) []int {
  res := NewBigIntFrom(digits).Add(NewBigIntFrom([]int{1}))

  return res.Digits()
}

type BigInt struct {
  digits []int
}

func NewBigIntFrom(digits []int) BigInt {
  return BigInt{digits: digits}
}

func (i BigInt) Digits() []int {
  return i.digits
}

func (i BigInt) Add(j BigInt) BigInt {
  a, b := len(i.digits)-1, len(j.digits)-1
  carry := 0
  res := make([]int, 0)

  for ; a >= 0 && b >= 0; a, b = a-1, b-1 {
    currSum := i.digits[a] + j.digits[b] + carry
    res = append(res, currSum%10)
    carry = currSum / 10
  }

  for ; a >= 0; a-- {
    currSum := i.digits[a] + carry
    res = append(res, currSum%10)
    carry = currSum / 10
  }

  for ; b >= 0; b-- {
    currSum := j.digits[b] + carry
    res = append(res, currSum%10)
    carry = currSum / 10
  }

  if carry > 0 {
    res = append(res, carry)
  }

  return BigInt{digits: reverseSlice(res)}
}

func reverseSlice[T any](s []T) []T {
  i, j := 0, len(s)-1

  for i < j {
    swap(&s[i], &s[j])
    i++
    j--
  }

  return s
}

func swap[T any](a *T, b *T) {
  *a, *b = *b, *a
}
