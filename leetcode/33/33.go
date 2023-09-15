package main

func search(nums []int, target int) int {
	s, e := 0, len(nums)-1

	for s <= e {
		m := (s + e) / 2

		if target == nums[m] {
			return m
		}

		if nums[s] <= nums[m] { // left half is sorted
			if nums[s] <= target && target < nums[m] {
				e = m - 1
			} else {
				s = m + 1
			}
		} else if nums[m] <= nums[e] { // right half is sorted
			if nums[m] < target && target <= nums[e] {
				s = m + 1
			} else {
				e = m - 1
			}
		}
	}

	return -1
}
