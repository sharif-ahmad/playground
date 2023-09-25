package main

func minSubArrayLen(target int, nums []int) int {
	minLen := 999999
	subArraySum := 0

	for i, j := 0, 0; i < len(nums); i++ {
		subArraySum += nums[i]

		for subArraySum >= target {
			currLen := i - j + 1
			if currLen < minLen {
				minLen = currLen
			}

			subArraySum -= nums[j]
			j++
		}
	}

	if minLen == 999999 {
		return 0
	}

	return minLen
}
