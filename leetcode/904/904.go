package main

func totalFruit(fruits []int) int {
	left, right := 0, 0
	fruitCount := make(map[int]int)
	maxFruit := 0
	currCount := 0

	for left < len(fruits) && right < len(fruits) {
		fruitCount[fruits[right]]++
		currCount++

		for len(fruitCount) > 2 { // we have added more than 2 types of fruit, we need to delete some
			fruitCount[fruits[left]]--
			currCount--
			if fruitCount[fruits[left]] == 0 {
				delete(fruitCount, fruits[left])
			}
			left++
		}

		if currCount > maxFruit {
			maxFruit = currCount
		}

		right++
	}

	return maxFruit
}
