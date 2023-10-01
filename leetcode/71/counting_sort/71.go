package main

func sortColors(nums []int) {
	counts := make([]int, 3)

	for _, n := range nums {
		counts[n]++
	}

	j := 0
	for i, count := range counts {
		repeat(nums[j:], i, count)
		j += count
	}
}

func repeat[T any](A []T, a T, n int) {
	for i := 0; i < n; i++ {
		A[i] = a
	}
}
