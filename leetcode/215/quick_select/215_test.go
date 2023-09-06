package quick_select

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestFindKthLargest(t *testing.T) {

  assert.Equal(t, 5, findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
  assert.Equal(t, 4, findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
  assert.Equal(t, 2, findKthLargest([]int{1, 1, 1, 2, 3, 4, 4}, 4))
  assert.Equal(t, 6, findKthLargest([]int{7, 6, 5, 4, 3, 2, 1}, 2))
  assert.Equal(t, 7, findKthLargest([]int{7, 7}, 2))
  assert.Equal(t, 5, findKthLargest([]int{5}, 1))
  assert.Equal(t, 3, findKthLargest([]int{7, 6, 5, 4, 3, 2, 1}, 5))
}

func TestPartition(t *testing.T) {
  assert.Equal(t, 0, partition([]int{7, 6, 5, 4, 3, 2, 1}, 0, 6))
}

func TestSwap_normal_variable(t *testing.T) {
  var a = 5
  var b = 10

  swap(&a, &b)

  assert.Equal(t, 10, a)
  assert.Equal(t, 5, b)
}

func TestSwap_array(t *testing.T) {
  A := []int{1, 10}

  swap(&A[0], &A[1])

  assert.Equal(t, []int{10, 1}, A)
}
