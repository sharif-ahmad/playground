package main

func searchRange(nums []int, target int) []int {
	lb := lowerBound(nums, target)
	if lb >= len(nums) || target != nums[lb] {
		return []int{-1, -1}
	}

	ub := upperBound(nums, target)

	return []int{lb, ub - 1}
}

// binarySearch C++ binary_search
func binarySearch(nums []int, target int, predicat func(int, int) bool) int {
	s, e := 0, len(nums)-1

	for s <= e {
		m := (s + e) / 2
		if predicat(nums[m], target) {
			s = m + 1
		} else {
			e = m - 1
		}
	}

	return s
}

// lowerBound C++ lower_bound
func lowerBound(nums []int, target int) int {
	return binarySearch(nums, target, func(a, b int) bool {
		return a < b
	})
}

// upperBound C++ upper_bound
func upperBound(nums []int, target int) int {
	return binarySearch(nums, target, func(a, b int) bool {
		return a <= b
	})
}
