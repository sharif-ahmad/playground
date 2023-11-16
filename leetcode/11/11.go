package main

import "math"

func maxArea(height []int) int {
	res := math.MinInt
	i, j := 0, len(height)-1
	for i < j {
		w := _min(height[i], height[j]) * (j - i)
		res = _max(res, w)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return res
}

func _max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func _min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
