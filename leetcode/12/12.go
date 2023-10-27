package main

import (
	"sort"
	"strings"
)

var valueList []int
var symbolList []string

func intToSymbol(i int) (int, string) {
	if valueList == nil && symbolList == nil {
		valueList = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
		symbolList = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	}

	k := sort.SearchInts(valueList, i)
	if k < len(valueList) && valueList[k] == i {
		return valueList[k], symbolList[k]
	}

	return valueList[k-1], symbolList[k-1]
}

func intToRoman(num int) string {
	res := strings.Builder{}

	for num > 0 {
		k, v := intToSymbol(num)
		res.WriteString(v)
		num -= k
	}

	return res.String()
}
