package main

import "math"

func maxSubArray(nums []int) int {
	return _maxSubArray(nums, 0, len(nums)-1)
}

func _maxSubArray(nums []int, l, r int) int {
	if l == r {
		return nums[l]
	}

	m := (l + r) / 2
	lsum := _maxSubArray(nums, l, m)
	rsum := _maxSubArray(nums, m+1, r)
	csum := _maxCrossingSum(nums, l, m, r)

	return max(lsum, rsum, csum)
}

func _maxCrossingSum(nums []int, l, m, r int) int {
	lmx := math.MinInt
	lsum := 0
	for i := m; i >= l; i-- {
		lsum += nums[i]
		lmx = max(lmx, lsum)
	}

	rmx := math.MinInt
	rsum := 0
	for i := m + 1; i <= r; i++ {
		rsum += nums[i]
		rmx = max(rmx, rsum)
	}

	return lmx + rmx
}
