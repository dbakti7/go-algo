package strings

func buildLPSTable(word string) []int {
	if len(word) == 0 {
		return []int{}
	}

	table := make([]int, len(word))
	curLength := 0
	table[0] = 0 // base case, always 0 for first index
	for i := 1; i < len(word); i++ {
		for word[curLength] != word[i] && curLength > 0 {
			curLength = table[curLength-1]
		}
		if word[curLength] == word[i] {
			curLength++
			table[i] = curLength
		} else {
			table[i] = 0
		}
	}

	return table
}

func FindSubstrings(haystack, needle string) []int {
	j := 0
	results := make([]int, 0)
	lps := buildLPSTable(needle)

	for i := 0; i < len(haystack); i++ {
		for haystack[i] != needle[j] && j > 0 {
			j = lps[j-1]
		}

		if haystack[i] == needle[j] {
			j++
			if j == len(needle) {
				results = append(results, i-j+1)
				j = lps[j-1]
			}
			continue
		}
	}
	return results
}
