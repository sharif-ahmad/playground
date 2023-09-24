package main

func findRepeatedDnaSequences(s string) []string {
	subStringSet := make(map[string]int)
	res := make([]string, 0)

	for i := 0; i+10 <= len(s); i++ {
		if subStringSet[s[i:i+10]] == 1 {
			res = append(res, s[i:i+10])
		}
		subStringSet[s[i:i+10]]++
	}

	return res
}
