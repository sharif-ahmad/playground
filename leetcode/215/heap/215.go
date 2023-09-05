package heap

import (
	"cmp"
	"errors"
)

// Could not submit this in leetcode because it does not support go 1.21 yet. :(
// cmp.Ordering is 1.21 feature

// Time complexity: O(n * log k)
// Space complexity: O(n + k)

func findKthLargest(nums []int, k int) int {
	minHeap := NewHeap[int](k, &LTCompareFn[int]{})

	for _, x := range nums {
		if minHeap.size < k {
			minHeap.Insert(x)
		} else if x > minHeap.Top() {
			minHeap.ExtractTop()
			minHeap.Insert(x)
		}
	}

	return minHeap.Top()
}

// auxiliary data structures

type LTCompareFn[T cmp.Ordered] struct{}

func (receiver *LTCompareFn[T]) call(a, b T) bool {
	return a < b
}

type CompareFn[T any] interface {
	call(T, T) bool
}

type Heap[T any] struct {
	cap       int
	size      int
	arr       []T
	compareFn CompareFn[T]
}

func NewHeap[T any](cap int, compareFn CompareFn[T]) *Heap[T] {
	return &Heap[T]{
		cap:       cap,
		size:      0,
		arr:       make([]T, cap),
		compareFn: compareFn,
	}
}

func NewHeapFrom[T any](arr []T, compareFn CompareFn[T]) *Heap[T] {
	size := len(arr)

	h := &Heap[T]{
		cap:       size,
		size:      size,
		arr:       arr,
		compareFn: compareFn,
	}

	for i := size/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}

	return h
}

func (h *Heap[T]) Size() int {
	return h.size
}

func (h *Heap[T]) parent(c int) int {
	return (c - 1) / 2
}

func (h *Heap[T]) leftChild(p int) int {
	return 2*p + 1
}

func (h *Heap[T]) rightChild(p int) int {
	return 2*p + 2
}

func (h *Heap[T]) heapify(i int) {
	newParent := i

	l := h.leftChild(i)
	if l < h.size && h.compareFn.call(h.arr[l], h.arr[newParent]) {
		newParent = l
	}

	r := h.rightChild(i)
	if r < h.size && h.compareFn.call(h.arr[r], h.arr[newParent]) {
		newParent = r
	}

	if newParent == i {
		return
	}

	swap(&h.arr[newParent], &h.arr[i])
	h.heapify(newParent)
}

func (h *Heap[T]) Insert(x T) error {
	if h.size == h.cap {
		return errors.New("capacity full")
	}

	i := h.size
	h.size++

	h.arr[i] = x

	for i > 0 {
		p := h.parent(i)
		if !h.compareFn.call(h.arr[i], h.arr[p]) {
			break
		}
		swap(&h.arr[i], &h.arr[p])

		i = p
	}

	return nil
}

func (h *Heap[T]) Top() T {
	return h.arr[0]
}

func (h *Heap[T]) ExtractTop() T {
	swap(&h.arr[0], &h.arr[h.size-1])
	h.size--

	h.heapify(0)

	return h.arr[h.size]
}

func swap[T any](a, b *T) {
	*a, *b = *b, *a
}
