package main

// sortColors
// [0,l) all 0s
// [l, i) all 1s
// [i, r] unprocessed
// (r,len(nums)) all 2s
func sortColors(nums []int) {
	l, i, r := 0, 0, len(nums)-1

	for i <= r {
		if nums[i] == 0 {
			swap(&nums[l], &nums[i])
			l++
			i++
		} else if nums[i] == 2 {
			swap(&nums[i], &nums[r])
			r--
		} else {
			i++
		}
	}
}

func swap[T any](a, b *T) {
	*a, *b = *b, *a
}
