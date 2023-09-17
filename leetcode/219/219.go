package main

func containsNearbyDuplicate(nums []int, k int) bool {
	freqMap := make(map[int]int)
	freqMap[nums[0]]++
	for i, j := 0, 1; j < len(nums); j++ {
		if j-i > k {
			freqMap[nums[i]]--
			i++
		}

		if freqMap[nums[j]] > 0 {
			return true
		}

		freqMap[nums[j]]++
	}

	return false
}
