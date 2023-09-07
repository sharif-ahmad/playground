package heap

import (
	"cmp"
	"github.com/stretchr/testify/assert"
	"testing"
)

// auxiliary method for testing
func (h *Heap[T]) getArr() []T {
	return h.arr[0:h.size]
}

type GTCompareFn[T cmp.Ordered] struct{}

func (c *GTCompareFn[T]) call(a, b T) bool {
	return a > b
}

func TestFindKthLargest(t *testing.T) {

	assert.Equal(t, 5, findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	assert.Equal(t, 4, findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	assert.Equal(t, 2, findKthLargest([]int{1, 1, 1, 2, 3, 4, 4}, 4))
	assert.Equal(t, 6, findKthLargest([]int{7, 6, 5, 4, 3, 2, 1}, 2))
	assert.Equal(t, 7, findKthLargest([]int{7, 7}, 2))
	assert.Equal(t, 5, findKthLargest([]int{5}, 1))
	assert.Equal(t, 3, findKthLargest([]int{7, 6, 5, 4, 3, 2, 1}, 5))
}

func TestNewHeapFrom(t *testing.T) {
	A := []int{1, 2, 4, 5, 5, 6}
	h := NewHeapFrom(A, &GTCompareFn[int]{})

	assert.Equal(t, 6, h.Top())
	assert.Equal(t, []int{6, 5, 4, 2, 5, 1}, h.getArr())
}

func TestMaxHeap_Insert(t *testing.T) {
	t.Run(`when heap is built with NewMaxHeap`, func(t *testing.T) {

		t.Run(`when heap is empty`, func(t *testing.T) {
			h := NewHeap[int](2, &GTCompareFn[int]{})
			err := h.Insert(5)

			assert.Nil(t, err)
			assert.Equal(t, []int{5}, h.getArr())
		})

		t.Run(`when heap is not empty`, func(t *testing.T) {
			h := NewHeap[int](2, &GTCompareFn[int]{})
			h.Insert(5)

			err := h.Insert(7)

			assert.Nil(t, err)
			assert.Equal(t, []int{7, 5}, h.getArr())
		})

		t.Run(`when heap is full`, func(t *testing.T) {
			h := NewHeap[int](2, &GTCompareFn[int]{})
			h.Insert(5)
			h.Insert(7)

			err := h.Insert(3)

			assert.NotNil(t, err)
		})

		t.Run(`example`, func(t *testing.T) {
			h := NewHeap[int](4, &LTCompareFn[int]{})
			h.Insert(3)
			h.Insert(3)
			h.Insert(4)
			h.Insert(5)
			h.ExtractTop()
			h.Insert(5)

			assert.Equal(t, 3, h.Top())
		})

	})

	//t.Run(`when heap is built with NewMaxHeapFrom`, func(t *testing.T) {
	//
	//})
}

func TestHeap_ExtractTop(t *testing.T) {
	t.Run(`When heap is full`, func(t *testing.T) {
		h := NewHeapFrom([]int{1, 1, 1, 2}, &LTCompareFn[int]{})

		assert.Equal(t, 1, h.ExtractTop())
	})
}
