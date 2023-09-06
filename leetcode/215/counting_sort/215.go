package counting_sort

// Does not work with negative numbers - got runtime error in leetcode

// Time complexity: O(n)
// Memory complexity: O(n)

func findKthLargest(nums []int, k int) int {
	m := findMax(nums)
	countList := make([]int, m+1)

	for _, n := range nums {
		countList[n]++
	}

	t := len(nums) - k + 1
	for i, n := range countList {
		t = t - n
		if t <= 0 {
			return i
		}
	}

	return -1
}

func findMax(nums []int) int {
	m := nums[0]
	for _, n := range nums[1:] {
		m = _max(m, n)
	}

	return m
}

func _max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
