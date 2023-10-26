package main

import "math"

func maxSubArray(nums []int) int {
	mx := math.MinInt
	sum := 0

	for _, n := range nums {
		sum += n
		if sum > mx {
			mx = sum
		}

		if sum < 0 {
			sum = 0
		}
	}

	return mx
}
