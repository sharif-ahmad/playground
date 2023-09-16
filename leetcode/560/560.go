package main

func subarraySum(nums []int, k int) int {
	var prefixSum []int

	prefixSum = append(prefixSum, nums[0])
	for i := 1; i < len(nums); i++ {
		prefixSum = append(prefixSum, prefixSum[i-1]+nums[i])
	}

	prefixSumMap := make(map[int]int)
	res := 0
	for i := 0; i < len(prefixSum); i++ {
		if prefixSum[i] == k {
			res++
		}

		res += prefixSumMap[prefixSum[i]-k]

		prefixSumMap[prefixSum[i]]++
	}

	return res
}
