package main

func symbolToValue(s string) (int, bool) {
	valueMap := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	r, ok := valueMap[s]
	return r, ok
}

func romanToInt(s string) int {
	res := 0
	for i := 0; i < len(s); {
		if i+1 < len(s) {
			if v, ok := symbolToValue(s[i : i+2]); ok {
				res += v
				i = i + 2
			} else {
				v, _ := symbolToValue(s[i : i+1])
				res += v
				i = i + 1
			}
		} else {
			v, _ := symbolToValue(s[i : i+1])
			res += v
			i = i + 1
		}
	}

	return res
}
