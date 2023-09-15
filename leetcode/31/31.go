package main

// Implementation from c++ stl std::next_permutation
func nextPermutation(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			j := reverseUpperBound(nums[i:], nums[i-1])
			nums[i-1], nums[i+j] = nums[i+j], nums[i-1] // swap with the smallest number larger than (i-1)th
			reverseSlice(nums[i:])
			return
		}
	}

	reverseSlice(nums)
}

// returns the index of last element that is larger than t in a reverse sorted slice, -1 if t is the largest element
func reverseUpperBound(nums []int, t int) int {
	e, s := 0, len(nums)-1 // because it's reverse

	for s >= e {
		m := (s + e) / 2

		if nums[m] <= t {
			s = m - 1
		} else {
			e = m + 1
		}
	}

	return s
}

func reverseSlice(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
