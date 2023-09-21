package main

func findMaxAverage(nums []int, k int) float64 {
	maxAvg := -10001.0
	sum := 0
	for i, j := 0, 0; j < len(nums); j++ {
		sum += nums[j]
		if elementBetween(i, j) > k {
			sum -= nums[i]
			i++
		}

		if elementBetween(i, j) == k {
			avg := float64(sum) / float64(k)
			if avg > maxAvg {
				maxAvg = avg
			}
		}
	}

	return maxAvg
}

func elementBetween(i, j int) int {
	return j - i + 1
}
